package main

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func mockExceptionHandler8(cmdData Element, err error, logFile *os.File, queue *Queue) error {
	return nil
}

func mockExceptionHandler9(cmdData Element, logFile *os.File, queue *Queue) error {
	return nil
}

func Test_8(t *testing.T) {
	t.Log("Given the need to test the 8th paragraph.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen common arguments.", testID)
		{
			err := mockExceptionHandler8(Element{}, nil, nil, &Queue{})
			require.Equal(t, nil, err)
		}
	}
}

func Test_9(t *testing.T) {
	t.Log("Given the need to test the 9th paragraph.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen common arguments.", testID)
		{
			err := mockExceptionHandler9(Element{}, nil, &Queue{})
			require.Equal(t, nil, err)
		}
	}
}
