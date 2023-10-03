package main

import (
	"OCRsearch/constants"
	"OCRsearch/database"
	"OCRsearch/router"
	"fmt"
	"net/http"
)

func main() {
	if err := database.Instance(constants.DBType).Initialize(); err != nil {
		fmt.Println(err)
		return
	}

	r := router.Init()

	fmt.Println("Server running on port", constants.Port)
	if err := http.ListenAndServe(fmt.Sprintf(":%s", constants.Port), r); err != nil {
		fmt.Println(err)
	}
}
