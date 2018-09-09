package main

import (
	"context"
	"encoding/json"
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
	Value int
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

func Dial(m Message) Respond {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := cmds.NewCommandsClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var result *cmds.Reply
	switch m.Command {
	case "Add":
		{
			result, err = c.Add(ctx, &cmds.Request{Value: []int64{int64(m.Args[0].Value), int64(m.Args[1].Value)}})
		}
	case "Sleep":
		{
			result, err = c.Sleep(ctx, &cmds.Request{Value: []int64{int64(m.Args[0].Value)}})
		}
	}

	if err != nil {
		log.Fatalf("Err: %v", err)
	}

	log.Printf("Command executed : %s", m.Command)

	return Respond{Command: m.Command, Args: m.Args, Results: []Result{Result{Name: result.Msg[0].Name, Value: int(result.Msg[0].Value)}}}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "Please send a request body", 400)
		return
	}
	var m Message
	err := json.NewDecoder(r.Body).Decode(&m)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	log.Printf("Command received: %s | Args count: %d", m.Command, len(m.Args))
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
