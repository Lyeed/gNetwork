package main

import (
	"log"
	"net"
	"time"

	cmds "github.com/Lyeed/gNetwork/commands"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":50051"

type server struct{}

func (s *server) Add(ctx context.Context, in *cmds.Request) (*cmds.Reply, error) {
	total := in.Value[0] + in.Value[1]
	log.Printf("%d + %d = %d", in.Value[0], in.Value[1], total)
	return &cmds.Reply{Msg: [](*cmds.Data){&cmds.Data{Name: "sum", Value: total}}}, nil
}

func (s *server) Sleep(ctx context.Context, in *cmds.Request) (*cmds.Reply, error) {
	log.Printf("Sleep %d", in.Value[0])
	time.Sleep(time.Duration(in.Value[0]) * time.Millisecond)
	log.Printf("Sleep done")
	return &cmds.Reply{Msg: [](*cmds.Data){&cmds.Data{Name: "actual_duration", Value: in.Value[0]}}}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	cmds.RegisterCommandsServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
