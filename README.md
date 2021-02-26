# Language package for Golang

Package language provides HTTP middleware for parsing language from HTTP request and passing it via context.

## Example of reading language from Accept-Language header

```golang
package main

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    
    "golang.org/x/text/language"
    
    languagepkg "github.com/muonsoft/language"
)

func main() {
    h := http.HandlerFunc(func (writer http.ResponseWriter, request *http.Request) {
        tag := languagepkg.FromContext(request.Context())
        fmt.Println("language:", tag)
    })
    m := languagepkg.NewMiddleware(h, languagepkg.SupportedLanguages(language.English, language.Russian))
    
    r := httptest.NewRequest(http.MethodGet, "/", nil)
    r.Header.Set("Accept-Language", "ru")
    w := httptest.NewRecorder()
    
    m.ServeHTTP(w, r)
    // Output: language: ru
}
```

## Example of reading language from Cookie

```golang
package main

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    
    "golang.org/x/text/language"
    
    languagepkg "github.com/muonsoft/language"
)

func main() {
    h := http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
        tag := languagepkg.FromContext(request.Context())
        fmt.Println("language:", tag)
    })
    m := language.NewMiddleware(
        h,
        languagepkg.SupportedLanguages(language.English, language.Russian),
        languagepkg.ReadFromCookie("lang"),
    )
    
    r := httptest.NewRequest(http.MethodGet, "/", nil)
    r.AddCookie(&http.Cookie{Name: "lang", Value: "ru"})
    w := httptest.NewRecorder()
    
    m.ServeHTTP(w, r)
    // Output: language: ru
}
```
