# *go-list_pipe_programmable-gtree*

This is simple program that outputs results of '`go list -deps`' command execution in tree.

## Using Templates

### 1. Install [`gonew`](https://pkg.go.dev/golang.org/x/tools/cmd/gonew) command

```console
$ go install golang.org/x/tools/cmd/gonew@latest
```

### 2. Clone by specifying template

```console
$ gonew github.com/ddddddO/gtree/example/go-list_pipe_programmable-gtree example.com/go-list_pipe_programmable-gtree
```

### 3. Run the program

```console
$ cd go-list_pipe_programmable-gtree/ && go list -deps . | go run main.go
```

### 4. Arrange the program to your liking!

Thanks for reading to the end!
