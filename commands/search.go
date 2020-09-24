package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

func Search() *cobra.Command {
	return &cobra.Command{
		Use:     "search <query string...>",
		Short:   "Search against all your notes and tasks",
		Long:    "Quickly retrieve any notes and tasks that contain your search string",
		Aliases: []string{"s"},
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(args)
		},
	}
}
