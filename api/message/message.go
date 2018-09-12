package message

import (
	"encoding/json"
	"net/http"

	"github.com/Lyeed/gNetwork/commands"
)

type Data struct {
	Name  string
	Value int64
}

type Message struct {
	Command string
	Args    []Data
	Results []Data
}

// NewMessage: Creates and return a structure decoded from the body request
func NewMessage(req *http.Request) (IMessage, error) {
	var msg Message
	err := json.NewDecoder(req.Body).Decode(&msg)
	return &msg, err
}

type IMessage interface {
	GetCommand() string
	NewCommandMessage() *commands.Message
	SetResults(*commands.Message)
}

// NewCommandMessage: Create and return a new message structured according to the commands server's need
func (msg Message) NewCommandMessage() *commands.Message {
	var cmdMsg commands.Message
	for _, element := range msg.Args {
		var data commands.Data
		data.Name = element.Name
		data.Value = element.Value
		cmdMsg.Msg = append(cmdMsg.Msg, &data)
	}
	return &cmdMsg
}

// SetResults: Set the message's result from the command execution result
func (msg *Message) SetResults(response *commands.Message) {
	for _, element := range response.Msg {
		var data Data
		data.Name = element.Name
		data.Value = element.Value
		msg.Results = append(msg.Results, data)
	}
}

func (msg *Message) GetCommand() string {
	return msg.Command
}
