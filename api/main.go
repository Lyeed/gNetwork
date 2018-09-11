package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/Lyeed/gNetwork/commands"
	"google.golang.org/grpc"
)

const address = "localhost:50051"
const port = ":8080"

type Data struct {
	Name  string
	Value int64
}

type Message struct {
	Command string
	Args    []Data
	Results []Data
}

func (m Message) NewCommandMessage() *commands.Message {
	var r commands.Message
	for _, element := range m.Args {
		var d commands.Data
		d.Name = element.Name
		d.Value = element.Value
		r.Msg = append(r.Msg, &d)
	}
	return &r
}

func (m *Message) SetResults(r *commands.Message) {
	for _, element := range r.Msg {
		var n Data
		n.Name = element.Name
		n.Value = element.Value
		m.Results = append(m.Results, n)
	}
}

type command func(commands.CommandsClient, context.Context, *commands.Message, ...grpc.CallOption) (*commands.Message, error)

func CallCommand(c commands.CommandsClient, ctx context.Context, m Message) (*commands.Message, error) {
	cmdsMap := map[string]command{
		"Add":   commands.CommandsClient.Add,
		"Sleep": commands.CommandsClient.Sleep,
	}
	cmd, found := cmdsMap[m.Command]
	if !found {
		return c.Error(ctx, m.NewCommandMessage())
	}
	return cmd(c, ctx, m.NewCommandMessage())
}

func Handler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Request body needed", http.StatusBadRequest)
		return
	}

	if strings.Compare(r.Header.Get("Content-type"), "application/json") != 0 {
		http.Error(w, "HTTP Content-Type must be application/json", http.StatusBadRequest)
		return
	}

	var m Message
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()

	c := commands.NewCommandsClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	reply, _ := CallCommand(c, ctx, m)
	m.SetResults(reply)

	js, err := json.Marshal(m)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func main() {
	http.HandleFunc("/command", Handler)
	log.Fatal(http.ListenAndServe(port, nil))
}
