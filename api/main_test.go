package main_test

import (
	"fmt"

	api "github.com/Lyeed/gNetwork/api"
	cmds "github.com/Lyeed/gNetwork/commands"
)

func ExampleNewCommandMessage() {
	m := api.Message{Command: "Add", Args: []api.Data{api.Data{Name: "op1", Value: 1}, api.Data{Name: "op2", Value: 1}}}
	c := m.NewCommandMessage()
	fmt.Printf("%+v\n", c)
	// Output: Msg:<Name:"op1" Value:1 > Msg:<Name:"op2" Value:1 >
}

func ExampleSetResults() {
	m := api.Message{Command: "Add", Args: []api.Data{api.Data{Name: "op1", Value: 1}, api.Data{Name: "op2", Value: 1}}}
	r := cmds.Message{Msg: [](*cmds.Data){&cmds.Data{Name: "sum", Value: 2}}}
	m.SetResults(&r)
	fmt.Printf("%+v\n", m)
	// Output: {Command:Add Args:[{Name:op1 Value:1} {Name:op2 Value:1}] Results:[{Name:sum Value:2}]}
}
