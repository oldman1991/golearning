package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func validAuth(s string) error {
	//if s != "123456" {
	//	return fmt.Errorf("%s", "bad auth token")
	//}
	return nil
}

func greetings(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome")
}

func logHandler(h http.Handler) http.Handler {

	tmp_func := func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t)
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(tmp_func)
	//return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	//	t := time.Now()
	//	log.Printf("[%s] %q %v\n", r.Method, r.URL.String(), t)
	//	h.ServeHTTP(w, r)
	//})
}

func authHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err := validAuth(r.Header.Get("Authorization"))
		if err != nil {
			http.Error(w, "bad auth param", http.StatusUnauthorized)
			return
		}
		h.ServeHTTP(w, r)
	})
}

func main() {

	http.ListenAndServe(":8000",
		logHandler(authHandler(http.HandlerFunc(greetings))))
}
