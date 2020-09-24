# clerk
Clerk started as a learning and practising exercise, but surprisingly enough I now use it on a regular basis. Clerk is supposed to be a simple personal assistant for the command-line. It allows you to manage notes and tasks in a simple way and search for stuff in your notes and tasks. If it sounds simple, that's because it **is** simple. My objective was to create a tool that helps me get shit done and retrieve information quickly, not something that I need to spend too much time tweaking or remembering how to use.

# Functionalities

### Tasks

- Add a new task: `clerk-cli task add <name> <contents>...`
- List existing tasks: `clerk-cli task list`
- Edit a task (replaces the existing contents): `clerk-cli task edit <name | id> <new contents>`
- Delete a task: `clerk-cli task del <name | id>`
- Mark a task as completed: `clerk-cli task done <name | id>`

### Notes

- Add a new note: `clerk-cli note add <name> <contents>...`
- List existing notes: `clerk-cli note list`
- Append contents to a note: `clerk-cli note append <name | id> <more contents>...`
- Show note contents: `clerk-cli note show <name | id>`
- Delete note: `clerk-cli note del <name | id>`

### Search
TODO

## Aliases

Most of the commands and subcommands have aliases, so that you don't need to type that much (you'll get shit done even faster...!!).

### Examples

```
# List notes
$ clerk-cli n ls

# Add new note called "test" with contents "some contents"
$ clerk-cli n a test some contents

# Show the contents of the newly added note
$ clerk-cli n sh test
- id: 1 | name: test | created_at: 0001-01-01 00:00:00
  Contents: some contents
```

Check the commands and subcommands' help to find their aliases: `clerk-cli <command> --help`.

# Dependencies

Clerk relies uses SQLite3 to store information. Most likely you'll have it already installed, but if you don't, then that's your only dependency.

# Installing

```
$ go get github.com/csixteen/clerk/cmd/clerk
```

# Building

The project uses Go modules, so you'll need a version of Go more recent than [1.11](https://blog.golang.org/using-go-modules).

```
$ make bin
go build -o clerk-cli cmd/clerk/*.go
```

# Testing

```
$ make test
go test -v pkg/actions/*.go
=== RUN   TestListTasks
--- PASS: TestListTasks (0.00s)
=== RUN   TestAddTask
--- PASS: TestAddTask (0.00s)
=== RUN   TestEditTask
--- PASS: TestEditTask (0.00s)
=== RUN   TestDeleteTask
--- PASS: TestDeleteTask (0.00s)
=== RUN   TestCompleteTask
--- PASS: TestCompleteTask (0.00s)
...
```

# References
- [go-sqlite3](https://github.com/mattn/go-sqlite3)
- [Package go-sqlite3](https://pkg.go.dev/github.com/mattn/go-sqlite3)
- [Unit Test (SQL) in Golang](https://medium.com/easyread/unit-test-sql-in-golang-5af19075e68e)
- [SQL As Understood by SQLite](https://sqlite.org/lang.html)
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)
- [SQLite Delete Cascade Not Working](https://stackoverflow.com/questions/13641250/sqlite-delete-cascade-not-working)
- [Cobra](https://github.com/spf13/cobra)
