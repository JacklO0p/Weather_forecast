package commands

type CommandReport struct {}

func (c *CommandReport) Command() string {
	return "/report"
}