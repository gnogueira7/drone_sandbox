package main

import (
  "fmt"
  "net/http"
)

func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "My Awesome Go Web App")
}

func setupRoutes() {
  http.HandleFunc("/", homePage)
}

func main() {
  fmt.Println("Go Web App Started on Port 80")
  setupRoutes()
  http.ListenAndServe(":80", nil)
}
