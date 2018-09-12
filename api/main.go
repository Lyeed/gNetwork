package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	msg "github.com/Lyeed/gNetwork/api/message"
	"github.com/Lyeed/gNetwork/commands"
	"google.golang.org/grpc"
)

const address = "localhost:50051" // commands service address and port
const port = ":8080"              // API listen port

type commandCalls func(commands.CommandsClient, context.Context, *commands.Message, ...grpc.CallOption) (*commands.Message, error)

// CallCommand: Searches the command in the map and execute it
// Calls Error in case the command called does not exist
// Returns the command response
func CallCommand(ctx context.Context, cmds commands.CommandsClient, msg msg.IMessage) (*commands.Message, error) {
	cmdsMap := map[string]commandCalls{
		"Add":   commands.CommandsClient.Add,
		"Sleep": commands.CommandsClient.Sleep,
	}
	cmd, found := cmdsMap[msg.GetCommand()]
	if !found {
		return cmds.Error(ctx, msg.NewCommandMessage())
	}
	return cmd(cmds, ctx, msg.NewCommandMessage())
}

// Handler: Handle the client request
// Checks the request, execute its command and send the response to the client
func Handler(writer http.ResponseWriter, req *http.Request) {
	if req.Body == nil {
		http.Error(writer, "Request body needed", http.StatusBadRequest)
		return
	}

	if strings.Compare(req.Header.Get("Content-type"), "application/json") != 0 {
		http.Error(writer, "HTTP Content-Type must be application/json", http.StatusBadRequest)
		return
	}

	mess, err := msg.NewMessage(req)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	log.Printf("Message: %+v\n", mess)

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	cmdClient := commands.NewCommandsClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reply, _ := CallCommand(ctx, cmdClient, mess)
	mess.SetResults(reply)
	log.Printf("Command response: %+v\n", mess.GetResults())
	log.Printf("API response: %+v\n\n", mess)

	js, err := json.Marshal(mess)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}

	writer.Header().Set("Content-Type", "application/json")
	writer.Write(js)
}

func main() {
	http.HandleFunc("/command", Handler)
	log.Fatal(http.ListenAndServe(port, nil))
}
