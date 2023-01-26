package game

import "fmt"

type ICommand interface {
	Execute() error
	GetName() string
}

type Command struct {
}

func (c *Command) Execute() error {
	return nil
}

func (c *Command) GetName() string {
	return "Command"
}

type GetErrorCommand struct {
}

func (c *GetErrorCommand) Execute() error {
	return fmt.Errorf("something went wrong")
}

func (c *GetErrorCommand) GetName() string {
	return "GetErrorCommand"
}
