package commands

type CommandNewLocation struct{}

func (c *CommandNewLocation) Command() string {
	return "/newLocation"
}
