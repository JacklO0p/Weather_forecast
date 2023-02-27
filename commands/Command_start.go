package commands

type CommandStart struct{}

func (c *CommandStart) Command() string {
	return "/start"
}
