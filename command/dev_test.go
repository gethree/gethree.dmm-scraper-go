package command

import (
	"testing"

	"github.com/mitchellh/cli"
)

func TestDevCommand_implement(t *testing.T) {
	var _ cli.Command = &DevCommand{}
}
