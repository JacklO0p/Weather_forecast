package commands

type CommandLocation struct{}

func (c *CommandLocation) Commad() string {
	return "/location"
}

