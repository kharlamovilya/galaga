package game

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMacroCommand_Execute(t *testing.T) {
	t.Log("Given the need to test MoveAndBurn's MoveAndBurnFuel command.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen common arguments.", testID)
		{
			commands := []ICommand{&Command{}, &Command{}}
			cmd := NewMacroCommand(commands)
			err := cmd.Execute()
			require.Equal(t, nil, err)
		}

		testID++
		t.Logf("\tTest %d:\tWhen no arguments.", testID)
		{
			var commands []ICommand
			cmd := NewMacroCommand(commands)
			err := cmd.Execute()
			require.Equal(t, nil, err)
		}

		testID++
		t.Logf("\tTest %d:\tWhen getting a wrong arguments.", testID)
		{
			commands := []ICommand{&GetErrorCommand{}}
			cmd := NewMacroCommand(commands)
			err := cmd.Execute()
			// Должны получить ошибку, т.к. в макро команду была передана команда, возвращающая ошибку
			require.NotEqual(t, nil, err)
		}
	}
}
