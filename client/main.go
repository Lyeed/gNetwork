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

const address = "http://localhost:8080/command" // API address and route

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

// OpenAndUnmarshallJSON: Open the Json file and unmarshall its content
func OpenAndUnmarshallJSON(str string) *Commands {
	file, err := os.Open(str)
	if err != nil {
		log.Fatalf("%v", err)
	}
	defer file.Close()
	byteValue, _ := ioutil.ReadAll(file)
	cmds := new(Commands)
	json.Unmarshal(byteValue, &cmds)
	return cmds
}

// PostCommand: Sends the command as json to the API
func PostCommand(cmd Command) (*http.Response, error) {
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(cmd)
	return http.Post(address, "application/json", buffer)
}

func main() {
	if len(os.Args) >= 2 {
		cmds := OpenAndUnmarshallJSON(os.Args[1])
		for i := 0; i < len(cmds.Commands); i++ {
			fmt.Printf("Command: %+v\n", cmds.Commands[i]) // printing the initial structure and its content

			response, err := PostCommand(cmds.Commands[i])
			// the variable response contains the command's response
			// the json is contained in the response's body accessible via r.Body

			// Example using command's response
			if err == nil {
				buffer := new(bytes.Buffer)
				buffer.ReadFrom(response.Body)
				str := buffer.String()
				fmt.Printf("Response type: %s\n", response.Header.Get("Content-type")) // printing content type
				fmt.Printf("Response content: %s\n\n", str)                            // printing as string the json returned
				response.Body.Close()
			} else {
				fmt.Printf("%s\n\n", err.Error()) // printing Post error
			}
		}
	} else {
		fmt.Printf("1st parameter must be a json file")
	}
}
