package commands

type CommandHelp struct{}

func (c *CommandHelp) Command() string {
	return "/help"
}