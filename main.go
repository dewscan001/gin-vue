package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe(os.Getenv("PORT"), http.FileServer(http.Dir("./")))
	if err != nil {
		fmt.Println("Failed to start server", err)
		return
	}
}
