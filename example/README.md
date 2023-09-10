# Templates for project using *`gtree`* library

You can set up either project with `gonew` command.<br>
The gtree library of each project generates programmable tree structure and outputs tree based on it to standard output.

## Templates
- *find_pipe_programmable-gtree*
  - This is simple program that outputs results of `find` command execution in tree.
  - Please refer to [README](find_pipe_programmable-gtree/README.md) for details.

- *go-list_pipe_programmable-gtree*
  - This is simple program that outputs results of '`go list -deps`' command execution in tree.
  - Please refer to [README](go-list_pipe_programmable-gtree/README.md) for details.

## Using Templates

### 1. Install [`gonew`](https://pkg.go.dev/golang.org/x/tools/cmd/gonew) command

```console
$ go install golang.org/x/tools/cmd/gonew@latest
```

### 2. Clone by specifying template

```console
$ gonew github.com/ddddddO/gtree/example/find_pipe_programmable-gtree example.com/find_pipe_programmable-gtree
```

or

```console
$ gonew github.com/ddddddO/gtree/example/go-list_pipe_programmable-gtree example.com/go-list_pipe_programmable-gtree
```

### 3. Run the program
#### For *find_pipe_programmable-gtree*
```console
$ find . -type d -name .git -prune -o -type f -print | go run main.go
```

... You may have felt that the result of executing the find command in the project root was not enough. If so, try executing it in a place with many directories!

#### For *go-list_pipe_programmable-gtree*
```console
$ go list -deps . | go run main.go
```

### 4. Arrange the program to your liking!

Thanks for reading to the end!
