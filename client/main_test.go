package main_test

import (
	"bytes"
	"fmt"
	"testing"

	m "github.com/Lyeed/gNetwork/client"
)

func TestAddition(t *testing.T) {
	r, _ := m.PostCommand(m.Command{Command: "Add", Args: []m.Arg{m.Arg{Name: "op1", Value: 1}, m.Arg{Name: "op2", Value: 1}}})
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()
	fmt.Println(newStr)
	// Output: {"Command":"Add","Args":[{"Name":"op1","Value":1},{"Name":"op2","Value":1}],"Results":[{"Name":"sum","Value":2}]}
}

func TestAdditionMissingArg(t *testing.T) {
	r, _ := m.PostCommand(m.Command{Command: "Add", Args: []m.Arg{m.Arg{Name: "op1", Value: 1}}})
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()
	fmt.Println(newStr)
	// Output: {"Command":"Add","Args":[{"Name":"op1","Value":1}],"Results":[{"Name":"wrong syntax","Value":-1}]}
}

func TestAdditionNoArg(t *testing.T) {
	r, _ := m.PostCommand(m.Command{Command: "Foobar", Args: []m.Arg{}})
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()
	fmt.Println(newStr)
	// Output: {"Command":"Foobar","Args":[],"Results":[{"Name":"unknown command","Value":-1}]}
}
