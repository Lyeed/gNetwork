package main

import (
	"log"
	"net"

	srvCmd "github.com/Lyeed/gNetwork/backend/serverCommands"
	"github.com/Lyeed/gNetwork/commands"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const port = ":50051" // commands server listen port

func main() {
	ln, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	serv := grpc.NewServer()
	commands.RegisterCommandsServer(serv, srvCmd.NewServerCommands())
	reflection.Register(serv)
	if serv.Serve(ln) != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
