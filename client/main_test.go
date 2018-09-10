package main_test

import (
	"bytes"
	"fmt"
	"testing"

	m "github.com/Lyeed/gNetwork/client"
)

func TestAddition(t *testing.T) {
	r := m.PostCommand(m.Command{Command: "Add", Args: []m.Arg{m.Arg{Name: "op1", Value: 1}, m.Arg{Name: "op2", Value: 1}}})
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()
	fmt.Printf("%s\n", newStr)
	// Output:{"Command":"Add","Args":[{"Name":"op1","Value":1},{"Name":"op2","Value":1}],"Results":[{"Name":"sum","Value":2}]}
}

func TestAdditionMissingArg(t *testing.T) {
	r := m.PostCommand(m.Command{Command: "Add", Args: []m.Arg{m.Arg{Name: "op1", Value: 1}}})
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()
	fmt.Printf("%s\n", newStr)
	// Output: {"Command":"Add","Args":[{"Name":"op1","Value":1},{"Name":"op2","Value":1}],"Results":[{"Name":"sum","Value":2}]}
}

func TestAdditionNoArg(t *testing.T) {
	r := m.PostCommand(m.Command{Command: "Add"})
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	newStr := buf.String()
	fmt.Printf("%s\n", newStr)
	// Output: {"Command":"Add","Args":[{"Name":"op1","Value":1},{"Name":"op2","Value":1}],"Results":[{"Name":"sum","Value":2}]}
}
