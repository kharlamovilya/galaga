package game

type IMovableAndBurnable interface {
	MoveAndBurnFuel(IMovable, IFuelBurnable, int) ICommand
}

type MoveAndBurn struct {
	movable  IMovable
	burnable IFuelBurnable
}

type MoveAndBurnCommand struct {
	movable         IMovable
	burnable        IFuelBurnable
	fuelConsumption int
}

func (c *MoveAndBurnCommand) Execute() error {
	cmd1 := fuelBurn{}.CheckFuel(c.burnable, c.fuelConsumption)
	cmd2 := fuelBurn{}.BurnFuel(c.burnable, c.fuelConsumption)
	// пока что команда Move не реализована, условие задания это предусматривает
	cmd3 := &Command{}
	mc := NewMacroCommand([]ICommand{cmd1, cmd2, cmd3})
	err := mc.Execute()
	return err
}

func (c *MoveAndBurnCommand) GetName() string {
	return "MoveAndBurnCommand"
}

func (c *MoveAndBurn) MoveAndBurnFuel(m IMovable, b IFuelBurnable, f int) ICommand {
	return &MoveAndBurnCommand{m, b, f}
}
