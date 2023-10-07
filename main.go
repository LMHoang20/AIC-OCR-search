package main

import (
	"OCRsearch/constants"
	"OCRsearch/database"
	"OCRsearch/router"
	"fmt"
	"net/http"
)

func RecoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil && err != http.ErrAbortHandler {
				fmt.Println(err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func main() {
	if err := database.Instance(constants.DBType).Initialize(); err != nil {
		fmt.Println(err)
		return
	}

	r := router.Init()

	corsMiddleware := func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Add("Access-Control-Allow-Origin", "*")
			next.ServeHTTP(w, r)
		})
	}

	r.Use(corsMiddleware)

	fmt.Println("Server running on port", constants.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", constants.Port), RecoveryMiddleware(r)); err != nil {
		fmt.Println(err)
	}
}
