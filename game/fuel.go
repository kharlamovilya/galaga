package game

import "fmt"

type IFuelBurnable interface {
	Fuel() (int, error)
	SetFuel(int) error
}

type fuelBurn struct {
	fuelBurnable IFuelBurnable
}

func NewFuelBurn(fuelBurnable IFuelBurnable) *fuelBurn {
	return &fuelBurn{fuelBurnable: fuelBurnable}
}

type BurnFuelCommand struct {
	fuelBurnable    IFuelBurnable
	fuelConsumption int
}

// Execute обязательно обработать ошибку!
func (c *BurnFuelCommand) Execute() error {
	fuel, _ := c.fuelBurnable.Fuel()
	newFuelLevel := fuel - c.fuelConsumption
	err := c.fuelBurnable.SetFuel(newFuelLevel)
	return err
}

func (c *BurnFuelCommand) GetName() string {
	return "BurnFuelCommand"
}

// BurnFuel передаваемые параметры: объект-потребитель топлива, кол-во единиц потребляемого топлива
func (f fuelBurn) BurnFuel(fuelBurnable IFuelBurnable, fuelConsumption int) ICommand {
	return &BurnFuelCommand{fuelBurnable: fuelBurnable, fuelConsumption: fuelConsumption}
}

type CheckFuelCommand struct {
	fuelBurnable    IFuelBurnable
	fuelConsumption int
}

func (c *CheckFuelCommand) Execute() error {
	fuel, _ := c.fuelBurnable.Fuel()
	newFuelLevel := fuel - c.fuelConsumption
	if newFuelLevel < 0 {
		return fmt.Errorf(
			"not enough fuel, current spaceship's fuel level: %d, reqiured: %d",
			fuel,
			c.fuelConsumption,
		)
	}
	return nil
}

func (c *CheckFuelCommand) GetName() string {
	return "CheckFuelCommand"
}

// CheckFuel передаваемые параметры: объект-потребитель топлива, кол-во единиц потребляемого топлива
func (f fuelBurn) CheckFuel(fuelBurnable IFuelBurnable, fuelConsumption int) ICommand {
	return &CheckFuelCommand{fuelBurnable: fuelBurnable, fuelConsumption: fuelConsumption}
}
