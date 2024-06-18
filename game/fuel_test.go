package game

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type mockfuelBurn struct {
}

type mockBurnFuelCommand struct {
}

// Execute обязательно обработать ошибку!
func (c *mockBurnFuelCommand) Execute() error {
	return nil
}

func (c *mockBurnFuelCommand) GetName() string {
	return "mockBurnFuelCommand"
}

// BurnFuel передаваемые параметры: объект-потребитель топлива, кол-во единиц потребляемого топлива
func (f mockfuelBurn) mockBurnFuel() ICommand {
	return &mockBurnFuelCommand{}
}

type mockCheckFuelCommand struct {
}

func (c *mockCheckFuelCommand) Execute() error {
	return nil
}

func (c *mockCheckFuelCommand) GetName() string {
	return "mockCheckFuelCommand"
}

// CheckFuel передаваемые параметры: объект-потребитель топлива, кол-во единиц потребляемого топлива
func (f mockfuelBurn) mockCheckFuel() ICommand {
	return &mockCheckFuelCommand{}
}

func TestFuelBurn_CheckFuel(t *testing.T) {
	t.Log("Given the need to test FuelBurn's CheckFuel command.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen common arguments.", testID)
		{
			obj := mockfuelBurn{}
			cmd := obj.mockCheckFuel()
			err := cmd.Execute()
			require.Equal(t, nil, err)
		}
	}
}

func TestFuelBurn_BurnFuel(t *testing.T) {
	t.Log("Given the need to test FuelBurn's CheckFuel command.")
	{
		testID := 0
		t.Logf("\tTest %d:\tWhen common arguments.", testID)
		{
			obj := mockfuelBurn{}
			cmd := obj.mockBurnFuel()
			err := cmd.Execute()
			require.Equal(t, nil, err)
		}
	}
}
