package main

import (
	"OCRsearch/router"
	"fmt"
	"net/http"
)

const PORT string = "3003"

func main() {
	r := router.Init()

	fmt.Println("Server running on port", PORT)
	http.ListenAndServe(fmt.Sprintf(":%s", PORT), r)
}
