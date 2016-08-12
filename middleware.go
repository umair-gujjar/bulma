// Copyright 2016 The Nanoninja Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package bulma

import "net/http"

// Middleware is a function that wraps a middleware to
// be compatible with a handler of middleware.
//    mux := http.NewServeMux()
//    mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
//        fmt.Fprintf(rw, "Welcome to the home page!")
//    })
//
//    mw := bulma.Middleware("username", "password", bulma.DefaultRealm)
//
//    http.ListenAndServe(":3000", mw(mux))
func Middleware(user, pass, realm string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return BasicAuthHandler(BasicAuth(user, pass), next, realm)
	}
}

// MiddlewareFunc is the same as Middleware but
// allows to use a custom authentication BasicAuthFunc function.
//    mux := http.NewServeMux()
//    mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
//        fmt.Fprintf(rw, "Welcome to the home page!")
//    })
//
//    f := func(user, pass) bool {
//        return user == "username" && pass == "password"
//    }
//
//    mw := bulma.MiddlewareFunc(f, bulma.DefaultRealm)
//
//    http.ListenAndServe(":3000", mw(mux))
func MiddlewareFunc(a Authentifier, realm string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return BasicAuthHandler(a, next, realm)
	}
}
