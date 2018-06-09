package main

import (
	"net/http"

	"github.com/parakeety/omikuji/kuji"
)

func main() {
	http.Handle("/", kuji.NewOmikujiHandler())
	http.ListenAndServe(":8080", nil)
}
