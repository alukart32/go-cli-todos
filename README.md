# ToDo list CLI manager

This application is a cli tool for working with todos list.

The main operations are:
1. add
2. done
3. list
4. clear

## ToDo list repo

The default db is [bolt](https://github.com/boltdb/bolt).

The ToDo list db is stored at the path `/pkg/boltdbx/todos.db`.

The default bucket is `todos`.

## CLI library

The [cobra](https://github.com/spf13/cobra) was used for creating cli manager.

## Commands

All commands must be executed at the root of the project.

## Build
```bash
go build
```

## Lunch
### Windows
```bash
.\todos.exe
```

### Linux
```bash
./todos
```

## Add todo
```bash
todos add word "sentece"
```

## List todos
```bash
todos list
```

## Done todo
```bash
todos done todo_id
```

## Clear todos
```bash
todos clear
```
