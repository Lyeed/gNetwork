package serverCommands

import (
	"log"
	"strings"
	"time"

	"github.com/Lyeed/gNetwork/commands"
	"golang.org/x/net/context"
)

type ServerCommands struct{}

type IServerCommands interface {
	Add(context.Context, *commands.Message) (*commands.Message, error)
	Sleep(context.Context, *commands.Message) (*commands.Message, error)
	Error(context.Context, *commands.Message) (*commands.Message, error)
}

func NewServerCommands() IServerCommands {
	return &ServerCommands{}
}

// Add: Sums two operands
func (s *ServerCommands) Add(ctx context.Context, in *commands.Message) (*commands.Message, error) {
	ope1 := GetParam(in.Msg, "op1")
	ope2 := GetParam(in.Msg, "op2")
	if ope1 == nil || ope2 == nil {
		return NewReply("wrong_syntax", -1), nil
	}
	log.Printf("%d + %d = %d\n", ope1.Value, ope2.Value, ope1.Value+ope2.Value)
	return NewReply("sum", ope1.Value+ope2.Value), nil
}

// Sleep: Sleeps for a certain duration
func (s *ServerCommands) Sleep(ctx context.Context, in *commands.Message) (*commands.Message, error) {
	dur := GetParam(in.Msg, "duration")
	if dur == nil {
		return NewReply("wrong_syntax", -1), nil
	}
	start := time.Now()
	log.Printf("Sleeping %dms\n", dur.Value)
	time.Sleep(time.Duration(dur.Value) * time.Millisecond)
	return NewReply("actual_duration", int64(time.Since(start)/time.Millisecond)), nil
}

// Error: Returns a "Unknown command" response
func (s *ServerCommands) Error(ctx context.Context, in *commands.Message) (*commands.Message, error) {
	return NewReply("unknown_command", -1), nil
}

// GetParam: Returns the argument matching with name
func GetParam(d []*commands.Data, name string) *commands.Data {
	for _, element := range d {
		if strings.Compare(element.Name, name) == 0 {
			return element
		}
	}
	return nil
}

// NewReply: Creates and return a response with parameters name and value as content
func NewReply(name string, value int64) *commands.Message {
	return &commands.Message{Msg: [](*commands.Data){&commands.Data{Name: name, Value: value}}}
}
