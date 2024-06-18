package game

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type mockRepeat struct {
}

type mockRepeatCommand struct {
}

func (c *mockRepeatCommand) Execute() error {
	return nil
}

func (c *mockRepeatCommand) GetName() string {
	return "mockRepeatCommand"
}

func (r mockRepeat) mockRepeat() ICommand {
	return &mockRepeatCommand{}
}

func TestRepeat_Repeat(t *testing.T) {
	t.Log("Given the need to test Repeat's Repeat command.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen common arguments.", testID)
		{
			obj := mockRepeat{}
			cmd := obj.mockRepeat()
			err := cmd.Execute()
			require.Equal(t, nil, err)
		}
	}
}
