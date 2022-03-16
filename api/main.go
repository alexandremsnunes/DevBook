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

	fmt.Printf("Escutando na porta: %d", config.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Port), r))

}
