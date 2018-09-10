package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/Lyeed/gNetwork/commands"
	"google.golang.org/grpc"
)

type Data struct {
	Name  string
	Value int64
}

type Message struct {
	Command string
	Args    []Data
	Results []Data
}

const address = "localhost:50051"

type command func(commands.CommandsClient, context.Context, *commands.Message, ...grpc.CallOption) (*commands.Message, error)

func Dial(m Message) Message {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("dial error: %v", err)
	}
	defer conn.Close()

	c := commands.NewCommandsClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reply, _ := CallCommand(c, ctx, m)
	m.Results = NewRespondMessage(reply)
	return m
}

func CallCommand(c commands.CommandsClient, ctx context.Context, m Message) (*commands.Message, error) {
	cmdsMap := map[string]command{
		"Add":   commands.CommandsClient.Add,
		"Sleep": commands.CommandsClient.Sleep,
	}
	req := NewCommandMessage(m)
	cmd, found := cmdsMap[m.Command]
	if !found {
		return c.Error(ctx, &req)
	}
	return cmd(c, ctx, &req)
}

func NewRespondMessage(r *commands.Message) []Data {
	var args []Data
	for _, element := range r.Msg {
		var n Data
		n.Name = element.Name
		n.Value = element.Value
		args = append(args, n)
	}
	return args
}

func NewCommandMessage(m Message) commands.Message {
	var r commands.Message
	for _, element := range m.Args {
		var d commands.Data
		d.Name = element.Name
		d.Value = element.Value
		r.Msg = append(r.Msg, &d)
	}
	return r
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Request body needed", http.StatusBadRequest)
		return
	}

	var m Message
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
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
	http.HandleFunc("/command", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
