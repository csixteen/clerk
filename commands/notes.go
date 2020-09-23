package commands

import (
	"strings"
	"time"

	u "github.com/csixteen/clerk/cmd/clerk/util"
	"github.com/csixteen/clerk/pkg/actions"
	"github.com/spf13/cobra"
)

// Notes returns the top level `note` command.
func Notes() *cobra.Command {
	notes := &cobra.Command{
		Use:     "note",
		Aliases: []string{"n"},
		Short:   "Manage your notes.",
		Long:    "Add, list, delete or change your existing notes.",
	}

	notes.AddCommand(listNotes())
	notes.AddCommand(addNote())
	notes.AddCommand(appendNote())
	notes.AddCommand(showNote())
	notes.AddCommand(deleteNote())

	return notes
}

func listNotes() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "Lists all the existing notes",
		Aliases: []string{"ls"},
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			notes, err := actions.ListNotes(database)
			if err != nil {
				// TODO - log
				return
			}

			for _, n := range notes {
				u.PrintColor(n.String(), u.ColorPurple)
			}
		},
	}
}

func addNote() *cobra.Command {
	return &cobra.Command{
		Use:     "add <name> <contents>...",
		Short:   "Adds a new note",
		Aliases: []string{"a"},
		Args:    cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			err := actions.AddNote(
				database,
				args[0],
				strings.Join(args[1:], " "),
				time.Now(),
			)
			if err != nil {
				//TODO - log
				return
			}
		},
	}
}

func appendNote() *cobra.Command {
	return &cobra.Command{
		Use:     "append <name-or-id> <contents>...",
		Short:   "Appends contents to an existing note",
		Long:    "Appends contents to an existing note, given its name or id. The id should be prefixed by a '#'",
		Aliases: []string{"app"},
		Args:    cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			err := actions.AppendNote(
				database,
				args[0],
				strings.Join(args[1:], " "),
			)
			if err != nil {
				//TODO - log
				return
			}
		},
	}
}

func showNote() *cobra.Command {
	return &cobra.Command{
		Use:     "show <name-or-id>",
		Short:   "Shows the contents of a note",
		Long:    "Shows the contents of a note given its name or id. The id should be prefixed by a '#'",
		Aliases: []string{"sh"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			n, err := actions.GetNote(database, args[0])
			if err != nil {
				//TODO - log
				return
			}

			u.PrintColor(n.String(), u.ColorPurple)
		},
	}
}

func deleteNote() *cobra.Command {
	return &cobra.Command{
		Use:     "del <name-or-id>",
		Short:   "Deletes an existing note",
		Long:    "Deletes an existing note given its name or id. The id should be prefixed by a '#'",
		Aliases: []string{"d"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := actions.DeleteNote(database, args[0])
			if err != nil {
				//TODO - log
				return
			}
		},
	}
}
