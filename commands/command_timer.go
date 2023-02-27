package commands

type CommandTimer struct{}

func (c *CommandTimer) Command() string {
	return "/timer"
}
