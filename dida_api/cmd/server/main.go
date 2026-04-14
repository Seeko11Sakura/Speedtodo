package main

import (
	"log"

	apihttp "dida_api/internal/http"
)

func main() {
	router := apihttp.NewRouter()
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
