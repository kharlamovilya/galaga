package exceptions

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
	err := c.queue.Put(c.cmd, c.retrying+1)
	return err
}

func (c *RetryCommand) GetName() string {
	return "RetryCommand"
}

// Создание экзмепляра команды

func (r Retry) Retry(queue *Queue, cmdData Element) ICommand {
	return &RetryCommand{queue: queue, cmd: cmdData.Cmd(), retrying: cmdData.Retrying()}
}

func (r Retry) Error() string {
	return "Command keeps failing"
}
