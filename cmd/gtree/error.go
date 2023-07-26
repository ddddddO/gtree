package main

import "github.com/urfave/cli/v2"

const (
	exitCodeErrOpts = iota + 1
	exitCodeErrOpen
	exitCodeErrOutput
	exitCodeErrMkdir
	exitCodeErrVerify
)

func exitErrOpts(err error) cli.ExitCoder {
	return cli.Exit(err, exitCodeErrOpts)
}

func exitErrOpen(err error) cli.ExitCoder {
	return cli.Exit(err, exitCodeErrOpen)
}

func exitErrOutput(err error) cli.ExitCoder {
	return cli.Exit(err, exitCodeErrOutput)
}

func exitErrMkdir(err error) cli.ExitCoder {
	return cli.Exit(err, exitCodeErrMkdir)
}

func exitErrVerify(err error) cli.ExitCoder {
	return cli.Exit(err, exitCodeErrVerify)
}
