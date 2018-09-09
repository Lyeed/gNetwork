package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const address = "http://localhost:8080/"

type Arg struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}
type Command struct {
	Command string `json:"command"`
	Args    []Arg  `json:"args"`
}
type Commands struct {
	Commands []Command `json:"commands"`
}

func OpenAndUnmarshallJSON(s string) *Commands {
	file, err := os.Open(s)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	cmds := new(Commands)
	json.Unmarshal(byteValue, &cmds)
	return cmds
}

func PostCommand(c Command) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(c)
	fmt.Printf("\n\nSending message to %s. Command: %s\n", address, c.Command)
	res, err := http.Post(address, "application/json; charset=utf-8", b)
	if err != nil {
		log.Fatalf("%v", err)
	}
	io.Copy(os.Stdout, res.Body)
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("1st parameter should be a json file")
	}

	cmds := OpenAndUnmarshallJSON(os.Args[1])
	for i := 0; i < len(cmds.Commands); i++ {
		PostCommand(cmds.Commands[i])
	}
}
