package main_test

import (
	"bytes"
	"fmt"

	m "github.com/Lyeed/gNetwork/client"
)

// ExampleAddition: Sums of 1 and 1
func ExampleAddition() {
	r, err := m.PostCommand(m.Command{Command: "Add", Args: []m.Arg{m.Arg{Name: "op1", Value: 1}, m.Arg{Name: "op2", Value: 1}}})
	if err == nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		newStr := buf.String()
		fmt.Println(newStr)
	} else {
		fmt.Println(err)
	}
	// Output: {"Command":"Add","Args":[{"Name":"op1","Value":1},{"Name":"op2","Value":1}],"Results":[{"Name":"sum","Value":2}]}
}

// ExampleAdditionMissingArg: Sums of 1 and a missing argument
func ExampleAdditionMissingArg() {
	r, err := m.PostCommand(m.Command{Command: "Add", Args: []m.Arg{m.Arg{Name: "op1", Value: 1}}})
	if err == nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		newStr := buf.String()
		fmt.Println(newStr)
	} else {
		fmt.Println(err)
	}
	// Output: {"Command":"Add","Args":[{"Name":"op1","Value":1}],"Results":[{"Name":"wrong_syntax","Value":-1}]}
}

// ExampleUnknownCommand: Tests an Unknown command
func ExampleUnknownCommand() {
	r, err := m.PostCommand(m.Command{Command: "Foobar", Args: []m.Arg{}})
	if err == nil {
		buf := new(bytes.Buffer)
		buf.ReadFrom(r.Body)
		newStr := buf.String()
		fmt.Println(newStr)
	} else {
		fmt.Println(err)
	}
	// Output: {"Command":"Foobar","Args":[],"Results":[{"Name":"unknown_command","Value":-1}]}
}
