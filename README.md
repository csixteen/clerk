# clerk

Clerk started as a learning and practising exercise, but surprisingly enough I now use it on a regular basis. Clerk is supposed to be a simple personal assistant for the command-line. It allows you to manage notes and tasks in a simple way and search for stuff in your notes and tasks. If it sounds simple, that's because it **is** simple. My objective was to create a tool that helps me get shit done and retrieve information quickly, not something that I need to spend too much time tweaking or remembering how to use.

# How to use

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

- `clerk-cli search|s <query>...`

```
$ clerk-cli s clerk
note
- id: 1 | name: clerk
  Contents: add unit tests

```

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
- id: 1 | name: test | created_at: 2020-09-24 08:10:10
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

# Limitations and Caveats

I'm not using **Full Text Search** feature from SQLite, as it requires the module `fts5` to be avavailable. As such, I'm executing simple `SELECT` queries on `notes` and `tasks` tables. This is ok because I didn't intend to perform ultra complex search queries anyways. Also, it doesn't have any noticeable impact on performance.

# TODO
- Refactor the code. Things became a bit messy since I've introduced the `clerk-server`.
- Comment the code a bit better.
- Increase code coverage.
- Don't show *completed* tasks with `task list`, unless flag `--all` is passed.
- Do some house cleaning in general.
- Better error handling, maybe?

# Contributing

Yes, please. If you find a bug or want to improve the tool in general, feel free to submit a Pull-request.

# References
- [go-sqlite3](https://github.com/mattn/go-sqlite3)
- [Unit Test (SQL) in Golang](https://medium.com/easyread/unit-test-sql-in-golang-5af19075e68e)
- [SQL As Understood by SQLite](https://sqlite.org/lang.html)
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)
- [SQLite Delete Cascade Not Working](https://stackoverflow.com/questions/13641250/sqlite-delete-cascade-not-working)
- [Cobra](https://github.com/spf13/cobra)
