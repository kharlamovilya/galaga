package exceptions

import "fmt"

type ICommand interface {
	Execute() error
	GetName() string
}

type Command struct {
}

func (c *Command) Execute() error {
	return fmt.Errorf("something went wrong")
}

func (c *Command) GetName() string {
	return "Command"
}
