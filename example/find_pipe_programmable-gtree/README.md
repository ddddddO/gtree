# *find_pipe_programmable-gtree*

This is simple program that outputs results of `find` command execution in tree.

## Using Templates

### 1. Install [`gonew`](https://pkg.go.dev/golang.org/x/tools/cmd/gonew) command

```console
$ go install golang.org/x/tools/cmd/gonew@latest
```

### 2. Clone by specifying template

```console
$ gonew github.com/ddddddO/gtree/example/find_pipe_programmable-gtree example.com/find_pipe_programmable-gtree
```

### 3. Run the program

```console
$ cd find_pipe_programmable-gtree/ && find . -type d -name .git -prune -o -type f -print | go run main.go
```

... You may have felt that the result of executing the find command in the project root was not enough. If so, try executing it in a place with many directories!

### 4. Arrange the program to your liking!

Thanks for reading to the end!
