package main

type SecondaryRetry struct {
}

// Команда, которая повторяет Команду, выбросившую исключение

type SecondaryRetryCommand struct {
	queue    *Queue
	cmd      ICommand
	retrying int
}

func (c *SecondaryRetryCommand) Execute() error {
	err := c.queue.Put(c.cmd, c.retrying)
	return err
}

// Создание экзмепляра команды

func (r SecondaryRetry) Retry(queue *Queue, cmdData Element) ICommand {
	return &RetryCommand{queue: queue, cmd: cmdData.cmd, retrying: cmdData.retrying}
}

func (r SecondaryRetry) Error() string {
	return "Command keeps failing"
}
