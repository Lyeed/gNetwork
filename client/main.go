package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type arg struct {
	Name  string
	Value int
}

type message struct {
	Command string
	//	Args    []arg
}

func main() {
	msg := message{Command: "add"}
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(msg)
	res, _ := http.Post("http://localhost:8080/", "application/json; charset=utf-8", b)
	io.Copy(os.Stdout, res.Body)
}
