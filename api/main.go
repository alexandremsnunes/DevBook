package main

import (
	"api/src/config"
	"api/src/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	config.Load()
	fmt.Println(config.BdConectionStr)
	r := router.Generate()

	fmt.Println("Escutando na porta: 5000")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))

}
