package language

import (
	"net/http"

	"golang.org/x/text/language"
)

// Middleware is used to parse language from Accept-Language header or custom cookie
// from current request and pass best matching language via context to next http.Handler.
type Middleware struct {
	matcher language.Matcher
	readers []func(r *http.Request) string
	next    http.Handler
}

// MiddlewareOption is used to set up Middleware.
type MiddlewareOption func(middleware *Middleware)

// ReadFromCookie can be used to set up middleware to read language value from cookie with given name.
func ReadFromCookie(name string) MiddlewareOption {
	return func(middleware *Middleware) {
		middleware.readers = append(middleware.readers, func(r *http.Request) string {
			cookie, _ := r.Cookie(name)
			return cookie.Value
		})
	}
}

// ReadFromAcceptHeader can be used to set up middleware to read language value from Accept-Language header.
func ReadFromAcceptHeader() MiddlewareOption {
	return func(middleware *Middleware) {
		middleware.readers = append(middleware.readers, readFromAcceptLanguageHeader)
	}
}

// SupportedLanguages is used to set up list of supported languages. See language.NewMatcher() for details.
func SupportedLanguages(tags ...Tag) MiddlewareOption {
	return func(middleware *Middleware) {
		middleware.matcher = language.NewMatcher(tags)
	}
}

// NewMiddleware creates middleware for parsing language from Accept-Language header or a cookie
// and passing its value via context.
//
// By default Middleware uses only English language and Accept-Language header as source.
//
// To set up supported languages list use SupportedLanguages option.
//
// To set up sources of language value use ReadFromCookie and ReadFromAcceptHeader options. Order of
// sources should be preserved.
func NewMiddleware(next http.Handler, options ...MiddlewareOption) *Middleware {
	middleware := &Middleware{next: next}

	for _, setOption := range options {
		setOption(middleware)
	}
	if middleware.matcher == nil {
		middleware.matcher = language.NewMatcher([]Tag{English})
	}
	if len(middleware.readers) == 0 {
		middleware.readers = append(middleware.readers, readFromAcceptLanguageHeader)
	}

	return middleware
}

func (middleware *Middleware) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	sources := make([]string, 0, len(middleware.readers))
	for _, readLanguage := range middleware.readers {
		sources = append(sources, readLanguage(request))
	}

	tag, _ := language.MatchStrings(middleware.matcher, sources...)
	ctx := WithContext(request.Context(), tag)

	middleware.next.ServeHTTP(writer, request.WithContext(ctx))
}

func readFromAcceptLanguageHeader(r *http.Request) string {
	return r.Header.Get("Accept-Language")
}
