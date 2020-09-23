package commands

import (
	"strings"
	"time"

	u "github.com/csixteen/clerk/cmd/clerk/util"
	"github.com/csixteen/clerk/pkg/actions"
	"github.com/spf13/cobra"
)

// Notes returns the top level `task` command.
func Tasks() *cobra.Command {
	notes := &cobra.Command{
		Use:     "task",
		Aliases: []string{"t"},
		Short:   "Manage your tasks",
		Long:    "Add, list, delete, change or mark tasks as completed.",
	}

	notes.AddCommand(listTasks())
	notes.AddCommand(addTask())
	notes.AddCommand(editTask())
	notes.AddCommand(deleteTask())
	notes.AddCommand(completeTask())

	return notes
}

func listTasks() *cobra.Command {
	return &cobra.Command{
		Use:     "list",
		Short:   "Lists all the existing tasks",
		Aliases: []string{"ls"},
		Args:    cobra.NoArgs,
		Run: func(cmd *cobra.Command, args []string) {
			tasks, err := actions.ListTasks(database)
			if err != nil {
				// TODO - log
				return
			}

			for _, t := range tasks {
				u.PrintColor(t.String(), u.ColorPurple)
			}
		},
	}
}

func addTask() *cobra.Command {
	return &cobra.Command{
		Use:     "add <name> <contents>...",
		Short:   "Adds a new task",
		Aliases: []string{"a"},
		Args:    cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			err := actions.AddTask(
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

func editTask() *cobra.Command {
	return &cobra.Command{
		Use:     "edit <name-or-id> <new contents>...",
		Short:   "Replace the contents of a task",
		Long:    "Replaces the contents of an existing task given its name or id. The id should be prefixed by a '#'",
		Aliases: []string{"e"},
		Args:    cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			err := actions.EditTask(
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

func deleteTask() *cobra.Command {
	return &cobra.Command{
		Use:     "del <name-or-id>",
		Short:   "Deletes an existing task",
		Long:    "Deletes an existing task given its name or id. The id should be prefixed by a '#'",
		Aliases: []string{"d"},
		Args:    cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := actions.DeleteTask(database, args[0])
			if err != nil {
				//TODO - log
				return
			}
		},
	}
}

func completeTask() *cobra.Command {
	return &cobra.Command{
		Use:   "done <name-or-id>",
		Short: "Marks an existing task as completed",
		Long:  "Marks an existing task as completed given its name or id. The id should be prefixed by a '#'",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			err := actions.CompleteTask(database, args[0], time.Now())
			if err != nil {
				//TODO - log
				return
			}
		},
	}
}
