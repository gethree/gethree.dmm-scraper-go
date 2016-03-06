package main

import (
	"github.com/gethree/gethree.dmm-scraper-go/command"
	"github.com/mitchellh/cli"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"dev": func() (cli.Command, error) {
			return &command.DevCommand{
				Meta: *meta,
			}, nil
		},
		"renameAndGet": func() (cli.Command, error) {
			return &command.RenameAndGetCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
