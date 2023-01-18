package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

type mockLogger struct {
}

func (l mockLogger) WriteLog(logFile *os.File, log string, err error) ICommand {
	return &mockWriteLogCommand{file: logFile, log: log, err: err}
}

type mockWriteLogCommand struct {
	file *os.File
	log  string
	err  error
}

func (c *mockWriteLogCommand) Execute() error {
	return nil
}

func mockExceptionHandler5(cmdData Element, err error, logFile *os.File, queue *Queue) error {
	// do something
	return nil
}

func Test_4(t *testing.T) {
	t.Log("Given the need to test the 4th paragraph.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen common arguments.", testID)
		{
			logger := mockLogger{}
			cmd := logger.WriteLog(&os.File{}, "", nil)
			err := cmd.Execute()
			require.Equal(t, nil, err)
		}
	}
}

func Test_5(t *testing.T) {
	t.Log("Given the need to test the 5th paragraph.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen common arguments.", testID)
		{
			err := mockExceptionHandler5(Element{}, nil, nil, &Queue{})
			require.Equal(t, nil, err)
		}
	}
}
