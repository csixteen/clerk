// MIT License
//
// Copyright (c) 2020 Pedro Rodrigues
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package commands

import (
	"strings"
	"time"

	u "github.com/csixteen/clerk/cmd/clerk/util"
	"github.com/csixteen/clerk/pkg/models"
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
		RunE: func(cmd *cobra.Command, args []string) error {
			tasks, err := models.ListTasks(database)
			if err != nil {
				return err
			}

			for _, t := range tasks {
				u.PrintColor(t.String(), u.ColorYellow)
			}

			return nil
		},
	}
}

func addTask() *cobra.Command {
	return &cobra.Command{
		Use:     "add <name> <contents>...",
		Short:   "Adds a new task",
		Aliases: []string{"a"},
		Args:    cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := models.AddTask(
				database,
				args[0],
				strings.Join(args[1:], " "),
				time.Now(),
			)

			return err
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
		RunE: func(cmd *cobra.Command, args []string) error {
			err := models.EditTask(
				database,
				args[0],
				strings.Join(args[1:], " "),
			)

			return err
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
		RunE: func(cmd *cobra.Command, args []string) error {
			return models.DeleteTask(database, args[0])
		},
	}
}

func completeTask() *cobra.Command {
	return &cobra.Command{
		Use:   "done <name-or-id>",
		Short: "Marks an existing task as completed",
		Long:  "Marks an existing task as completed given its name or id. The id should be prefixed by a '#'",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return models.CompleteTask(database, args[0], time.Now())
		},
	}
}
