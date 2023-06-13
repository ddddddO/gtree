package main

import (
	"fmt"
)

type template string

func (t template) print() error {
	_, err := fmt.Print(string(t))
	return err
}

func (t template) println() error {
	if err := t.print(); err != nil {
		return err
	}
	_, err := fmt.Println()
	return err
}

const directory template = "" +
	"- gtree\n" +
	"	- cmd\n" +
	"		- gtree\n" +
	"			- main.go\n" +
	"	- testdata\n" +
	"		- sample1.md\n" +
	"		- sample2.md\n" +
	"	- Makefile\n" +
	"	- tree.go"

const description template = "" +
	"- # Description\n" +
	"	- Output tree from markdown or programmatically.\n" +
	"		- Output format is tree|yaml|toml|json.\n" +
	"		- Default tree.\n" +
	"	- Make directories from markdown or programmatically.\n" +
	"		- It is possible to dry run.\n" +
	"		- You can use `-e` flag to make specified extensions as file.\n" +
	"	- Output a markdown template that can be used with either `output` subcommand or `mkdir` subcommand.\n" +
	"	- Provide CLI, Go library and Web."
