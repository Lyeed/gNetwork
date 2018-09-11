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

const address = "http://localhost:8080/command"

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

func PostCommand(c Command) (*http.Response, error) {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(c)
	return http.Post(address, "application/json", b)
}

func main() {
	if len(os.Args) >= 2 {
		cmds := OpenAndUnmarshallJSON(os.Args[1])
		for i := 0; i < len(cmds.Commands); i++ {
			fmt.Printf("Command: %+v\n", cmds.Commands[i]) // printing the initial structure and its content

			r, err := PostCommand(cmds.Commands[i])
			// the variable r contains the command's response
			// the json is contained in the response's body accessible via r.Body

			// Example using command's response
			if err == nil {
				buf := new(bytes.Buffer)
				buf.ReadFrom(r.Body)
				newStr := buf.String()
				fmt.Printf("Response type: %s\n", r.Header.Get("Content-type")) // printing content type
				fmt.Printf("Response content: %s\n\n", newStr)                  // printing as string the json returned
				r.Body.Close()
			} else {
				fmt.Printf("%s\n\n", err.Error()) // printing Post error
			}
		}
	} else {
		fmt.Printf("1st parameter must be a json file")
	}
}
