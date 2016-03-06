package command

import (
	"strings"
)

type RenameAndGetCommand struct {
	Meta
}

func (c *RenameAndGetCommand) Run(args []string) int {
	// Write your code here

	return 0
}

func (c *RenameAndGetCommand) Synopsis() string {
	return "rename file and get package image"
}

func (c *RenameAndGetCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
