package main

import "fmt"

type ICommand interface {
	Execute() error
}

type Command struct {
}

func (c *Command) Execute() error {
	return fmt.Errorf("something went wrong")
}
