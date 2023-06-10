package main

import "github.com/urfave/cli/v2"

const (
	exitCodeErrOpts = iota + 1
	exitCodeErrOutput
	exitCodeErrOpen
	exitCodeErrMkdir
)

func exitErrOpts(err error) cli.ExitCoder {
	return cli.Exit(err, exitCodeErrOpts)
}

func exitErrOutput(err error) cli.ExitCoder {
	return cli.Exit(err, exitCodeErrOutput)
}

func exitErrOpen(err error) cli.ExitCoder {
	return cli.Exit(err, exitCodeErrOpen)
}

func exitErrMkdir(err error) cli.ExitCoder {
	return cli.Exit(err, exitCodeErrMkdir)
}
