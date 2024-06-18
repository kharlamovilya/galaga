package game

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type mockMoveAndBurn struct {
}

type mockMoveAndBurnCommand struct {
}

func (c *mockMoveAndBurnCommand) Execute() error {
	return nil
}

func (c *mockMoveAndBurnCommand) GetName() string {
	return "mockMoveAndBurnCommand"
}

func (c *mockMoveAndBurn) mockMoveAndBurnFuel() ICommand {
	return &mockMoveAndBurnCommand{}
}

func TestMoveAndBurn_MoveAndBurnFuel(t *testing.T) {
	t.Log("Given the need to test MoveAndBurn's MoveAndBurnFuel command.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen common arguments.", testID)
		{
			obj := mockMoveAndBurn{}
			cmd := obj.mockMoveAndBurnFuel()
			err := cmd.Execute()
			require.Equal(t, nil, err)
		}
	}
}
