package main

import (
	"log"
	"net"
	"strings"
	"time"

	"github.com/Lyeed/gNetwork/commands"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":50051"

type serverCommands struct{}

func getParam(d []*commands.Data, name string) *commands.Data {
	for _, element := range d {
		if strings.Compare(element.Name, name) == 0 {
			return element
		}
	}
	return nil
}

func NewReply(name string, value int64) *commands.Message {
	return &commands.Message{Msg: [](*commands.Data){&commands.Data{Name: name, Value: value}}}
}

func (s *serverCommands) Add(ctx context.Context, in *commands.Message) (*commands.Message, error) {
	ope1 := getParam(in.Msg, "op1")
	ope2 := getParam(in.Msg, "op2")
	if ope1 == nil || ope2 == nil {
		return NewReply("wrong_syntax", -1), nil
	}
	return NewReply("sum", ope1.Value+ope2.Value), nil
}

func (s *serverCommands) Sleep(ctx context.Context, in *commands.Message) (*commands.Message, error) {
	dur := getParam(in.Msg, "duration")
	if dur == nil {
		return NewReply("wrong_syntax", -1), nil
	}
	start := time.Now()
	time.Sleep(time.Duration(dur.Value) * time.Millisecond)
	return NewReply("actual_duration", int64(time.Since(start)/time.Millisecond)), nil
}

func (s *serverCommands) Error(ctx context.Context, in *commands.Message) (*commands.Message, error) {
	return NewReply("unknown_command", -1), nil
}

func main() {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	commands.RegisterCommandsServer(s, &serverCommands{})
	reflection.Register(s)
	if err := s.Serve(ln); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
