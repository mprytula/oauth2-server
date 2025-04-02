package main

import (
	"fmt"
	"log"
	"net/http"
	"qweinke/oauth/internal/handlers"
)

const ADRESS = "127.0.0.1"
const PORT = "8000"
const SERVER_ADRESS = ADRESS + ":" + PORT

func main() {
	fmt.Println("[ START ] Server started listening at " + SERVER_ADRESS)
	var mutex = http.NewServeMux()
	handlers.AttachRouting(mutex)
	log.Fatal(http.ListenAndServe(SERVER_ADRESS, mutex))
}