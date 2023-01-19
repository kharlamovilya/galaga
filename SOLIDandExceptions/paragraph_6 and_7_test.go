package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type mockRetry struct {
}

type mockRetryCommand struct {
	queue    *Queue
	cmd      ICommand
	retrying int
}

func (c *mockRetryCommand) Execute() error {
	return nil
}

// Создание экзмепляра команды

func (r mockRetry) Retry(queue *Queue, cmdData Element) ICommand {
	return &mockRetryCommand{queue: queue, cmd: cmdData.cmd, retrying: cmdData.retrying}
}

func (r mockRetry) Error() string {
	return "Command keeps failing"
}

func mockExceptionHandler7(cmdData Element, queue *Queue) error {
	return nil
}

func Test_6(t *testing.T) {
	t.Log("Given the need to test the 6th paragraph.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen common arguments.", testID)
		{
			retryer := mockRetry{}
			cmd := retryer.Retry(&Queue{}, Element{})
			err := cmd.Execute()
			require.Equal(t, nil, err)
		}
	}
}

func Test_7(t *testing.T) {
	t.Log("Given the need to test the 7th paragraph.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen common arguments.", testID)
		{
			err := mockExceptionHandler7(Element{}, &Queue{})
			require.Equal(t, nil, err)
		}
	}
}
