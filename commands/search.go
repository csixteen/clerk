package commands

import (
	u "github.com/csixteen/clerk/cmd/clerk/util"
	"github.com/csixteen/clerk/pkg/actions"
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
			results, err := actions.Search(database, args)
			if err != nil {
				//TODO - log
				return
			}

			for _, res := range results {
				u.PrintColor(res.String(), u.ColorYellow)
			}
		},
	}
}
