package main_test

import (
	"fmt"

	msg "github.com/Lyeed/gNetwork/api/message"
	cmds "github.com/Lyeed/gNetwork/commands"
)

// ExampleNewCommandMessage: Generates a command message for the Add function
func ExampleNewCommandMessage() {
	m := msg.Message{Command: "Add", Args: []msg.Data{msg.Data{Name: "op1", Value: 1}, msg.Data{Name: "op2", Value: 1}}}
	c := m.NewCommandMessage()
	fmt.Printf("%+v\n", c)
	// Output: Msg:<Name:"op1" Value:1 > Msg:<Name:"op2" Value:1 >
}

// ExampleSetResults: Sets the sum of the Add function to the message structure
func ExampleSetResults() {
	m := msg.Message{Command: "Add", Args: []msg.Data{msg.Data{Name: "op1", Value: 1}, msg.Data{Name: "op2", Value: 1}}}
	r := cmds.Message{Msg: [](*cmds.Data){&cmds.Data{Name: "sum", Value: 2}}}
	m.SetResults(&r)
	fmt.Printf("%+v\n", m)
	// Output: {Command:Add Args:[{Name:op1 Value:1} {Name:op2 Value:1}] Results:[{Name:sum Value:2}]}
}
