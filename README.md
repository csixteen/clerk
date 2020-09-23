# clerk
Clerk is a simple personal assistant for the command-line. It allows you to manage notes and tasks in a simple way and search for stuff in your notes and tasks. If it sounds too simple, it's because it is. I prefer to have a single CLI tool that allows me to easily store notes and search them later instead of managing text files and have them scattered all over the place.

# Functionalities
- Tasks
 - Add a new task: `clerk-cli task add <name> <contents>...`
 - List existing tasks: `clerk-cli task list`
 - Edit a task (replaces the existing contents): `clerk-cli task edit <name | id> <new contents>`
 - Delete a task: `clerk-cli task del <name | id>`
 - Mark a task as completed: `clerk-cli task done <name | id>`

# References
- [go-sqlite3](https://github.com/mattn/go-sqlite3)
- [Package go-sqlite3](https://pkg.go.dev/github.com/mattn/go-sqlite3)
- [Unit Test (SQL) in Golang](https://medium.com/easyread/unit-test-sql-in-golang-5af19075e68e)
- [SQL As Understood by SQLite](https://sqlite.org/lang.html)
- [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests/)
- [SQLite Delete Cascade Not Working](https://stackoverflow.com/questions/13641250/sqlite-delete-cascade-not-working)
- [Cobra](https://github.com/spf13/cobra)
