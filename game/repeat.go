package game

import "awesomeProject/exceptions"

type IRepeater interface {
	Repeat(*exceptions.Queue, exceptions.Element) ICommand
}

type Repeat struct {
}

// Команда, которая повторяет Команду, выбросившую исключение

type RepeatCommand struct {
	queue    *exceptions.Queue
	cmd      ICommand
	retrying int
}

func (c *RepeatCommand) Execute() error {
	err := c.queue.Put(c.cmd, c.retrying)
	return err
}

func (c *RepeatCommand) GetName() string {
	return "RepeatCommand"
}

// Создание экзмепляра команды

func (r Repeat) Repeat(queue *exceptions.Queue, cmdData exceptions.Element) ICommand {
	return &RepeatCommand{queue: queue, cmd: cmdData.Cmd(), retrying: cmdData.Retrying()}
}
