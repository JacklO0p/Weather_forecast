package commands

type CommandNewTimer struct{}

func (c *CommandNewTimer) Command() string {
	return "/newTimer"
}