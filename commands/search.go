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
	"fmt"
	"log"
	"strings"

	u "github.com/csixteen/clerk/cmd/clerk/util"
	"github.com/csixteen/clerk/pkg/actions"
	"github.com/spf13/cobra"
)

func highlightText(s string, query string) string {
	return strings.ReplaceAll(
		s,
		query,
		string(u.ColorRed)+query+string(u.ColorReset),
	)
}

func Search() *cobra.Command {
	return &cobra.Command{
		Use:     "search <query string...>",
		Short:   "Search against all your notes and tasks",
		Long:    "Quickly retrieve any notes and tasks that contain your search string",
		Aliases: []string{"s"},
		Args:    cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			query := strings.Join(args, " ")
			results, err := actions.Search(database, query)
			if err != nil {
				log.Fatal(err)
			}

			for _, res := range results {
				u.PrintColor(res.Type(), u.ColorCyan)
				fmt.Println(highlightText(res.String(), query))
			}
		},
	}
}
