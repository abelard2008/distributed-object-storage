package main

import (
	"log"
	"net/http"
	"os"
)
import (
	"fmt"
	"./objects"
)

func main() {
	http.HandleFunc("/objects/", objects.Handler)
	fmt.Println("pengcz")
	log.Fatal(http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"),nil))

}
