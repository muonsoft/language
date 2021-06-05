package language_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/muonsoft/language"
	textlanguage "golang.org/x/text/language"
)

func ExampleEqual() {
	fmt.Println(language.Equal(language.English, language.English))
	fmt.Println(language.Equal(language.English, language.Russian))
	fmt.Println(language.Equal(textlanguage.MustParse("ru"), textlanguage.MustParse("ru-RU")))
	// Output:
	// true
	// false
	// true
}

func ExampleMiddleware_ServeHTTP_readFromAcceptLanguageHeader() {
	h := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		tag := language.FromContext(request.Context())
		fmt.Println("language:", tag)
	})
	m := language.NewMiddleware(h, language.SupportedLanguages(language.English, language.Russian))

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.Header.Set("Accept-Language", "ru")
	w := httptest.NewRecorder()

	m.ServeHTTP(w, r)
	// Output: language: ru
}

func ExampleMiddleware_ServeHTTP_readFromCookie() {
	h := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		tag := language.FromContext(request.Context())
		fmt.Println("language:", tag)
	})
	m := language.NewMiddleware(
		h,
		language.SupportedLanguages(language.English, language.Russian),
		language.ReadFromCookie("lang"),
	)

	r := httptest.NewRequest(http.MethodGet, "/", nil)
	r.AddCookie(&http.Cookie{Name: "lang", Value: "ru"})
	w := httptest.NewRecorder()

	m.ServeHTTP(w, r)
	// Output: language: ru
}
