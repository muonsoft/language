package language

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"golang.org/x/text/language"
)

func ExampleMiddleware_ServeHTTP_readFromAcceptLanguageHeader() {
	h := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		tag := FromContext(request.Context())
		fmt.Println("language:", tag)
	})
	m := NewMiddleware(h, SupportedLanguages(language.English, language.Russian))

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("Accept-Language", "ru")
	w := httptest.NewRecorder()

	m.ServeHTTP(w, r)
	// Output: language: ru
}

func ExampleMiddleware_ServeHTTP_readFromCookie() {
	h := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		tag := FromContext(request.Context())
		fmt.Println("language:", tag)
	})
	m := NewMiddleware(
		h,
		SupportedLanguages(language.English, language.Russian),
		ReadFromCookie("lang"),
	)

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.AddCookie(&http.Cookie{Name: "lang", Value: "ru"})
	w := httptest.NewRecorder()

	m.ServeHTTP(w, r)
	// Output: language: ru
}

func TestMiddleware_ServeHTTP(t *testing.T) {
	tests := []struct {
		name             string
		request          *http.Request
		options          []MiddlewareOption
		expectedLanguage language.Tag
	}{
		{
			name:             "no options",
			request:          httptest.NewRequest(http.MethodGet, "/", nil),
			options:          nil,
			expectedLanguage: language.English,
		},
		{
			name:             "supported languages with one language",
			request:          httptest.NewRequest(http.MethodGet, "/", nil),
			options:          []MiddlewareOption{SupportedLanguages(language.Russian)},
			expectedLanguage: language.Russian,
		},
		{
			name:    "supported languages and source in cookie",
			request: givenRequestWithCookie("language_cookie", "ru"),
			options: []MiddlewareOption{
				SupportedLanguages(language.English, language.Russian),
				ReadFromCookie("language_cookie"),
			},
			expectedLanguage: language.Russian,
		},
		{
			name:    "supported languages and source in header",
			request: givenRequestWithAcceptLanguageHeader("ru"),
			options: []MiddlewareOption{
				SupportedLanguages(language.English, language.Russian),
				ReadFromAcceptHeader(),
			},
			expectedLanguage: language.Russian,
		},
		{
			name:    "read from accept header by default",
			request: givenRequestWithAcceptLanguageHeader("ru"),
			options: []MiddlewareOption{
				SupportedLanguages(language.English, language.Russian),
			},
			expectedLanguage: language.Russian,
		},
		{
			name:    "read from accept header by with priorities",
			request: givenRequestWithAcceptLanguageHeader("ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7"),
			options: []MiddlewareOption{
				SupportedLanguages(language.English, language.Russian),
			},
			expectedLanguage: language.Russian,
		},
		{
			name:    "read from accept header by with reversed priorities",
			request: givenRequestWithAcceptLanguageHeader("fr-BE;q=0.8, ru-RU;q=0.9, fr-BE;q=0.6"),
			options: []MiddlewareOption{
				SupportedLanguages(language.English, language.Russian),
			},
			expectedLanguage: language.Russian,
		},
		{
			name:    "read from cookie first",
			request: givenRequestWithCookieAndHeader("lang", "ru", "en"),
			options: []MiddlewareOption{
				SupportedLanguages(language.English, language.Russian),
				ReadFromCookie("lang"),
				ReadFromAcceptHeader(),
			},
			expectedLanguage: language.Russian,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			var capturedContext context.Context
			next := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
				capturedContext = request.Context()
			})
			middleware := NewMiddleware(next, test.options...)
			recorder := httptest.NewRecorder()

			middleware.ServeHTTP(recorder, test.request)
			tag := FromContext(capturedContext)

			if tag != test.expectedLanguage && tag.Parent() != test.expectedLanguage {
				t.Errorf("actual language: %s, expected language: %s", tag, test.expectedLanguage)
			}
		})
	}
}

func givenRequestWithCookie(cookieName string, value string) *http.Request {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.AddCookie(&http.Cookie{
		Name:  cookieName,
		Value: value,
	})

	return r
}

func givenRequestWithAcceptLanguageHeader(value string) *http.Request {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("Accept-Language", value)

	return r
}

func givenRequestWithCookieAndHeader(cookieName string, value string, langHeaderValue string) *http.Request {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.AddCookie(&http.Cookie{
		Name:  cookieName,
		Value: value,
	})
	r.Header.Set("Accept-Language", langHeaderValue)

	return r
}
