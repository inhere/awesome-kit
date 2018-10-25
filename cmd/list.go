package cmd

import "github.com/gookit/cliapp"

func ListCommand() *cliapp.Command {
	c := &cliapp.Command{
		Name: "list",
		UseFor: "list all contains awesome repositories",
	}



	return c
}