# Language package for Golang

[![Go Reference](https://pkg.go.dev/badge/github.com/muonsoft/language.svg)](https://pkg.go.dev/github.com/muonsoft/language)
[![Go Report Card](https://goreportcard.com/badge/github.com/muonsoft/language)](https://goreportcard.com/report/github.com/muonsoft/language)
[![CI](https://github.com/muonsoft/language/actions/workflows/main.yml/badge.svg)](https://github.com/muonsoft/language/actions/workflows/main.yml)

Package language provides HTTP middleware for parsing language from HTTP request and passing it via context.

## How to install

Run the following command to install the package:

```bash
go get -u github.com/muonsoft/language
```

## Example of reading language from Accept-Language header

```golang
package main

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    
    "github.com/muonsoft/language"
)

func main() {
    h := http.HandlerFunc(func (writer http.ResponseWriter, request *http.Request) {
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
```

## Example of reading language from Cookie

```golang
package main

import (
    "fmt"
    "net/http"
    "net/http/httptest"
    
    "github.com/muonsoft/language"
)

func main() {
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
```
