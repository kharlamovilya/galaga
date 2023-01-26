package game

type macroCommand struct {
	Command
	commands []ICommand
}

func NewMacroCommand(commands []ICommand) *macroCommand {
	return &macroCommand{commands: commands}
}

func (mc macroCommand) Execute() error {
	for _, cmd := range mc.commands {
		err := cmd.Execute()
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *macroCommand) GetName() string {
	return "macroCommand"
}
