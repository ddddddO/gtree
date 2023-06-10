package main

const template = `
- gtree
	- cmd
		- gtree
			- main.go
	- testdata
		- sample1.md
		- sample2.md
	- makefile
	- tree.go
`

const description = "- # Description\n" +
	"	- Output tree from markdown or programmatically.\n" +
	"		- Output format is tree|yaml|toml|json.\n" +
	"		- Default tree.\n" +
	"	- Make directories from markdown or programmatically.\n" +
	"		- It is possible to dry run.\n" +
	"		- You can use `-e` flag to make specified extensions as file.\n" +
	"	- Output a markdown template that can be used with either `output` subcommand or `mkdir` subcommand.\n" +
	"	- Provide CLI, Go library and Web.\n"
