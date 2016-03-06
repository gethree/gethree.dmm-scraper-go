package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestRenameAndGetCommand_implement(t *testing.T) {
	var _ cli.Command = &RenameAndGetCommand{}
}
