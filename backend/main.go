package main

import (
	"log"
	"net"
	"strings"
	"time"

	cmds "github.com/Lyeed/gNetwork/commands"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":50051"

type serverCommands struct{}

func getParam(d []*cmds.Data, name string) *cmds.Data {
	for _, element := range d {
		if strings.Compare(element.Name, name) == 0 {
			return element
		}
	}
	return nil
}

func NewReply(name string, value int64) *cmds.Reply {
	return &cmds.Reply{Msg: [](*cmds.Data){&cmds.Data{Name: name, Value: value}}}
}

func (s *serverCommands) Add(ctx context.Context, in *cmds.Request) (*cmds.Reply, error) {
	ope1 := getParam(in.Msg, "op1")
	ope2 := getParam(in.Msg, "op2")
	if ope1 == nil || ope2 == nil {
		return NewReply("wrong syntax", -1), nil
	}
	return NewReply("sum", ope1.Value+ope2.Value), nil
}

func (s *serverCommands) Sleep(ctx context.Context, in *cmds.Request) (*cmds.Reply, error) {
	dur := getParam(in.Msg, "duration")
	if dur == nil {
		return NewReply("wrong syntax", -1), nil
	}
	start := time.Now()
	time.Sleep(time.Duration(dur.Value) * time.Millisecond)
	return NewReply("actual_duration", int64(time.Since(start)/time.Millisecond)), nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	cmds.RegisterCommandsServer(s, &serverCommands{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
