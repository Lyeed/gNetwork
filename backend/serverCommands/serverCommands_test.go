package serverCommands_test

import (
	"fmt"

	srvCmd "github.com/Lyeed/gNetwork/backend/serverCommands"
	cmd "github.com/Lyeed/gNetwork/commands"
)

// ExampleAdd: "Add" example function
func ExampleAdd() {
	resp, _ := srvCmd.NewServerCommands().Add(nil, &cmd.Message{Msg: [](*cmd.Data){&cmd.Data{Name: "op1", Value: 2}, &cmd.Data{Name: "op2", Value: 2}}})
	fmt.Printf("%+v\n", resp)
	// Output: Msg:<Name:"sum" Value:4 >
}

// ExampleSleep: "Sleep" example function
func ExampleSleep() {
	resp, _ := srvCmd.NewServerCommands().Sleep(nil, &cmd.Message{Msg: [](*cmd.Data){&cmd.Data{Name: "duration", Value: 1000}}})
	fmt.Printf("%+v\n", resp)
	// Output: Msg:<Name:"actual_duration" Value:1000 >
}
