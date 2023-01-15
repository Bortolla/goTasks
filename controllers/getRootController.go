package controllers

import (
	"fmt"
	"io"
	"net/http"
)


func GetRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("got / request\n")
	io.WriteString(w, "This is my website!\n")
}