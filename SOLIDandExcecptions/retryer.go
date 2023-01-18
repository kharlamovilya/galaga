package main

type IRetryer interface {
	Retry(*Queue, Element) ICommand
}

type Retry struct {
}

// Команда, которая повторяет Команду, выбросившую исключение

type RetryCommand struct {
	queue    *Queue
	cmd      ICommand
	retrying int
}

func (c *RetryCommand) Execute() error {
	err := c.queue.Put(c.cmd, c.retrying)
	return err
}

// Создание экзмепляра команды

func (r Retry) Retry(queue *Queue, cmdData Element) ICommand {
	return &RetryCommand{queue: queue, cmd: cmdData.cmd, retrying: cmdData.retrying}
}

func (r Retry) Error() string {
	return "Command keeps failing"
}
