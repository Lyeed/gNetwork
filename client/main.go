package main

import (
	"bytes"
	"encoding/json"
	"fmt"
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

func PostCommand(c Command) *http.Response {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(c)
	res, err := http.Post(address, "application/json; charset=utf-8", b)
	if err != nil {
		log.Fatalf("%v", err)
	}
	return res
}

func main() {
	if len(os.Args) < 2 {
		log.Fatalf("1st parameter must be a json file")
	}

	cmds := OpenAndUnmarshallJSON(os.Args[1])
	for i := 0; i < len(cmds.Commands); i++ {
		fmt.Printf("\n\n%+v\n", cmds.Commands[i])
		r := PostCommand(cmds.Commands[i])
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		newStr := buf.String()
		fmt.Printf("%s\n", newStr)
	}
}
