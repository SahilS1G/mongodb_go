package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/SahilS1G/mongodb_go/router"
)

func main() {
	fmt.Println("hello")
	r := router.Router()
	fmt.Println("server is getting started ...")
	log.Fatal(http.ListenAndServe(":3000", r))
	fmt.Println("listening at port ...")

}
