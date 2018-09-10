package main

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	cmds "github.com/Lyeed/gNetwork/commands"
	"google.golang.org/grpc"
)

type Argument struct {
	Name  string
	Value int
}

type Result struct {
	Name  string
	Value int64
}

type Message struct {
	Command string
	Args    []Argument
}

type Respond struct {
	Command string
	Args    []Argument
	Results []Result
}

const address = "localhost:50051"

type command func(cmds.CommandsClient, context.Context, *cmds.Request, ...grpc.CallOption) (*cmds.Reply, error)

func Dial(m Message) Respond {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial error: %v", err)
	}
	defer conn.Close()

	c := cmds.NewCommandsClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reply, error := CallCommand(c, ctx, m)
	if error != nil {
		return StructureRespond(m, &cmds.Reply{Msg: [](*cmds.Data){&cmds.Data{Name: "unknown command", Value: -1}}})
	}
	return StructureRespond(m, reply)
}

func CallCommand(c cmds.CommandsClient, ctx context.Context, m Message) (*cmds.Reply, error) {
	cmdsMap := map[string]command{
		"Add":   cmds.CommandsClient.Add,
		"Sleep": cmds.CommandsClient.Sleep,
	}
	req := NewRequest(m)
	cmd, found := cmdsMap[m.Command]
	if !found {
		return nil, errors.New("commands: unknown command")
	}
	return cmd(c, ctx, &req)
}

func StructureRespond(m Message, reply *cmds.Reply) Respond {
	var respond Respond
	respond.Command = m.Command
	respond.Args = m.Args
	for _, element := range reply.Msg {
		var n Result
		n.Name = element.Name
		n.Value = element.Value
		respond.Results = append(respond.Results, n)
	}
	return respond
}

func NewRequest(m Message) cmds.Request {
	var r cmds.Request
	for _, element := range m.Args {
		var d cmds.Data
		d.Name = element.Name
		d.Value = int64(element.Value)
		r.Msg = append(r.Msg, &d)
	}
	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Request body needed", 400)
		return
	}

	var m Message
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	result := Dial(m)
	js, err := json.Marshal(result)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
