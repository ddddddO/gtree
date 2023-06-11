package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ddddddO/gtree"
)

// Example:
// $ cd github.com/ddddddO/gtree
// $ find . -type d -name .git -prune -o -type f -print
// ./config.go
// ./node_generator_test.go
// ./example/like_cli/adapter/indentation.go
// ./example/like_cli/adapter/executor.go
// ./example/like_cli/main.go
// ./example/find_pipe_programmable-gtree/main.go
// ...
// $ find . -type d -name .git -prune -o -type f -print | go run example/find_pipe_programmable-gtree/main.go
// << See "Output:" below. >>
func main() {
	var (
		root *gtree.Node
		node *gtree.Node
	)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text() // e.g.) "./example/find_pipe_programmable-gtree/main.go"
		splited := strings.Split(line, "/") // e.g.) [. example find_pipe_programmable-gtree main.go]

		for i, s := range splited {
			if i == 0 {
				if root == nil {
					root = gtree.NewRoot(s)  // s := "."
					node = root
				}
				continue
			}
			tmp := node.Add(s)
			node = tmp
		}
		node = root
	}

	if err := gtree.OutputProgrammably(os.Stdout, root); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	// Output:
	// .
	// ├── config.go
	// ├── node_generator_test.go
	// ├── example
	// │   ├── like_cli
	// │   │   ├── adapter
	// │   │   │   ├── indentation.go
	// │   │   │   └── executor.go
	// │   │   └── main.go
	// │   ├── find_pipe_programmable-gtree
	// │   │   └── main.go
	// │   ├── go-list_pipe_programmable-gtree
	// │   │   └── main.go
	// │   └── programmable
	// │       └── main.go
	// ├── file_considerer.go
	// ├── node.go
	// ├── node_generator.go
	// ├── .gitignore
	// ├── wasm_tree.go
	// ├── input_spliter.go
	// ├── wasm_tree_handler.go
	// ├── tree_handler_mkdir_test.go
	// ├── tree_handler_benchmark_test.go
	// ├── LICENSE
	// ├── README.md
	// ├── tree_handler_programmably.go
	// ├── cli_mkdir_dryrun.png
	// ├── tree_mkdirer.go
	// ├── docs
	// │   ├── main.css
	// │   ├── robots.txt
	// │   ├── wasm_exec.js
	// │   ├── main.js
	// │   ├── main.wasm
	// │   ├── sitemap.xml
	// │   ├── toast.css
	// │   ├── service_worker.js
	// │   ├── index.html
	// │   ├── tab.js
	// │   └── toast.js
	// ├── tree_spreader.go
	// ├── stack.go
	// ├── CREDITS
	// ├── web_example.gif
	// ├── counter.go
	// ├── tmp.md
	// ├── .github
	// │   ├── dependabot.yml
	// │   └── workflows
	// │       ├── cd.yaml
	// │       └── ci.yaml
	// ├── .golangci.yaml
	// ├── markdown
	// │   ├── parser.go
	// │   ├── markdown.go
	// │   └── parser_test.go
	// ├── tree_grower.go
	// ├── performance.svg
	// ├── stack_test.go
	// ├── tree.go
	// ├── tree_handler.go
	// ├── tree_handler_programmably_mkdir_test.go
	// ├── root_generator.go
	// ├── testdata
	// │   ├── sample2.md
	// │   ├── sample4.md
	// │   ├── sample5.md
	// │   ├── sample1.md
	// │   ├── sample3.md
	// │   ├── sample8.md
	// │   ├── sample6.md
	// │   ├── demo.md
	// │   └── sample7.md
	// ├── cmd
	// │   ├── gtree
	// │   │   ├── indent.go
	// │   │   ├── mkdir.go
	// │   │   ├── main.go
	// │   │   ├── template.go
	// │   │   ├── error.go
	// │   │   └── output.go
	// │   └── gtree-wasm
	// │       ├── node_modules
	// │       │   ├── fastq
	// │       │   │   ├── example.mjs
	// │       │   │   ├── example.js
	// │       │   │   ├── package.json
	// │       │   │   ├── bench.js
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── index.d.ts
	// │       │   │   ├── .github
	// │       │   │   │   ├── dependabot.yml
	// │       │   │   │   └── workflows
	// │       │   │   │       └── ci.yml
	// │       │   │   ├── test
	// │       │   │   │   ├── test.js
	// │       │   │   │   ├── example.ts
	// │       │   │   │   ├── tsconfig.json
	// │       │   │   │   └── promise.js
	// │       │   │   └── queue.js
	// │       │   ├── queue-microtask
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   └── index.d.ts
	// │       │   ├── postcss-selector-parser
	// │       │   │   ├── package.json
	// │       │   │   ├── README.md
	// │       │   │   ├── API.md
	// │       │   │   ├── dist
	// │       │   │   │   ├── sortAscending.js
	// │       │   │   │   ├── parser.js
	// │       │   │   │   ├── tokenize.js
	// │       │   │   │   ├── index.js
	// │       │   │   │   ├── util
	// │       │   │   │   │   ├── unesc.js
	// │       │   │   │   │   ├── index.js
	// │       │   │   │   │   ├── stripComments.js
	// │       │   │   │   │   ├── getProp.js
	// │       │   │   │   │   └── ensureObject.js
	// │       │   │   │   ├── tokenTypes.js
	// │       │   │   │   ├── processor.js
	// │       │   │   │   └── selectors
	// │       │   │   │       ├── attribute.js
	// │       │   │   │       ├── universal.js
	// │       │   │   │       ├── guards.js
	// │       │   │   │       ├── root.js
	// │       │   │   │       ├── string.js
	// │       │   │   │       ├── types.js
	// │       │   │   │       ├── comment.js
	// │       │   │   │       ├── id.js
	// │       │   │   │       ├── index.js
	// │       │   │   │       ├── container.js
	// │       │   │   │       ├── constructors.js
	// │       │   │   │       ├── node.js
	// │       │   │   │       ├── tag.js
	// │       │   │   │       ├── combinator.js
	// │       │   │   │       ├── selector.js
	// │       │   │   │       ├── nesting.js
	// │       │   │   │       ├── pseudo.js
	// │       │   │   │       ├── namespace.js
	// │       │   │   │       └── className.js
	// │       │   │   ├── postcss-selector-parser.d.ts
	// │       │   │   ├── LICENSE-MIT
	// │       │   │   └── CHANGELOG.md
	// │       │   ├── dlv
	// │       │   │   ├── package.json
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   └── dist
	// │       │   │       ├── dlv.es.js
	// │       │   │       ├── dlv.js.map
	// │       │   │       ├── dlv.js
	// │       │   │       ├── dlv.umd.js.map
	// │       │   │       ├── dlv.umd.js
	// │       │   │       └── dlv.es.js.map
	// │       │   ├── path-is-absolute
	// │       │   │   ├── license
	// │       │   │   ├── package.json
	// │       │   │   ├── index.js
	// │       │   │   └── readme.md
	// │       │   ├── is-core-module
	// │       │   │   ├── package.json
	// │       │   │   ├── core.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   ├── .nycrc
	// │       │   │   ├── test
	// │       │   │   │   └── index.js
	// │       │   │   ├── .eslintrc
	// │       │   │   └── CHANGELOG.md
	// │       │   ├── thenify-all
	// │       │   │   ├── History.md
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── index.js
	// │       │   ├── arg
	// │       │   │   ├── package.json
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   ├── index.d.ts
	// │       │   │   └── LICENSE.md
	// │       │   ├── readdirp
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   └── index.d.ts
	// │       │   ├── any-promise
	// │       │   │   ├── loader.js
	// │       │   │   ├── register.d.ts
	// │       │   │   ├── register-shim.js
	// │       │   │   ├── package.json
	// │       │   │   ├── implementation.d.ts
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   ├── index.d.ts
	// │       │   │   ├── optional.js
	// │       │   │   ├── implementation.js
	// │       │   │   ├── .jshintrc
	// │       │   │   ├── .npmignore
	// │       │   │   ├── register.js
	// │       │   │   └── register
	// │       │   │       ├── rsvp.js
	// │       │   │       ├── native-promise-only.d.ts
	// │       │   │       ├── bluebird.js
	// │       │   │       ├── lie.js
	// │       │   │       ├── pinkie.d.ts
	// │       │   │       ├── lie.d.ts
	// │       │   │       ├── bluebird.d.ts
	// │       │   │       ├── promise.js
	// │       │   │       ├── when.d.ts
	// │       │   │       ├── vow.js
	// │       │   │       ├── pinkie.js
	// │       │   │       ├── promise.d.ts
	// │       │   │       ├── es6-promise.js
	// │       │   │       ├── native-promise-only.js
	// │       │   │       ├── when.js
	// │       │   │       ├── es6-promise.d.ts
	// │       │   │       ├── q.js
	// │       │   │       ├── vow.d.ts
	// │       │   │       ├── rsvp.d.ts
	// │       │   │       └── q.d.ts
	// │       │   ├── picomatch
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   ├── lib
	// │       │   │   │   ├── scan.js
	// │       │   │   │   ├── utils.js
	// │       │   │   │   ├── parse.js
	// │       │   │   │   ├── constants.js
	// │       │   │   │   └── picomatch.js
	// │       │   │   └── CHANGELOG.md
	// │       │   ├── binary-extensions
	// │       │   │   ├── binary-extensions.json
	// │       │   │   ├── license
	// │       │   │   ├── package.json
	// │       │   │   ├── index.js
	// │       │   │   ├── index.d.ts
	// │       │   │   ├── readme.md
	// │       │   │   └── binary-extensions.json.d.ts
	// │       │   ├── thenify
	// │       │   │   ├── History.md
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── index.js
	// │       │   ├── .package-lock.json
	// │       │   ├── fs.realpath
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   └── old.js
	// │       │   ├── anymatch
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   └── index.d.ts
	// │       │   ├── has
	// │       │   │   ├── src
	// │       │   │   │   └── index.js
	// │       │   │   ├── package.json
	// │       │   │   ├── README.md
	// │       │   │   ├── test
	// │       │   │   │   └── index.js
	// │       │   │   └── LICENSE-MIT
	// │       │   ├── once
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── once.js
	// │       │   ├── resolve
	// │       │   │   ├── example
	// │       │   │   │   ├── sync.js
	// │       │   │   │   └── async.js
	// │       │   │   ├── .editorconfig
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── index.js
	// │       │   │   ├── lib
	// │       │   │   │   ├── caller.js
	// │       │   │   │   ├── core.json
	// │       │   │   │   ├── node-modules-paths.js
	// │       │   │   │   ├── normalize-options.js
	// │       │   │   │   ├── core.js
	// │       │   │   │   ├── homedir.js
	// │       │   │   │   ├── sync.js
	// │       │   │   │   ├── is-core.js
	// │       │   │   │   └── async.js
	// │       │   │   ├── readme.markdown
	// │       │   │   ├── .github
	// │       │   │   │   └── FUNDING.yml
	// │       │   │   ├── test
	// │       │   │   │   ├── shadowed_core
	// │       │   │   │   │   └── node_modules
	// │       │   │   │   │       └── util
	// │       │   │   │   │           └── index.js
	// │       │   │   │   ├── home_paths_sync.js
	// │       │   │   │   ├── shadowed_core.js
	// │       │   │   │   ├── pathfilter.js
	// │       │   │   │   ├── node-modules-paths.js
	// │       │   │   │   ├── dotdot.js
	// │       │   │   │   ├── nonstring.js
	// │       │   │   │   ├── filter.js
	// │       │   │   │   ├── faulty_basedir.js
	// │       │   │   │   ├── core.js
	// │       │   │   │   ├── subdirs.js
	// │       │   │   │   ├── symlinks.js
	// │       │   │   │   ├── resolver
	// │       │   │   │   │   ├── incorrect_main
	// │       │   │   │   │   │   ├── package.json
	// │       │   │   │   │   │   └── index.js
	// │       │   │   │   │   ├── dot_main
	// │       │   │   │   │   │   ├── package.json
	// │       │   │   │   │   │   └── index.js
	// │       │   │   │   │   ├── symlinked
	// │       │   │   │   │   │   ├── package
	// │       │   │   │   │   │   │   ├── package.json
	// │       │   │   │   │   │   │   └── bar.js
	// │       │   │   │   │   │   └── _
	// │       │   │   │   │   │       ├── node_modules
	// │       │   │   │   │   │       │   └── foo.js
	// │       │   │   │   │   │       └── symlink_target
	// │       │   │   │   │   │           └── .gitkeep
	// │       │   │   │   │   ├── mug.coffee
	// │       │   │   │   │   ├── invalid_main
	// │       │   │   │   │   │   └── package.json
	// │       │   │   │   │   ├── foo.js
	// │       │   │   │   │   ├── quux
	// │       │   │   │   │   │   └── foo
	// │       │   │   │   │   │       └── index.js
	// │       │   │   │   │   ├── nested_symlinks
	// │       │   │   │   │   │   └── mylib
	// │       │   │   │   │   │       ├── package.json
	// │       │   │   │   │   │       ├── sync.js
	// │       │   │   │   │   │       └── async.js
	// │       │   │   │   │   ├── baz
	// │       │   │   │   │   │   ├── package.json
	// │       │   │   │   │   │   ├── doom.js
	// │       │   │   │   │   │   └── quux.js
	// │       │   │   │   │   ├── false_main
	// │       │   │   │   │   │   ├── package.json
	// │       │   │   │   │   │   └── index.js
	// │       │   │   │   │   ├── dot_slash_main
	// │       │   │   │   │   │   ├── package.json
	// │       │   │   │   │   │   └── index.js
	// │       │   │   │   │   ├── other_path
	// │       │   │   │   │   │   ├── root.js
	// │       │   │   │   │   │   └── lib
	// │       │   │   │   │   │       └── other-lib.js
	// │       │   │   │   │   ├── without_basedir
	// │       │   │   │   │   │   └── main.js
	// │       │   │   │   │   ├── malformed_package_json
	// │       │   │   │   │   │   ├── package.json
	// │       │   │   │   │   │   └── index.js
	// │       │   │   │   │   ├── mug.js
	// │       │   │   │   │   ├── same_names
	// │       │   │   │   │   │   ├── foo.js
	// │       │   │   │   │   │   └── foo
	// │       │   │   │   │   │       └── index.js
	// │       │   │   │   │   ├── cup.coffee
	// │       │   │   │   │   ├── multirepo
	// │       │   │   │   │   │   ├── packages
	// │       │   │   │   │   │   │   ├── package-b
	// │       │   │   │   │   │   │   │   ├── package.json
	// │       │   │   │   │   │   │   │   └── index.js
	// │       │   │   │   │   │   │   └── package-a
	// │       │   │   │   │   │   │       ├── package.json
	// │       │   │   │   │   │   │       └── index.js
	// │       │   │   │   │   │   ├── lerna.json
	// │       │   │   │   │   │   └── package.json
	// │       │   │   │   │   └── browser_field
	// │       │   │   │   │       ├── package.json
	// │       │   │   │   │       ├── b.js
	// │       │   │   │   │       └── a.js
	// │       │   │   │   ├── pathfilter
	// │       │   │   │   │   └── deep_ref
	// │       │   │   │   │       └── main.js
	// │       │   │   │   ├── resolver.js
	// │       │   │   │   ├── module_dir.js
	// │       │   │   │   ├── module_dir
	// │       │   │   │   │   ├── ymodules
	// │       │   │   │   │   │   └── aaa
	// │       │   │   │   │   │       └── index.js
	// │       │   │   │   │   ├── xmodules
	// │       │   │   │   │   │   └── aaa
	// │       │   │   │   │   │       └── index.js
	// │       │   │   │   │   └── zmodules
	// │       │   │   │   │       └── bbb
	// │       │   │   │   │           ├── package.json
	// │       │   │   │   │           └── main.js
	// │       │   │   │   ├── filter_sync.js
	// │       │   │   │   ├── precedence.js
	// │       │   │   │   ├── resolver_sync.js
	// │       │   │   │   ├── dotdot
	// │       │   │   │   │   ├── abc
	// │       │   │   │   │   │   └── index.js
	// │       │   │   │   │   └── index.js
	// │       │   │   │   ├── mock.js
	// │       │   │   │   ├── home_paths.js
	// │       │   │   │   ├── precedence
	// │       │   │   │   │   ├── bbb.js
	// │       │   │   │   │   ├── bbb
	// │       │   │   │   │   │   └── main.js
	// │       │   │   │   │   ├── aaa.js
	// │       │   │   │   │   └── aaa
	// │       │   │   │   │       ├── index.js
	// │       │   │   │   │       └── main.js
	// │       │   │   │   ├── node_path
	// │       │   │   │   │   ├── y
	// │       │   │   │   │   │   ├── ccc
	// │       │   │   │   │   │   │   └── index.js
	// │       │   │   │   │   │   └── bbb
	// │       │   │   │   │   │       └── index.js
	// │       │   │   │   │   └── x
	// │       │   │   │   │       ├── ccc
	// │       │   │   │   │       │   └── index.js
	// │       │   │   │   │       └── aaa
	// │       │   │   │   │           └── index.js
	// │       │   │   │   ├── mock_sync.js
	// │       │   │   │   └── node_path.js
	// │       │   │   ├── SECURITY.md
	// │       │   │   ├── .eslintrc
	// │       │   │   ├── bin
	// │       │   │   │   └── resolve
	// │       │   │   ├── sync.js
	// │       │   │   └── async.js
	// │       │   ├── run-parallel
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── index.js
	// │       │   ├── concat-map
	// │       │   │   ├── example
	// │       │   │   │   └── map.js
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.markdown
	// │       │   │   ├── index.js
	// │       │   │   ├── test
	// │       │   │   │   └── map.js
	// │       │   │   └── .travis.yml
	// │       │   ├── jiti
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── dist
	// │       │   │   │   ├── types.d.ts
	// │       │   │   │   ├── jiti.js
	// │       │   │   │   ├── plugins
	// │       │   │   │   │   ├── import-meta-env.d.ts
	// │       │   │   │   │   └── babel-plugin-transform-import-meta.d.ts
	// │       │   │   │   ├── utils.d.ts
	// │       │   │   │   ├── babel.d.ts
	// │       │   │   │   ├── babel.js
	// │       │   │   │   └── jiti.d.ts
	// │       │   │   ├── lib
	// │       │   │   │   └── index.js
	// │       │   │   ├── bin
	// │       │   │   │   └── jiti.js
	// │       │   │   └── register.js
	// │       │   ├── is-extglob
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── index.js
	// │       │   ├── normalize-path
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── index.js
	// │       │   ├── fast-glob
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── out
	// │       │   │       ├── types
	// │       │   │       │   ├── index.js
	// │       │   │       │   └── index.d.ts
	// │       │   │       ├── index.js
	// │       │   │       ├── index.d.ts
	// │       │   │       ├── managers
	// │       │   │       │   ├── patterns.d.ts
	// │       │   │       │   ├── tasks.d.ts
	// │       │   │       │   ├── patterns.js
	// │       │   │       │   └── tasks.js
	// │       │   │       ├── settings.d.ts
	// │       │   │       ├── providers
	// │       │   │       │   ├── provider.d.ts
	// │       │   │       │   ├── stream.d.ts
	// │       │   │       │   ├── provider.js
	// │       │   │       │   ├── filters
	// │       │   │       │   │   ├── error.js
	// │       │   │       │   │   ├── entry.js
	// │       │   │       │   │   ├── deep.js
	// │       │   │       │   │   ├── deep.d.ts
	// │       │   │       │   │   ├── entry.d.ts
	// │       │   │       │   │   └── error.d.ts
	// │       │   │       │   ├── matchers
	// │       │   │       │   │   ├── partial.d.ts
	// │       │   │       │   │   ├── matcher.js
	// │       │   │       │   │   ├── matcher.d.ts
	// │       │   │       │   │   └── partial.js
	// │       │   │       │   ├── async.d.ts
	// │       │   │       │   ├── sync.d.ts
	// │       │   │       │   ├── transformers
	// │       │   │       │   │   ├── entry.js
	// │       │   │       │   │   └── entry.d.ts
	// │       │   │       │   ├── sync.js
	// │       │   │       │   ├── stream.js
	// │       │   │       │   └── async.js
	// │       │   │       ├── utils
	// │       │   │       │   ├── errno.d.ts
	// │       │   │       │   ├── errno.js
	// │       │   │       │   ├── pattern.js
	// │       │   │       │   ├── string.d.ts
	// │       │   │       │   ├── stream.d.ts
	// │       │   │       │   ├── path.d.ts
	// │       │   │       │   ├── path.js
	// │       │   │       │   ├── string.js
	// │       │   │       │   ├── index.js
	// │       │   │       │   ├── index.d.ts
	// │       │   │       │   ├── fs.js
	// │       │   │       │   ├── array.js
	// │       │   │       │   ├── pattern.d.ts
	// │       │   │       │   ├── stream.js
	// │       │   │       │   ├── fs.d.ts
	// │       │   │       │   └── array.d.ts
	// │       │   │       ├── settings.js
	// │       │   │       └── readers
	// │       │   │           ├── stream.d.ts
	// │       │   │           ├── reader.js
	// │       │   │           ├── async.d.ts
	// │       │   │           ├── reader.d.ts
	// │       │   │           ├── sync.d.ts
	// │       │   │           ├── sync.js
	// │       │   │           ├── stream.js
	// │       │   │           └── async.js
	// │       │   ├── nanoid
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   ├── index.d.ts
	// │       │   │   ├── async
	// │       │   │   │   ├── index.native.js
	// │       │   │   │   ├── package.json
	// │       │   │   │   ├── index.js
	// │       │   │   │   ├── index.d.ts
	// │       │   │   │   ├── index.browser.js
	// │       │   │   │   ├── index.cjs
	// │       │   │   │   └── index.browser.cjs
	// │       │   │   ├── index.browser.js
	// │       │   │   ├── non-secure
	// │       │   │   │   ├── package.json
	// │       │   │   │   ├── index.js
	// │       │   │   │   ├── index.d.ts
	// │       │   │   │   └── index.cjs
	// │       │   │   ├── index.cjs
	// │       │   │   ├── bin
	// │       │   │   │   └── nanoid.cjs
	// │       │   │   ├── url-alphabet
	// │       │   │   │   ├── package.json
	// │       │   │   │   ├── index.js
	// │       │   │   │   └── index.cjs
	// │       │   │   ├── index.browser.cjs
	// │       │   │   └── nanoid.js
	// │       │   ├── read-cache
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── index.js
	// │       │   ├── sucrase
	// │       │   │   ├── node_modules
	// │       │   │   │   ├── commander
	// │       │   │   │   │   ├── package.json
	// │       │   │   │   │   ├── LICENSE
	// │       │   │   │   │   ├── index.js
	// │       │   │   │   │   ├── typings
	// │       │   │   │   │   │   └── index.d.ts
	// │       │   │   │   │   ├── Readme.md
	// │       │   │   │   │   └── CHANGELOG.md
	// │       │   │   │   └── glob
	// │       │   │   │       ├── package.json
	// │       │   │   │       ├── LICENSE
	// │       │   │   │       ├── README.md
	// │       │   │   │       ├── changelog.md
	// │       │   │   │       ├── glob.js
	// │       │   │   │       ├── sync.js
	// │       │   │   │       └── common.js
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── dist
	// │       │   │   │   ├── TokenProcessor.js
	// │       │   │   │   ├── types
	// │       │   │   │   │   ├── HelperManager.d.ts
	// │       │   │   │   │   ├── identifyShadowedGlobals.d.ts
	// │       │   │   │   │   ├── register.d.ts
	// │       │   │   │   │   ├── computeSourceMap.d.ts
	// │       │   │   │   │   ├── cli.d.ts
	// │       │   │   │   │   ├── index.d.ts
	// │       │   │   │   │   ├── Options.d.ts
	// │       │   │   │   │   ├── TokenProcessor.d.ts
	// │       │   │   │   │   ├── CJSImportProcessor.d.ts
	// │       │   │   │   │   ├── util
	// │       │   │   │   │   │   ├── getImportExportSpecifierInfo.d.ts
	// │       │   │   │   │   │   ├── getDeclarationInfo.d.ts
	// │       │   │   │   │   │   ├── getJSXPragmaInfo.d.ts
	// │       │   │   │   │   │   ├── isIdentifier.d.ts
	// │       │   │   │   │   │   ├── getTSImportedNames.d.ts
	// │       │   │   │   │   │   ├── formatTokens.d.ts
	// │       │   │   │   │   │   ├── shouldElideDefaultExport.d.ts
	// │       │   │   │   │   │   ├── elideImportEquals.d.ts
	// │       │   │   │   │   │   ├── removeMaybeImportAssertion.d.ts
	// │       │   │   │   │   │   ├── isAsyncOperation.d.ts
	// │       │   │   │   │   │   ├── getNonTypeIdentifiers.d.ts
	// │       │   │   │   │   │   ├── getIdentifierNames.d.ts
	// │       │   │   │   │   │   └── getClassInfo.d.ts
	// │       │   │   │   │   ├── parser
	// │       │   │   │   │   │   ├── index.d.ts
	// │       │   │   │   │   │   ├── plugins
	// │       │   │   │   │   │   │   ├── types.d.ts
	// │       │   │   │   │   │   │   ├── jsx
	// │       │   │   │   │   │   │   │   ├── index.d.ts
	// │       │   │   │   │   │   │   │   └── xhtml.d.ts
	// │       │   │   │   │   │   │   ├── typescript.d.ts
	// │       │   │   │   │   │   │   └── flow.d.ts
	// │       │   │   │   │   │   ├── util
	// │       │   │   │   │   │   │   ├── charcodes.d.ts
	// │       │   │   │   │   │   │   ├── identifier.d.ts
	// │       │   │   │   │   │   │   └── whitespace.d.ts
	// │       │   │   │   │   │   ├── tokenizer
	// │       │   │   │   │   │   │   ├── state.d.ts
	// │       │   │   │   │   │   │   ├── types.d.ts
	// │       │   │   │   │   │   │   ├── readWordTree.d.ts
	// │       │   │   │   │   │   │   ├── index.d.ts
	// │       │   │   │   │   │   │   ├── keywords.d.ts
	// │       │   │   │   │   │   │   └── readWord.d.ts
	// │       │   │   │   │   │   └── traverser
	// │       │   │   │   │   │       ├── expression.d.ts
	// │       │   │   │   │   │       ├── util.d.ts
	// │       │   │   │   │   │       ├── lval.d.ts
	// │       │   │   │   │   │       ├── base.d.ts
	// │       │   │   │   │   │       ├── index.d.ts
	// │       │   │   │   │   │       └── statement.d.ts
	// │       │   │   │   │   ├── Options-gen-types.d.ts
	// │       │   │   │   │   ├── transformers
	// │       │   │   │   │   │   ├── OptionalCatchBindingTransformer.d.ts
	// │       │   │   │   │   │   ├── JSXTransformer.d.ts
	// │       │   │   │   │   │   ├── Transformer.d.ts
	// │       │   │   │   │   │   ├── JestHoistTransformer.d.ts
	// │       │   │   │   │   │   ├── ESMImportTransformer.d.ts
	// │       │   │   │   │   │   ├── NumericSeparatorTransformer.d.ts
	// │       │   │   │   │   │   ├── TypeScriptTransformer.d.ts
	// │       │   │   │   │   │   ├── RootTransformer.d.ts
	// │       │   │   │   │   │   ├── ReactHotLoaderTransformer.d.ts
	// │       │   │   │   │   │   ├── ReactDisplayNameTransformer.d.ts
	// │       │   │   │   │   │   ├── CJSImportTransformer.d.ts
	// │       │   │   │   │   │   ├── OptionalChainingNullishTransformer.d.ts
	// │       │   │   │   │   │   └── FlowTransformer.d.ts
	// │       │   │   │   │   └── NameManager.d.ts
	// │       │   │   │   ├── HelperManager.js
	// │       │   │   │   ├── esm
	// │       │   │   │   │   ├── TokenProcessor.js
	// │       │   │   │   │   ├── HelperManager.js
	// │       │   │   │   │   ├── CJSImportProcessor.js
	// │       │   │   │   │   ├── NameManager.js
	// │       │   │   │   │   ├── index.js
	// │       │   │   │   │   ├── Options-gen-types.js
	// │       │   │   │   │   ├── Options.js
	// │       │   │   │   │   ├── util
	// │       │   │   │   │   │   ├── isIdentifier.js
	// │       │   │   │   │   │   ├── getJSXPragmaInfo.js
	// │       │   │   │   │   │   ├── getDeclarationInfo.js
	// │       │   │   │   │   │   ├── formatTokens.js
	// │       │   │   │   │   │   ├── elideImportEquals.js
	// │       │   │   │   │   │   ├── getClassInfo.js
	// │       │   │   │   │   │   ├── getImportExportSpecifierInfo.js
	// │       │   │   │   │   │   ├── getTSImportedNames.js
	// │       │   │   │   │   │   ├── getIdentifierNames.js
	// │       │   │   │   │   │   ├── getNonTypeIdentifiers.js
	// │       │   │   │   │   │   ├── removeMaybeImportAssertion.js
	// │       │   │   │   │   │   ├── isAsyncOperation.js
	// │       │   │   │   │   │   └── shouldElideDefaultExport.js
	// │       │   │   │   │   ├── parser
	// │       │   │   │   │   │   ├── index.js
	// │       │   │   │   │   │   ├── plugins
	// │       │   │   │   │   │   │   ├── jsx
	// │       │   │   │   │   │   │   │   ├── index.js
	// │       │   │   │   │   │   │   │   └── xhtml.js
	// │       │   │   │   │   │   │   ├── types.js
	// │       │   │   │   │   │   │   ├── typescript.js
	// │       │   │   │   │   │   │   └── flow.js
	// │       │   │   │   │   │   ├── util
	// │       │   │   │   │   │   │   ├── whitespace.js
	// │       │   │   │   │   │   │   ├── identifier.js
	// │       │   │   │   │   │   │   └── charcodes.js
	// │       │   │   │   │   │   ├── tokenizer
	// │       │   │   │   │   │   │   ├── types.js
	// │       │   │   │   │   │   │   ├── index.js
	// │       │   │   │   │   │   │   ├── readWord.js
	// │       │   │   │   │   │   │   ├── readWordTree.js
	// │       │   │   │   │   │   │   ├── keywords.js
	// │       │   │   │   │   │   │   └── state.js
	// │       │   │   │   │   │   └── traverser
	// │       │   │   │   │   │       ├── index.js
	// │       │   │   │   │   │       ├── base.js
	// │       │   │   │   │   │       ├── lval.js
	// │       │   │   │   │   │       ├── expression.js
	// │       │   │   │   │   │       ├── statement.js
	// │       │   │   │   │   │       └── util.js
	// │       │   │   │   │   ├── cli.js
	// │       │   │   │   │   ├── computeSourceMap.js
	// │       │   │   │   │   ├── transformers
	// │       │   │   │   │   │   ├── FlowTransformer.js
	// │       │   │   │   │   │   ├── RootTransformer.js
	// │       │   │   │   │   │   ├── Transformer.js
	// │       │   │   │   │   │   ├── OptionalCatchBindingTransformer.js
	// │       │   │   │   │   │   ├── NumericSeparatorTransformer.js
	// │       │   │   │   │   │   ├── ReactHotLoaderTransformer.js
	// │       │   │   │   │   │   ├── JestHoistTransformer.js
	// │       │   │   │   │   │   ├── JSXTransformer.js
	// │       │   │   │   │   │   ├── OptionalChainingNullishTransformer.js
	// │       │   │   │   │   │   ├── ReactDisplayNameTransformer.js
	// │       │   │   │   │   │   ├── TypeScriptTransformer.js
	// │       │   │   │   │   │   ├── CJSImportTransformer.js
	// │       │   │   │   │   │   └── ESMImportTransformer.js
	// │       │   │   │   │   ├── register.js
	// │       │   │   │   │   └── identifyShadowedGlobals.js
	// │       │   │   │   ├── CJSImportProcessor.js
	// │       │   │   │   ├── NameManager.js
	// │       │   │   │   ├── index.js
	// │       │   │   │   ├── Options-gen-types.js
	// │       │   │   │   ├── Options.js
	// │       │   │   │   ├── util
	// │       │   │   │   │   ├── isIdentifier.js
	// │       │   │   │   │   ├── getJSXPragmaInfo.js
	// │       │   │   │   │   ├── getDeclarationInfo.js
	// │       │   │   │   │   ├── formatTokens.js
	// │       │   │   │   │   ├── elideImportEquals.js
	// │       │   │   │   │   ├── getClassInfo.js
	// │       │   │   │   │   ├── getImportExportSpecifierInfo.js
	// │       │   │   │   │   ├── getTSImportedNames.js
	// │       │   │   │   │   ├── getIdentifierNames.js
	// │       │   │   │   │   ├── getNonTypeIdentifiers.js
	// │       │   │   │   │   ├── removeMaybeImportAssertion.js
	// │       │   │   │   │   ├── isAsyncOperation.js
	// │       │   │   │   │   └── shouldElideDefaultExport.js
	// │       │   │   │   ├── parser
	// │       │   │   │   │   ├── index.js
	// │       │   │   │   │   ├── plugins
	// │       │   │   │   │   │   ├── jsx
	// │       │   │   │   │   │   │   ├── index.js
	// │       │   │   │   │   │   │   └── xhtml.js
	// │       │   │   │   │   │   ├── types.js
	// │       │   │   │   │   │   ├── typescript.js
	// │       │   │   │   │   │   └── flow.js
	// │       │   │   │   │   ├── util
	// │       │   │   │   │   │   ├── whitespace.js
	// │       │   │   │   │   │   ├── identifier.js
	// │       │   │   │   │   │   └── charcodes.js
	// │       │   │   │   │   ├── tokenizer
	// │       │   │   │   │   │   ├── types.js
	// │       │   │   │   │   │   ├── index.js
	// │       │   │   │   │   │   ├── readWord.js
	// │       │   │   │   │   │   ├── readWordTree.js
	// │       │   │   │   │   │   ├── keywords.js
	// │       │   │   │   │   │   └── state.js
	// │       │   │   │   │   └── traverser
	// │       │   │   │   │       ├── index.js
	// │       │   │   │   │       ├── base.js
	// │       │   │   │   │       ├── lval.js
	// │       │   │   │   │       ├── expression.js
	// │       │   │   │   │       ├── statement.js
	// │       │   │   │   │       └── util.js
	// │       │   │   │   ├── cli.js
	// │       │   │   │   ├── computeSourceMap.js
	// │       │   │   │   ├── transformers
	// │       │   │   │   │   ├── FlowTransformer.js
	// │       │   │   │   │   ├── RootTransformer.js
	// │       │   │   │   │   ├── Transformer.js
	// │       │   │   │   │   ├── OptionalCatchBindingTransformer.js
	// │       │   │   │   │   ├── NumericSeparatorTransformer.js
	// │       │   │   │   │   ├── ReactHotLoaderTransformer.js
	// │       │   │   │   │   ├── JestHoistTransformer.js
	// │       │   │   │   │   ├── JSXTransformer.js
	// │       │   │   │   │   ├── OptionalChainingNullishTransformer.js
	// │       │   │   │   │   ├── ReactDisplayNameTransformer.js
	// │       │   │   │   │   ├── TypeScriptTransformer.js
	// │       │   │   │   │   ├── CJSImportTransformer.js
	// │       │   │   │   │   └── ESMImportTransformer.js
	// │       │   │   │   ├── register.js
	// │       │   │   │   └── identifyShadowedGlobals.js
	// │       │   │   ├── bin
	// │       │   │   │   ├── sucrase-node
	// │       │   │   │   └── sucrase
	// │       │   │   ├── ts-node-plugin
	// │       │   │   │   └── index.js
	// │       │   │   └── register
	// │       │   │       ├── tsx-legacy-module-interop.js
	// │       │   │       ├── index.js
	// │       │   │       ├── tsx.js
	// │       │   │       ├── ts.js
	// │       │   │       ├── jsx.js
	// │       │   │       ├── ts-legacy-module-interop.js
	// │       │   │       └── js.js
	// │       │   ├── @tailwindcss
	// │       │   │   ├── forms
	// │       │   │   │   ├── src
	// │       │   │   │   │   ├── index.js
	// │       │   │   │   │   └── index.d.ts
	// │       │   │   │   ├── package.json
	// │       │   │   │   ├── LICENSE
	// │       │   │   │   ├── README.md
	// │       │   │   │   ├── tailwind.config.js
	// │       │   │   │   ├── .github
	// │       │   │   │   │   ├── workflows
	// │       │   │   │   │   │   └── release-insiders.yml
	// │       │   │   │   │   └── ISSUE_TEMPLATE
	// │       │   │   │   │       ├── 1.bug_report.yml
	// │       │   │   │   │       └── config.yml
	// │       │   │   │   ├── index.html
	// │       │   │   │   ├── kitchen-sink.html
	// │       │   │   │   └── CHANGELOG.md
	// │       │   │   └── aspect-ratio
	// │       │   │       ├── src
	// │       │   │       │   ├── index.js
	// │       │   │       │   └── index.d.ts
	// │       │   │       ├── package.json
	// │       │   │       ├── tests
	// │       │   │       │   └── test.js
	// │       │   │       ├── README.md
	// │       │   │       ├── .github
	// │       │   │       │   ├── workflows
	// │       │   │       │   │   └── release-insiders.yml
	// │       │   │       │   └── ISSUE_TEMPLATE
	// │       │   │       │       ├── 1.bug_report.yml
	// │       │   │       │       └── config.yml
	// │       │   │       └── CHANGELOG.md
	// │       │   ├── postcss-import
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   └── lib
	// │       │   │       ├── parse-statements.js
	// │       │   │       ├── join-media.js
	// │       │   │       ├── process-content.js
	// │       │   │       ├── load-content.js
	// │       │   │       ├── assign-layer-names.js
	// │       │   │       ├── data-url.js
	// │       │   │       ├── resolve-id.js
	// │       │   │       └── join-layer.js
	// │       │   ├── lilconfig
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── dist
	// │       │   │   │   ├── index.js
	// │       │   │   │   └── index.d.ts
	// │       │   │   └── readme.md
	// │       │   ├── commander
	// │       │   │   ├── package.json
	// │       │   │   ├── esm.mjs
	// │       │   │   ├── LICENSE
	// │       │   │   ├── package-support.json
	// │       │   │   ├── index.js
	// │       │   │   ├── typings
	// │       │   │   │   └── index.d.ts
	// │       │   │   ├── Readme.md
	// │       │   │   └── CHANGELOG.md
	// │       │   ├── @jridgewell
	// │       │   │   ├── trace-mapping
	// │       │   │   │   ├── node_modules
	// │       │   │   │   │   └── @jridgewell
	// │       │   │   │   │       └── sourcemap-codec
	// │       │   │   │   │           ├── src
	// │       │   │   │   │           │   └── sourcemap-codec.ts
	// │       │   │   │   │           ├── package.json
	// │       │   │   │   │           ├── LICENSE
	// │       │   │   │   │           ├── README.md
	// │       │   │   │   │           └── dist
	// │       │   │   │   │               ├── sourcemap-codec.mjs.map
	// │       │   │   │   │               ├── types
	// │       │   │   │   │               │   └── sourcemap-codec.d.ts
	// │       │   │   │   │               ├── sourcemap-codec.umd.js.map
	// │       │   │   │   │               ├── sourcemap-codec.umd.js
	// │       │   │   │   │               └── sourcemap-codec.mjs
	// │       │   │   │   ├── package.json
	// │       │   │   │   ├── LICENSE
	// │       │   │   │   ├── README.md
	// │       │   │   │   └── dist
	// │       │   │   │       ├── trace-mapping.mjs.map
	// │       │   │   │       ├── types
	// │       │   │   │       │   ├── types.d.ts
	// │       │   │   │       │   ├── any-map.d.ts
	// │       │   │   │       │   ├── trace-mapping.d.ts
	// │       │   │   │       │   ├── strip-filename.d.ts
	// │       │   │   │       │   ├── binary-search.d.ts
	// │       │   │   │       │   ├── sourcemap-segment.d.ts
	// │       │   │   │       │   ├── resolve.d.ts
	// │       │   │   │       │   ├── sort.d.ts
	// │       │   │   │       │   └── by-source.d.ts
	// │       │   │   │       ├── trace-mapping.umd.js.map
	// │       │   │   │       ├── trace-mapping.umd.js
	// │       │   │   │       └── trace-mapping.mjs
	// │       │   │   ├── resolve-uri
	// │       │   │   │   ├── package.json
	// │       │   │   │   ├── LICENSE
	// │       │   │   │   ├── README.md
	// │       │   │   │   └── dist
	// │       │   │   │       ├── types
	// │       │   │   │       │   └── resolve-uri.d.ts
	// │       │   │   │       ├── resolve-uri.umd.js.map
	// │       │   │   │       ├── resolve-uri.mjs
	// │       │   │   │       ├── resolve-uri.umd.js
	// │       │   │   │       └── resolve-uri.mjs.map
	// │       │   │   ├── sourcemap-codec
	// │       │   │   │   ├── package.json
	// │       │   │   │   ├── LICENSE
	// │       │   │   │   ├── README.md
	// │       │   │   │   └── dist
	// │       │   │   │       ├── sourcemap-codec.mjs.map
	// │       │   │   │       ├── types
	// │       │   │   │       │   └── sourcemap-codec.d.ts
	// │       │   │   │       ├── sourcemap-codec.umd.js.map
	// │       │   │   │       ├── sourcemap-codec.umd.js
	// │       │   │   │       └── sourcemap-codec.mjs
	// │       │   │   ├── set-array
	// │       │   │   │   ├── src
	// │       │   │   │   │   └── set-array.ts
	// │       │   │   │   ├── package.json
	// │       │   │   │   ├── LICENSE
	// │       │   │   │   ├── README.md
	// │       │   │   │   └── dist
	// │       │   │   │       ├── set-array.mjs
	// │       │   │   │       ├── set-array.umd.js
	// │       │   │   │       ├── types
	// │       │   │   │       │   └── set-array.d.ts
	// │       │   │   │       ├── set-array.umd.js.map
	// │       │   │   │       └── set-array.mjs.map
	// │       │   │   └── gen-mapping
	// │       │   │       ├── package.json
	// │       │   │       ├── LICENSE
	// │       │   │       ├── README.md
	// │       │   │       └── dist
	// │       │   │           ├── gen-mapping.mjs
	// │       │   │           ├── types
	// │       │   │           │   ├── types.d.ts
	// │       │   │           │   ├── gen-mapping.d.ts
	// │       │   │           │   └── sourcemap-segment.d.ts
	// │       │   │           ├── gen-mapping.umd.js
	// │       │   │           ├── gen-mapping.umd.js.map
	// │       │   │           └── gen-mapping.mjs.map
	// │       │   ├── mini-svg-data-uri
	// │       │   │   ├── package.json
	// │       │   │   ├── shorter-css-color-names.js
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── index.test-d.ts
	// │       │   │   ├── index.js
	// │       │   │   ├── index.d.ts
	// │       │   │   └── cli.js
	// │       │   ├── function-bind
	// │       │   │   ├── .editorconfig
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   ├── .jscs.json
	// │       │   │   ├── implementation.js
	// │       │   │   ├── test
	// │       │   │   │   ├── index.js
	// │       │   │   │   └── .eslintrc
	// │       │   │   ├── .npmignore
	// │       │   │   ├── .eslintrc
	// │       │   │   └── .travis.yml
	// │       │   ├── reusify
	// │       │   │   ├── package.json
	// │       │   │   ├── reusify.js
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── test.js
	// │       │   │   ├── benchmarks
	// │       │   │   │   ├── reuseNoCodeFunction.js
	// │       │   │   │   ├── fib.js
	// │       │   │   │   └── createNoCodeFunction.js
	// │       │   │   ├── .travis.yml
	// │       │   │   └── .coveralls.yml
	// │       │   ├── glob
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── glob.js
	// │       │   │   ├── sync.js
	// │       │   │   └── common.js
	// │       │   ├── postcss-value-parser
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── lib
	// │       │   │       ├── unit.js
	// │       │   │       ├── index.js
	// │       │   │       ├── index.d.ts
	// │       │   │       ├── walk.js
	// │       │   │       ├── parse.js
	// │       │   │       └── stringify.js
	// │       │   ├── source-map-js
	// │       │   │   ├── source-map.d.ts
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── lib
	// │       │   │   │   ├── binary-search.js
	// │       │   │   │   ├── source-map-consumer.js
	// │       │   │   │   ├── source-node.js
	// │       │   │   │   ├── base64-vlq.js
	// │       │   │   │   ├── base64.js
	// │       │   │   │   ├── source-map-generator.js
	// │       │   │   │   ├── quick-sort.js
	// │       │   │   │   ├── mapping-list.js
	// │       │   │   │   ├── util.js
	// │       │   │   │   └── array-set.js
	// │       │   │   ├── source-map.js
	// │       │   │   └── CHANGELOG.md
	// │       │   ├── mz
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── readline.js
	// │       │   │   ├── child_process.js
	// │       │   │   ├── index.js
	// │       │   │   ├── fs.js
	// │       │   │   ├── zlib.js
	// │       │   │   ├── crypto.js
	// │       │   │   ├── HISTORY.md
	// │       │   │   └── dns.js
	// │       │   ├── pify
	// │       │   │   ├── license
	// │       │   │   ├── package.json
	// │       │   │   ├── index.js
	// │       │   │   └── readme.md
	// │       │   ├── object-assign
	// │       │   │   ├── license
	// │       │   │   ├── package.json
	// │       │   │   ├── index.js
	// │       │   │   └── readme.md
	// │       │   ├── uglify-js
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── lib
	// │       │   │   │   ├── transform.js
	// │       │   │   │   ├── ast.js
	// │       │   │   │   ├── sourcemap.js
	// │       │   │   │   ├── utils.js
	// │       │   │   │   ├── compress.js
	// │       │   │   │   ├── parse.js
	// │       │   │   │   ├── output.js
	// │       │   │   │   ├── minify.js
	// │       │   │   │   ├── mozilla-ast.js
	// │       │   │   │   ├── propmangle.js
	// │       │   │   │   └── scope.js
	// │       │   │   ├── bin
	// │       │   │   │   └── uglifyjs
	// │       │   │   └── tools
	// │       │   │       ├── exports.js
	// │       │   │       ├── tty.js
	// │       │   │       ├── node.js
	// │       │   │       ├── domprops.json
	// │       │   │       └── domprops.html
	// │       │   ├── inflight
	// │       │   │   ├── inflight.js
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   └── README.md
	// │       │   ├── is-glob
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── index.js
	// │       │   ├── clean-css
	// │       │   │   ├── History.md
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   └── lib
	// │       │   │       ├── clean.js
	// │       │   │       ├── options
	// │       │   │       │   ├── rounding-precision.js
	// │       │   │       │   ├── plugins.js
	// │       │   │       │   ├── fetch.js
	// │       │   │       │   ├── format.js
	// │       │   │       │   ├── inline-request.js
	// │       │   │       │   ├── inline-timeout.js
	// │       │   │       │   ├── compatibility.js
	// │       │   │       │   ├── optimization-level.js
	// │       │   │       │   ├── rebase-to.js
	// │       │   │       │   ├── inline.js
	// │       │   │       │   └── rebase.js
	// │       │   │       ├── writer
	// │       │   │       │   ├── source-maps.js
	// │       │   │       │   ├── simple.js
	// │       │   │       │   ├── one-time.js
	// │       │   │       │   └── helpers.js
	// │       │   │       ├── reader
	// │       │   │       │   ├── restore-import.js
	// │       │   │       │   ├── is-allowed-resource.js
	// │       │   │       │   ├── read-sources.js
	// │       │   │       │   ├── normalize-path.js
	// │       │   │       │   ├── load-remote-resource.js
	// │       │   │       │   ├── rebase-local-map.js
	// │       │   │       │   ├── apply-source-maps.js
	// │       │   │       │   ├── input-source-map-tracker.js
	// │       │   │       │   ├── extract-import-url-and-media.js
	// │       │   │       │   ├── match-data-uri.js
	// │       │   │       │   ├── rewrite-url.js
	// │       │   │       │   ├── rebase-remote-map.js
	// │       │   │       │   ├── load-original-sources.js
	// │       │   │       │   └── rebase.js
	// │       │   │       ├── optimizer
	// │       │   │       │   ├── validator.js
	// │       │   │       │   ├── vendor-prefixes.js
	// │       │   │       │   ├── level-2
	// │       │   │       │   │   ├── restore-with-components.js
	// │       │   │       │   │   ├── remove-duplicates.js
	// │       │   │       │   │   ├── merge-non-adjacent-by-selector.js
	// │       │   │       │   │   ├── specificities-overlap.js
	// │       │   │       │   │   ├── merge-media-queries.js
	// │       │   │       │   │   ├── remove-duplicate-font-at-rules.js
	// │       │   │       │   │   ├── optimize.js
	// │       │   │       │   │   ├── merge-non-adjacent-by-body.js
	// │       │   │       │   │   ├── tidy-rule-duplicates.js
	// │       │   │       │   │   ├── restructure.js
	// │       │   │       │   │   ├── extract-properties.js
	// │       │   │       │   │   ├── remove-duplicate-media-queries.js
	// │       │   │       │   │   ├── remove-unused-at-rules.js
	// │       │   │       │   │   ├── rules-overlap.js
	// │       │   │       │   │   ├── merge-adjacent.js
	// │       │   │       │   │   ├── specificity.js
	// │       │   │       │   │   ├── properties
	// │       │   │       │   │   │   ├── override-properties.js
	// │       │   │       │   │   │   ├── optimize.js
	// │       │   │       │   │   │   ├── has-inherit.js
	// │       │   │       │   │   │   ├── find-component-in.js
	// │       │   │       │   │   │   ├── merge-into-shorthands.js
	// │       │   │       │   │   │   ├── has-unset.js
	// │       │   │       │   │   │   ├── has-same-values.js
	// │       │   │       │   │   │   ├── is-mergeable-shorthand.js
	// │       │   │       │   │   │   ├── every-values-pair.js
	// │       │   │       │   │   │   ├── overrides-non-component-shorthand.js
	// │       │   │       │   │   │   ├── populate-components.js
	// │       │   │       │   │   │   └── is-component-of.js
	// │       │   │       │   │   ├── reduce-non-adjacent.js
	// │       │   │       │   │   ├── is-mergeable.js
	// │       │   │       │   │   └── reorderable.js
	// │       │   │       │   ├── wrap-for-optimizing.js
	// │       │   │       │   ├── hack.js
	// │       │   │       │   ├── configuration.js
	// │       │   │       │   ├── invalid-property-error.js
	// │       │   │       │   ├── remove-unused.js
	// │       │   │       │   ├── level-1
	// │       │   │       │   │   ├── tidy-block.js
	// │       │   │       │   │   ├── tidy-at-rule.js
	// │       │   │       │   │   ├── optimize.js
	// │       │   │       │   │   ├── property-optimizers.js
	// │       │   │       │   │   ├── tidy-rules.js
	// │       │   │       │   │   ├── sort-selectors.js
	// │       │   │       │   │   ├── value-optimizers
	// │       │   │       │   │   │   ├── color.js
	// │       │   │       │   │   │   ├── whitespace.js
	// │       │   │       │   │   │   ├── url-quotes.js
	// │       │   │       │   │   │   ├── precision.js
	// │       │   │       │   │   │   ├── unit.js
	// │       │   │       │   │   │   ├── fraction.js
	// │       │   │       │   │   │   ├── degrees.js
	// │       │   │       │   │   │   ├── url-whitespace.js
	// │       │   │       │   │   │   ├── starts-as-url.js
	// │       │   │       │   │   │   ├── url-prefix.js
	// │       │   │       │   │   │   ├── time.js
	// │       │   │       │   │   │   ├── text-quotes.js
	// │       │   │       │   │   │   ├── color
	// │       │   │       │   │   │   │   ├── shorten-hsl.js
	// │       │   │       │   │   │   │   ├── shorten-rgb.js
	// │       │   │       │   │   │   │   └── shorten-hex.js
	// │       │   │       │   │   │   └── zero.js
	// │       │   │       │   │   ├── property-optimizers
	// │       │   │       │   │   │   ├── outline.js
	// │       │   │       │   │   │   ├── padding.js
	// │       │   │       │   │   │   ├── filter.js
	// │       │   │       │   │   │   ├── border-radius.js
	// │       │   │       │   │   │   ├── box-shadow.js
	// │       │   │       │   │   │   ├── background.js
	// │       │   │       │   │   │   ├── margin.js
	// │       │   │       │   │   │   └── font-weight.js
	// │       │   │       │   │   └── value-optimizers.js
	// │       │   │       │   ├── level-0
	// │       │   │       │   │   └── optimize.js
	// │       │   │       │   ├── restore-from-optimizing.js
	// │       │   │       │   ├── configuration
	// │       │   │       │   │   ├── restore.js
	// │       │   │       │   │   ├── break-up.js
	// │       │   │       │   │   ├── properties
	// │       │   │       │   │   │   └── understandable.js
	// │       │   │       │   │   └── can-override.js
	// │       │   │       │   └── clone.js
	// │       │   │       ├── utils
	// │       │   │       │   ├── is-http-resource.js
	// │       │   │       │   ├── clone-array.js
	// │       │   │       │   ├── is-data-uri-resource.js
	// │       │   │       │   ├── natural-compare.js
	// │       │   │       │   ├── is-import.js
	// │       │   │       │   ├── is-remote-resource.js
	// │       │   │       │   ├── split.js
	// │       │   │       │   ├── has-protocol.js
	// │       │   │       │   ├── override.js
	// │       │   │       │   ├── format-position.js
	// │       │   │       │   └── is-https-resource.js
	// │       │   │       └── tokenizer
	// │       │   │           ├── tokenize.js
	// │       │   │           ├── token.js
	// │       │   │           └── marker.js
	// │       │   ├── balanced-match
	// │       │   │   ├── package.json
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   ├── .github
	// │       │   │   │   └── FUNDING.yml
	// │       │   │   └── LICENSE.md
	// │       │   ├── path-parse
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── index.js
	// │       │   ├── braces
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   ├── lib
	// │       │   │   │   ├── utils.js
	// │       │   │   │   ├── parse.js
	// │       │   │   │   ├── expand.js
	// │       │   │   │   ├── compile.js
	// │       │   │   │   ├── constants.js
	// │       │   │   │   └── stringify.js
	// │       │   │   └── CHANGELOG.md
	// │       │   ├── pirates
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── index.d.ts
	// │       │   │   └── lib
	// │       │   │       └── index.js
	// │       │   ├── @alloc
	// │       │   │   └── quick-lru
	// │       │   │       ├── license
	// │       │   │       ├── package.json
	// │       │   │       ├── index.js
	// │       │   │       ├── index.d.ts
	// │       │   │       └── readme.md
	// │       │   ├── ts-interface-checker
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── dist
	// │       │   │       ├── types.d.ts
	// │       │   │       ├── util.d.ts
	// │       │   │       ├── types.js
	// │       │   │       ├── index.js
	// │       │   │       ├── index.d.ts
	// │       │   │       └── util.js
	// │       │   ├── object-hash
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── index.js
	// │       │   │   ├── dist
	// │       │   │   │   └── object_hash.js
	// │       │   │   └── readme.markdown
	// │       │   ├── camelcase-css
	// │       │   │   ├── license
	// │       │   │   ├── package.json
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   └── index-es5.js
	// │       │   ├── is-number
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── index.js
	// │       │   ├── didyoumean
	// │       │   │   ├── didYouMean-1.2.1.js
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── didYouMean-1.2.1.min.js
	// │       │   ├── micromatch
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── index.js
	// │       │   ├── postcss-load-config
	// │       │   │   ├── src
	// │       │   │   │   ├── plugins.js
	// │       │   │   │   ├── req.js
	// │       │   │   │   ├── index.js
	// │       │   │   │   ├── index.d.ts
	// │       │   │   │   └── options.js
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   └── README.md
	// │       │   ├── clean-css-cli
	// │       │   │   ├── History.md
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   └── bin
	// │       │   │       └── cleancss
	// │       │   ├── merge2
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── index.js
	// │       │   ├── lines-and-columns
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── build
	// │       │   │       ├── index.js
	// │       │   │       └── index.d.ts
	// │       │   ├── source-map
	// │       │   │   ├── source-map.d.ts
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── dist
	// │       │   │   │   ├── source-map.min.js
	// │       │   │   │   ├── source-map.js
	// │       │   │   │   ├── source-map.debug.js
	// │       │   │   │   └── source-map.min.js.map
	// │       │   │   ├── lib
	// │       │   │   │   ├── binary-search.js
	// │       │   │   │   ├── source-map-consumer.js
	// │       │   │   │   ├── source-node.js
	// │       │   │   │   ├── base64-vlq.js
	// │       │   │   │   ├── base64.js
	// │       │   │   │   ├── source-map-generator.js
	// │       │   │   │   ├── quick-sort.js
	// │       │   │   │   ├── mapping-list.js
	// │       │   │   │   ├── util.js
	// │       │   │   │   └── array-set.js
	// │       │   │   ├── source-map.js
	// │       │   │   └── CHANGELOG.md
	// │       │   ├── fill-range
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── index.js
	// │       │   ├── glob-parent
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   └── CHANGELOG.md
	// │       │   ├── cssesc
	// │       │   │   ├── man
	// │       │   │   │   └── cssesc.1
	// │       │   │   ├── package.json
	// │       │   │   ├── README.md
	// │       │   │   ├── cssesc.js
	// │       │   │   ├── LICENSE-MIT.txt
	// │       │   │   └── bin
	// │       │   │       └── cssesc
	// │       │   ├── chokidar
	// │       │   │   ├── types
	// │       │   │   │   └── index.d.ts
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   └── lib
	// │       │   │       ├── fsevents-handler.js
	// │       │   │       ├── constants.js
	// │       │   │       └── nodefs-handler.js
	// │       │   ├── minimatch
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── minimatch.js
	// │       │   ├── yaml
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── dist
	// │       │   │   │   ├── compose
	// │       │   │   │   │   ├── resolve-block-map.d.ts
	// │       │   │   │   │   ├── resolve-flow-scalar.d.ts
	// │       │   │   │   │   ├── compose-node.d.ts
	// │       │   │   │   │   ├── compose-doc.js
	// │       │   │   │   │   ├── util-contains-newline.js
	// │       │   │   │   │   ├── compose-scalar.d.ts
	// │       │   │   │   │   ├── util-empty-scalar-position.js
	// │       │   │   │   │   ├── resolve-end.js
	// │       │   │   │   │   ├── util-flow-indent-check.js
	// │       │   │   │   │   ├── compose-doc.d.ts
	// │       │   │   │   │   ├── composer.js
	// │       │   │   │   │   ├── compose-collection.js
	// │       │   │   │   │   ├── composer.d.ts
	// │       │   │   │   │   ├── resolve-flow-collection.js
	// │       │   │   │   │   ├── resolve-props.js
	// │       │   │   │   │   ├── resolve-block-seq.js
	// │       │   │   │   │   ├── resolve-block-map.js
	// │       │   │   │   │   ├── util-contains-newline.d.ts
	// │       │   │   │   │   ├── compose-node.js
	// │       │   │   │   │   ├── util-map-includes.d.ts
	// │       │   │   │   │   ├── util-flow-indent-check.d.ts
	// │       │   │   │   │   ├── resolve-end.d.ts
	// │       │   │   │   │   ├── util-map-includes.js
	// │       │   │   │   │   ├── resolve-block-scalar.js
	// │       │   │   │   │   ├── compose-scalar.js
	// │       │   │   │   │   ├── resolve-flow-collection.d.ts
	// │       │   │   │   │   ├── compose-collection.d.ts
	// │       │   │   │   │   ├── resolve-block-seq.d.ts
	// │       │   │   │   │   ├── resolve-flow-scalar.js
	// │       │   │   │   │   ├── util-empty-scalar-position.d.ts
	// │       │   │   │   │   ├── resolve-props.d.ts
	// │       │   │   │   │   └── resolve-block-scalar.d.ts
	// │       │   │   │   ├── public-api.d.ts
	// │       │   │   │   ├── errors.d.ts
	// │       │   │   │   ├── visit.d.ts
	// │       │   │   │   ├── test-events.d.ts
	// │       │   │   │   ├── doc
	// │       │   │   │   │   ├── createNode.d.ts
	// │       │   │   │   │   ├── createNode.js
	// │       │   │   │   │   ├── anchors.d.ts
	// │       │   │   │   │   ├── directives.js
	// │       │   │   │   │   ├── directives.d.ts
	// │       │   │   │   │   ├── anchors.js
	// │       │   │   │   │   ├── applyReviver.js
	// │       │   │   │   │   ├── Document.js
	// │       │   │   │   │   ├── applyReviver.d.ts
	// │       │   │   │   │   └── Document.d.ts
	// │       │   │   │   ├── util.d.ts
	// │       │   │   │   ├── schema
	// │       │   │   │   │   ├── types.d.ts
	// │       │   │   │   │   ├── tags.d.ts
	// │       │   │   │   │   ├── json-schema.d.ts
	// │       │   │   │   │   ├── tags.js
	// │       │   │   │   │   ├── Schema.js
	// │       │   │   │   │   ├── Schema.d.ts
	// │       │   │   │   │   ├── json
	// │       │   │   │   │   │   ├── schema.js
	// │       │   │   │   │   │   └── schema.d.ts
	// │       │   │   │   │   ├── yaml-1.1
	// │       │   │   │   │   │   ├── binary.js
	// │       │   │   │   │   │   ├── omap.js
	// │       │   │   │   │   │   ├── float.d.ts
	// │       │   │   │   │   │   ├── schema.js
	// │       │   │   │   │   │   ├── float.js
	// │       │   │   │   │   │   ├── int.d.ts
	// │       │   │   │   │   │   ├── omap.d.ts
	// │       │   │   │   │   │   ├── set.d.ts
	// │       │   │   │   │   │   ├── schema.d.ts
	// │       │   │   │   │   │   ├── pairs.js
	// │       │   │   │   │   │   ├── timestamp.d.ts
	// │       │   │   │   │   │   ├── int.js
	// │       │   │   │   │   │   ├── binary.d.ts
	// │       │   │   │   │   │   ├── pairs.d.ts
	// │       │   │   │   │   │   ├── bool.d.ts
	// │       │   │   │   │   │   ├── bool.js
	// │       │   │   │   │   │   ├── set.js
	// │       │   │   │   │   │   └── timestamp.js
	// │       │   │   │   │   ├── common
	// │       │   │   │   │   │   ├── null.d.ts
	// │       │   │   │   │   │   ├── map.d.ts
	// │       │   │   │   │   │   ├── seq.d.ts
	// │       │   │   │   │   │   ├── string.d.ts
	// │       │   │   │   │   │   ├── string.js
	// │       │   │   │   │   │   ├── null.js
	// │       │   │   │   │   │   ├── seq.js
	// │       │   │   │   │   │   └── map.js
	// │       │   │   │   │   └── core
	// │       │   │   │   │       ├── float.d.ts
	// │       │   │   │   │       ├── schema.js
	// │       │   │   │   │       ├── float.js
	// │       │   │   │   │       ├── int.d.ts
	// │       │   │   │   │       ├── schema.d.ts
	// │       │   │   │   │       ├── int.js
	// │       │   │   │   │       ├── bool.d.ts
	// │       │   │   │   │       └── bool.js
	// │       │   │   │   ├── errors.js
	// │       │   │   │   ├── nodes
	// │       │   │   │   │   ├── addPairToJSMap.d.ts
	// │       │   │   │   │   ├── identity.js
	// │       │   │   │   │   ├── toJS.d.ts
	// │       │   │   │   │   ├── YAMLSeq.d.ts
	// │       │   │   │   │   ├── Collection.d.ts
	// │       │   │   │   │   ├── Alias.js
	// │       │   │   │   │   ├── Pair.d.ts
	// │       │   │   │   │   ├── YAMLSeq.js
	// │       │   │   │   │   ├── Node.d.ts
	// │       │   │   │   │   ├── Pair.js
	// │       │   │   │   │   ├── toJS.js
	// │       │   │   │   │   ├── addPairToJSMap.js
	// │       │   │   │   │   ├── Alias.d.ts
	// │       │   │   │   │   ├── Node.js
	// │       │   │   │   │   ├── Scalar.d.ts
	// │       │   │   │   │   ├── YAMLMap.js
	// │       │   │   │   │   ├── identity.d.ts
	// │       │   │   │   │   ├── YAMLMap.d.ts
	// │       │   │   │   │   ├── Collection.js
	// │       │   │   │   │   └── Scalar.js
	// │       │   │   │   ├── index.js
	// │       │   │   │   ├── index.d.ts
	// │       │   │   │   ├── log.d.ts
	// │       │   │   │   ├── log.js
	// │       │   │   │   ├── visit.js
	// │       │   │   │   ├── test-events.js
	// │       │   │   │   ├── parse
	// │       │   │   │   │   ├── cst-scalar.js
	// │       │   │   │   │   ├── cst-scalar.d.ts
	// │       │   │   │   │   ├── parser.js
	// │       │   │   │   │   ├── cst-stringify.js
	// │       │   │   │   │   ├── parser.d.ts
	// │       │   │   │   │   ├── cst-visit.js
	// │       │   │   │   │   ├── cst.js
	// │       │   │   │   │   ├── line-counter.js
	// │       │   │   │   │   ├── line-counter.d.ts
	// │       │   │   │   │   ├── cst-visit.d.ts
	// │       │   │   │   │   ├── cst.d.ts
	// │       │   │   │   │   ├── lexer.d.ts
	// │       │   │   │   │   ├── cst-stringify.d.ts
	// │       │   │   │   │   └── lexer.js
	// │       │   │   │   ├── public-api.js
	// │       │   │   │   ├── options.d.ts
	// │       │   │   │   ├── stringify
	// │       │   │   │   │   ├── foldFlowLines.d.ts
	// │       │   │   │   │   ├── stringify.d.ts
	// │       │   │   │   │   ├── stringifyCollection.d.ts
	// │       │   │   │   │   ├── stringifyNumber.js
	// │       │   │   │   │   ├── stringifyCollection.js
	// │       │   │   │   │   ├── stringifyComment.d.ts
	// │       │   │   │   │   ├── stringifyString.d.ts
	// │       │   │   │   │   ├── foldFlowLines.js
	// │       │   │   │   │   ├── stringifyPair.js
	// │       │   │   │   │   ├── stringifyDocument.d.ts
	// │       │   │   │   │   ├── stringifyDocument.js
	// │       │   │   │   │   ├── stringifyComment.js
	// │       │   │   │   │   ├── stringify.js
	// │       │   │   │   │   ├── stringifyPair.d.ts
	// │       │   │   │   │   ├── stringifyString.js
	// │       │   │   │   │   └── stringifyNumber.d.ts
	// │       │   │   │   └── util.js
	// │       │   │   ├── browser
	// │       │   │   │   ├── package.json
	// │       │   │   │   ├── index.js
	// │       │   │   │   └── dist
	// │       │   │   │       ├── node_modules
	// │       │   │   │       │   └── tslib
	// │       │   │   │       │       └── tslib.es6.js
	// │       │   │   │       ├── compose
	// │       │   │   │       │   ├── compose-doc.js
	// │       │   │   │       │   ├── util-contains-newline.js
	// │       │   │   │       │   ├── util-empty-scalar-position.js
	// │       │   │   │       │   ├── resolve-end.js
	// │       │   │   │       │   ├── util-flow-indent-check.js
	// │       │   │   │       │   ├── composer.js
	// │       │   │   │       │   ├── compose-collection.js
	// │       │   │   │       │   ├── resolve-flow-collection.js
	// │       │   │   │       │   ├── resolve-props.js
	// │       │   │   │       │   ├── resolve-block-seq.js
	// │       │   │   │       │   ├── resolve-block-map.js
	// │       │   │   │       │   ├── compose-node.js
	// │       │   │   │       │   ├── util-map-includes.js
	// │       │   │   │       │   ├── resolve-block-scalar.js
	// │       │   │   │       │   ├── compose-scalar.js
	// │       │   │   │       │   └── resolve-flow-scalar.js
	// │       │   │   │       ├── doc
	// │       │   │   │       │   ├── createNode.js
	// │       │   │   │       │   ├── directives.js
	// │       │   │   │       │   ├── anchors.js
	// │       │   │   │       │   ├── applyReviver.js
	// │       │   │   │       │   └── Document.js
	// │       │   │   │       ├── schema
	// │       │   │   │       │   ├── tags.js
	// │       │   │   │       │   ├── Schema.js
	// │       │   │   │       │   ├── json
	// │       │   │   │       │   │   └── schema.js
	// │       │   │   │       │   ├── yaml-1.1
	// │       │   │   │       │   │   ├── binary.js
	// │       │   │   │       │   │   ├── omap.js
	// │       │   │   │       │   │   ├── schema.js
	// │       │   │   │       │   │   ├── float.js
	// │       │   │   │       │   │   ├── pairs.js
	// │       │   │   │       │   │   ├── int.js
	// │       │   │   │       │   │   ├── bool.js
	// │       │   │   │       │   │   ├── set.js
	// │       │   │   │       │   │   └── timestamp.js
	// │       │   │   │       │   ├── common
	// │       │   │   │       │   │   ├── string.js
	// │       │   │   │       │   │   ├── null.js
	// │       │   │   │       │   │   ├── seq.js
	// │       │   │   │       │   │   └── map.js
	// │       │   │   │       │   └── core
	// │       │   │   │       │       ├── schema.js
	// │       │   │   │       │       ├── float.js
	// │       │   │   │       │       ├── int.js
	// │       │   │   │       │       └── bool.js
	// │       │   │   │       ├── errors.js
	// │       │   │   │       ├── nodes
	// │       │   │   │       │   ├── identity.js
	// │       │   │   │       │   ├── Alias.js
	// │       │   │   │       │   ├── YAMLSeq.js
	// │       │   │   │       │   ├── Pair.js
	// │       │   │   │       │   ├── toJS.js
	// │       │   │   │       │   ├── addPairToJSMap.js
	// │       │   │   │       │   ├── Node.js
	// │       │   │   │       │   ├── YAMLMap.js
	// │       │   │   │       │   ├── Collection.js
	// │       │   │   │       │   └── Scalar.js
	// │       │   │   │       ├── index.js
	// │       │   │   │       ├── log.js
	// │       │   │   │       ├── visit.js
	// │       │   │   │       ├── parse
	// │       │   │   │       │   ├── cst-scalar.js
	// │       │   │   │       │   ├── parser.js
	// │       │   │   │       │   ├── cst-stringify.js
	// │       │   │   │       │   ├── cst-visit.js
	// │       │   │   │       │   ├── cst.js
	// │       │   │   │       │   ├── line-counter.js
	// │       │   │   │       │   └── lexer.js
	// │       │   │   │       ├── public-api.js
	// │       │   │   │       ├── stringify
	// │       │   │   │       │   ├── stringifyNumber.js
	// │       │   │   │       │   ├── stringifyCollection.js
	// │       │   │   │       │   ├── foldFlowLines.js
	// │       │   │   │       │   ├── stringifyPair.js
	// │       │   │   │       │   ├── stringifyDocument.js
	// │       │   │   │       │   ├── stringifyComment.js
	// │       │   │   │       │   ├── stringify.js
	// │       │   │   │       │   └── stringifyString.js
	// │       │   │   │       └── util.js
	// │       │   │   └── util.js
	// │       │   ├── brace-expansion
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── index.js
	// │       │   ├── is-binary-path
	// │       │   │   ├── license
	// │       │   │   ├── package.json
	// │       │   │   ├── index.js
	// │       │   │   ├── index.d.ts
	// │       │   │   └── readme.md
	// │       │   ├── tailwindcss
	// │       │   │   ├── node_modules
	// │       │   │   │   └── glob-parent
	// │       │   │   │       ├── package.json
	// │       │   │   │       ├── LICENSE
	// │       │   │   │       ├── README.md
	// │       │   │   │       └── index.js
	// │       │   │   ├── components.css
	// │       │   │   ├── src
	// │       │   │   │   ├── corePlugins.js
	// │       │   │   │   ├── cli
	// │       │   │   │   │   ├── index.js
	// │       │   │   │   │   ├── build
	// │       │   │   │   │   │   ├── index.js
	// │       │   │   │   │   │   ├── utils.js
	// │       │   │   │   │   │   ├── plugin.js
	// │       │   │   │   │   │   ├── deps.js
	// │       │   │   │   │   │   └── watching.js
	// │       │   │   │   │   ├── init
	// │       │   │   │   │   │   └── index.js
	// │       │   │   │   │   └── help
	// │       │   │   │   │       └── index.js
	// │       │   │   │   ├── corePluginList.js
	// │       │   │   │   ├── cli-peer-dependencies.js
	// │       │   │   │   ├── index.js
	// │       │   │   │   ├── processTailwindFeatures.js
	// │       │   │   │   ├── oxide
	// │       │   │   │   │   ├── cli
	// │       │   │   │   │   │   ├── build
	// │       │   │   │   │   │   │   ├── plugin.ts
	// │       │   │   │   │   │   │   ├── watching.ts
	// │       │   │   │   │   │   │   ├── utils.ts
	// │       │   │   │   │   │   │   ├── deps.ts
	// │       │   │   │   │   │   │   └── index.ts
	// │       │   │   │   │   │   ├── init
	// │       │   │   │   │   │   │   └── index.ts
	// │       │   │   │   │   │   ├── help
	// │       │   │   │   │   │   │   └── index.ts
	// │       │   │   │   │   │   └── index.ts
	// │       │   │   │   │   ├── postcss-plugin.ts
	// │       │   │   │   │   └── cli.ts
	// │       │   │   │   ├── lib
	// │       │   │   │   │   ├── partitionApplyAtRules.js
	// │       │   │   │   │   ├── setupContextUtils.js
	// │       │   │   │   │   ├── setupTrackingContext.js
	// │       │   │   │   │   ├── resolveDefaultsAtRules.js
	// │       │   │   │   │   ├── generateRules.js
	// │       │   │   │   │   ├── expandApplyAtRules.js
	// │       │   │   │   │   ├── collapseDuplicateDeclarations.js
	// │       │   │   │   │   ├── defaultExtractor.js
	// │       │   │   │   │   ├── expandTailwindAtRules.js
	// │       │   │   │   │   ├── substituteScreenAtRules.js
	// │       │   │   │   │   ├── regex.js
	// │       │   │   │   │   ├── collapseAdjacentRules.js
	// │       │   │   │   │   ├── sharedState.js
	// │       │   │   │   │   ├── evaluateTailwindFunctions.js
	// │       │   │   │   │   ├── cacheInvalidation.js
	// │       │   │   │   │   ├── normalizeTailwindDirectives.js
	// │       │   │   │   │   ├── remap-bitfield.js
	// │       │   │   │   │   ├── getModuleDependencies.js
	// │       │   │   │   │   ├── offsets.js
	// │       │   │   │   │   ├── detectNesting.js
	// │       │   │   │   │   ├── load-config.ts
	// │       │   │   │   │   ├── content.js
	// │       │   │   │   │   └── findAtConfigPath.js
	// │       │   │   │   ├── plugin.js
	// │       │   │   │   ├── util
	// │       │   │   │   │   ├── parseObjectStyles.js
	// │       │   │   │   │   ├── color.js
	// │       │   │   │   │   ├── resolveConfigPath.js
	// │       │   │   │   │   ├── dataTypes.js
	// │       │   │   │   │   ├── parseGlob.js
	// │       │   │   │   │   ├── validateFormalSyntax.js
	// │       │   │   │   │   ├── isPlainObject.js
	// │       │   │   │   │   ├── defaults.js
	// │       │   │   │   │   ├── responsive.js
	// │       │   │   │   │   ├── getAllConfigs.js
	// │       │   │   │   │   ├── negateValue.js
	// │       │   │   │   │   ├── pseudoElements.js
	// │       │   │   │   │   ├── configurePlugins.js
	// │       │   │   │   │   ├── cloneNodes.js
	// │       │   │   │   │   ├── createPlugin.js
	// │       │   │   │   │   ├── withAlphaVariable.js
	// │       │   │   │   │   ├── bigSign.js
	// │       │   │   │   │   ├── splitAtTopLevelOnly.js
	// │       │   │   │   │   ├── resolveConfig.js
	// │       │   │   │   │   ├── nameClass.js
	// │       │   │   │   │   ├── toColorValue.js
	// │       │   │   │   │   ├── isSyntacticallyValidPropertyValue.js
	// │       │   │   │   │   ├── formatVariantSelector.js
	// │       │   │   │   │   ├── log.js
	// │       │   │   │   │   ├── prefixSelector.js
	// │       │   │   │   │   ├── cloneDeep.js
	// │       │   │   │   │   ├── toPath.js
	// │       │   │   │   │   ├── buildMediaQuery.js
	// │       │   │   │   │   ├── parseAnimationValue.js
	// │       │   │   │   │   ├── createUtilityPlugin.js
	// │       │   │   │   │   ├── isKeyframeRule.js
	// │       │   │   │   │   ├── normalizeConfig.js
	// │       │   │   │   │   ├── flattenColorPalette.js
	// │       │   │   │   │   ├── applyImportantSelector.js
	// │       │   │   │   │   ├── hashConfig.js
	// │       │   │   │   │   ├── transformThemeValue.js
	// │       │   │   │   │   ├── escapeCommas.js
	// │       │   │   │   │   ├── validateConfig.js
	// │       │   │   │   │   ├── removeAlphaVariables.js
	// │       │   │   │   │   ├── escapeClassName.js
	// │       │   │   │   │   ├── parseBoxShadowValue.js
	// │       │   │   │   │   ├── parseDependency.js
	// │       │   │   │   │   ├── colorNames.js
	// │       │   │   │   │   ├── tap.js
	// │       │   │   │   │   ├── pluginUtils.js
	// │       │   │   │   │   └── normalizeScreens.js
	// │       │   │   │   ├── css
	// │       │   │   │   │   ├── preflight.css
	// │       │   │   │   │   └── LICENSE
	// │       │   │   │   ├── featureFlags.js
	// │       │   │   │   ├── cli.js
	// │       │   │   │   ├── public
	// │       │   │   │   │   ├── load-config.js
	// │       │   │   │   │   ├── default-config.js
	// │       │   │   │   │   ├── colors.js
	// │       │   │   │   │   ├── resolve-config.js
	// │       │   │   │   │   ├── default-theme.js
	// │       │   │   │   │   └── create-plugin.js
	// │       │   │   │   └── postcss-plugins
	// │       │   │   │       └── nesting
	// │       │   │   │           ├── README.md
	// │       │   │   │           ├── index.js
	// │       │   │   │           └── plugin.js
	// │       │   │   ├── resolveConfig.d.ts
	// │       │   │   ├── types
	// │       │   │   │   ├── config.d.ts
	// │       │   │   │   ├── generated
	// │       │   │   │   │   ├── .gitkeep
	// │       │   │   │   │   ├── corePluginList.d.ts
	// │       │   │   │   │   ├── default-theme.d.ts
	// │       │   │   │   │   └── colors.d.ts
	// │       │   │   │   └── index.d.ts
	// │       │   │   ├── package.json
	// │       │   │   ├── base.css
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── resolveConfig.js
	// │       │   │   ├── colors.js
	// │       │   │   ├── variants.css
	// │       │   │   ├── scripts
	// │       │   │   │   ├── release-channel.js
	// │       │   │   │   ├── type-utils.js
	// │       │   │   │   ├── swap-engines.js
	// │       │   │   │   ├── generate-types.js
	// │       │   │   │   ├── create-plugin-list.js
	// │       │   │   │   └── release-notes.js
	// │       │   │   ├── prettier.config.js
	// │       │   │   ├── defaultTheme.d.ts
	// │       │   │   ├── defaultTheme.js
	// │       │   │   ├── lib
	// │       │   │   │   ├── corePlugins.js
	// │       │   │   │   ├── cli
	// │       │   │   │   │   ├── index.js
	// │       │   │   │   │   ├── build
	// │       │   │   │   │   │   ├── index.js
	// │       │   │   │   │   │   ├── utils.js
	// │       │   │   │   │   │   ├── plugin.js
	// │       │   │   │   │   │   ├── deps.js
	// │       │   │   │   │   │   └── watching.js
	// │       │   │   │   │   ├── init
	// │       │   │   │   │   │   └── index.js
	// │       │   │   │   │   └── help
	// │       │   │   │   │       └── index.js
	// │       │   │   │   ├── corePluginList.js
	// │       │   │   │   ├── cli-peer-dependencies.js
	// │       │   │   │   ├── index.js
	// │       │   │   │   ├── processTailwindFeatures.js
	// │       │   │   │   ├── oxide
	// │       │   │   │   │   ├── cli
	// │       │   │   │   │   │   ├── index.js
	// │       │   │   │   │   │   ├── build
	// │       │   │   │   │   │   │   ├── index.js
	// │       │   │   │   │   │   │   ├── utils.js
	// │       │   │   │   │   │   │   ├── plugin.js
	// │       │   │   │   │   │   │   ├── deps.js
	// │       │   │   │   │   │   │   └── watching.js
	// │       │   │   │   │   │   ├── init
	// │       │   │   │   │   │   │   └── index.js
	// │       │   │   │   │   │   └── help
	// │       │   │   │   │   │       └── index.js
	// │       │   │   │   │   ├── cli.js
	// │       │   │   │   │   └── postcss-plugin.js
	// │       │   │   │   ├── lib
	// │       │   │   │   │   ├── partitionApplyAtRules.js
	// │       │   │   │   │   ├── setupContextUtils.js
	// │       │   │   │   │   ├── setupTrackingContext.js
	// │       │   │   │   │   ├── resolveDefaultsAtRules.js
	// │       │   │   │   │   ├── load-config.js
	// │       │   │   │   │   ├── generateRules.js
	// │       │   │   │   │   ├── expandApplyAtRules.js
	// │       │   │   │   │   ├── collapseDuplicateDeclarations.js
	// │       │   │   │   │   ├── defaultExtractor.js
	// │       │   │   │   │   ├── expandTailwindAtRules.js
	// │       │   │   │   │   ├── substituteScreenAtRules.js
	// │       │   │   │   │   ├── regex.js
	// │       │   │   │   │   ├── collapseAdjacentRules.js
	// │       │   │   │   │   ├── sharedState.js
	// │       │   │   │   │   ├── evaluateTailwindFunctions.js
	// │       │   │   │   │   ├── cacheInvalidation.js
	// │       │   │   │   │   ├── normalizeTailwindDirectives.js
	// │       │   │   │   │   ├── remap-bitfield.js
	// │       │   │   │   │   ├── getModuleDependencies.js
	// │       │   │   │   │   ├── offsets.js
	// │       │   │   │   │   ├── detectNesting.js
	// │       │   │   │   │   ├── content.js
	// │       │   │   │   │   └── findAtConfigPath.js
	// │       │   │   │   ├── plugin.js
	// │       │   │   │   ├── util
	// │       │   │   │   │   ├── parseObjectStyles.js
	// │       │   │   │   │   ├── color.js
	// │       │   │   │   │   ├── resolveConfigPath.js
	// │       │   │   │   │   ├── dataTypes.js
	// │       │   │   │   │   ├── parseGlob.js
	// │       │   │   │   │   ├── validateFormalSyntax.js
	// │       │   │   │   │   ├── isPlainObject.js
	// │       │   │   │   │   ├── defaults.js
	// │       │   │   │   │   ├── responsive.js
	// │       │   │   │   │   ├── getAllConfigs.js
	// │       │   │   │   │   ├── negateValue.js
	// │       │   │   │   │   ├── pseudoElements.js
	// │       │   │   │   │   ├── configurePlugins.js
	// │       │   │   │   │   ├── cloneNodes.js
	// │       │   │   │   │   ├── createPlugin.js
	// │       │   │   │   │   ├── withAlphaVariable.js
	// │       │   │   │   │   ├── bigSign.js
	// │       │   │   │   │   ├── splitAtTopLevelOnly.js
	// │       │   │   │   │   ├── resolveConfig.js
	// │       │   │   │   │   ├── nameClass.js
	// │       │   │   │   │   ├── toColorValue.js
	// │       │   │   │   │   ├── isSyntacticallyValidPropertyValue.js
	// │       │   │   │   │   ├── formatVariantSelector.js
	// │       │   │   │   │   ├── log.js
	// │       │   │   │   │   ├── prefixSelector.js
	// │       │   │   │   │   ├── cloneDeep.js
	// │       │   │   │   │   ├── toPath.js
	// │       │   │   │   │   ├── buildMediaQuery.js
	// │       │   │   │   │   ├── parseAnimationValue.js
	// │       │   │   │   │   ├── createUtilityPlugin.js
	// │       │   │   │   │   ├── isKeyframeRule.js
	// │       │   │   │   │   ├── normalizeConfig.js
	// │       │   │   │   │   ├── flattenColorPalette.js
	// │       │   │   │   │   ├── applyImportantSelector.js
	// │       │   │   │   │   ├── hashConfig.js
	// │       │   │   │   │   ├── transformThemeValue.js
	// │       │   │   │   │   ├── escapeCommas.js
	// │       │   │   │   │   ├── validateConfig.js
	// │       │   │   │   │   ├── removeAlphaVariables.js
	// │       │   │   │   │   ├── escapeClassName.js
	// │       │   │   │   │   ├── parseBoxShadowValue.js
	// │       │   │   │   │   ├── parseDependency.js
	// │       │   │   │   │   ├── colorNames.js
	// │       │   │   │   │   ├── tap.js
	// │       │   │   │   │   ├── pluginUtils.js
	// │       │   │   │   │   └── normalizeScreens.js
	// │       │   │   │   ├── css
	// │       │   │   │   │   ├── preflight.css
	// │       │   │   │   │   └── LICENSE
	// │       │   │   │   ├── featureFlags.js
	// │       │   │   │   ├── cli.js
	// │       │   │   │   ├── public
	// │       │   │   │   │   ├── load-config.js
	// │       │   │   │   │   ├── default-config.js
	// │       │   │   │   │   ├── colors.js
	// │       │   │   │   │   ├── resolve-config.js
	// │       │   │   │   │   ├── default-theme.js
	// │       │   │   │   │   └── create-plugin.js
	// │       │   │   │   └── postcss-plugins
	// │       │   │   │       └── nesting
	// │       │   │   │           ├── README.md
	// │       │   │   │           ├── index.js
	// │       │   │   │           └── plugin.js
	// │       │   │   ├── tailwind.css
	// │       │   │   ├── loadConfig.d.ts
	// │       │   │   ├── loadConfig.js
	// │       │   │   ├── plugin.js
	// │       │   │   ├── stubs
	// │       │   │   │   ├── config.full.js
	// │       │   │   │   ├── config.simple.js
	// │       │   │   │   ├── tailwind.config.js
	// │       │   │   │   ├── postcss.config.js
	// │       │   │   │   ├── postcss.config.cjs
	// │       │   │   │   ├── .prettierrc.json
	// │       │   │   │   ├── tailwind.config.ts
	// │       │   │   │   ├── .npmignore
	// │       │   │   │   └── tailwind.config.cjs
	// │       │   │   ├── screens.css
	// │       │   │   ├── nesting
	// │       │   │   │   └── index.js
	// │       │   │   ├── defaultConfig.js
	// │       │   │   ├── defaultConfig.d.ts
	// │       │   │   ├── colors.d.ts
	// │       │   │   ├── peers
	// │       │   │   │   └── index.js
	// │       │   │   ├── utilities.css
	// │       │   │   ├── plugin.d.ts
	// │       │   │   └── CHANGELOG.md
	// │       │   ├── wrappy
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── wrappy.js
	// │       │   ├── @nodelib
	// │       │   │   ├── fs.walk
	// │       │   │   │   ├── package.json
	// │       │   │   │   ├── LICENSE
	// │       │   │   │   ├── README.md
	// │       │   │   │   └── out
	// │       │   │   │       ├── types
	// │       │   │   │       │   ├── index.js
	// │       │   │   │       │   └── index.d.ts
	// │       │   │   │       ├── index.js
	// │       │   │   │       ├── index.d.ts
	// │       │   │   │       ├── settings.d.ts
	// │       │   │   │       ├── providers
	// │       │   │   │       │   ├── stream.d.ts
	// │       │   │   │       │   ├── index.js
	// │       │   │   │       │   ├── index.d.ts
	// │       │   │   │       │   ├── async.d.ts
	// │       │   │   │       │   ├── sync.d.ts
	// │       │   │   │       │   ├── sync.js
	// │       │   │   │       │   ├── stream.js
	// │       │   │   │       │   └── async.js
	// │       │   │   │       ├── settings.js
	// │       │   │   │       └── readers
	// │       │   │   │           ├── reader.js
	// │       │   │   │           ├── common.d.ts
	// │       │   │   │           ├── async.d.ts
	// │       │   │   │           ├── reader.d.ts
	// │       │   │   │           ├── sync.d.ts
	// │       │   │   │           ├── sync.js
	// │       │   │   │           ├── common.js
	// │       │   │   │           └── async.js
	// │       │   │   ├── fs.scandir
	// │       │   │   │   ├── package.json
	// │       │   │   │   ├── LICENSE
	// │       │   │   │   ├── README.md
	// │       │   │   │   └── out
	// │       │   │   │       ├── adapters
	// │       │   │   │       │   ├── fs.js
	// │       │   │   │       │   └── fs.d.ts
	// │       │   │   │       ├── types
	// │       │   │   │       │   ├── index.js
	// │       │   │   │       │   └── index.d.ts
	// │       │   │   │       ├── index.js
	// │       │   │   │       ├── index.d.ts
	// │       │   │   │       ├── settings.d.ts
	// │       │   │   │       ├── providers
	// │       │   │   │       │   ├── common.d.ts
	// │       │   │   │       │   ├── async.d.ts
	// │       │   │   │       │   ├── sync.d.ts
	// │       │   │   │       │   ├── sync.js
	// │       │   │   │       │   ├── common.js
	// │       │   │   │       │   └── async.js
	// │       │   │   │       ├── constants.d.ts
	// │       │   │   │       ├── utils
	// │       │   │   │       │   ├── index.js
	// │       │   │   │       │   ├── index.d.ts
	// │       │   │   │       │   ├── fs.js
	// │       │   │   │       │   └── fs.d.ts
	// │       │   │   │       ├── settings.js
	// │       │   │   │       └── constants.js
	// │       │   │   └── fs.stat
	// │       │   │       ├── package.json
	// │       │   │       ├── LICENSE
	// │       │   │       ├── README.md
	// │       │   │       └── out
	// │       │   │           ├── adapters
	// │       │   │           │   ├── fs.js
	// │       │   │           │   └── fs.d.ts
	// │       │   │           ├── types
	// │       │   │           │   ├── index.js
	// │       │   │           │   └── index.d.ts
	// │       │   │           ├── index.js
	// │       │   │           ├── index.d.ts
	// │       │   │           ├── settings.d.ts
	// │       │   │           ├── providers
	// │       │   │           │   ├── async.d.ts
	// │       │   │           │   ├── sync.d.ts
	// │       │   │           │   ├── sync.js
	// │       │   │           │   └── async.js
	// │       │   │           └── settings.js
	// │       │   ├── inherits
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── inherits_browser.js
	// │       │   │   └── inherits.js
	// │       │   ├── postcss-js
	// │       │   │   ├── parser.js
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── process-result.js
	// │       │   │   ├── index.js
	// │       │   │   ├── objectifier.js
	// │       │   │   ├── index.mjs
	// │       │   │   ├── sync.js
	// │       │   │   └── async.js
	// │       │   ├── picocolors
	// │       │   │   ├── picocolors.js
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── picocolors.d.ts
	// │       │   │   ├── picocolors.browser.js
	// │       │   │   └── types.ts
	// │       │   ├── util-deprecate
	// │       │   │   ├── History.md
	// │       │   │   ├── browser.js
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── node.js
	// │       │   ├── supports-preserve-symlinks-flag
	// │       │   │   ├── browser.js
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   ├── .github
	// │       │   │   │   └── FUNDING.yml
	// │       │   │   ├── .nycrc
	// │       │   │   ├── test
	// │       │   │   │   └── index.js
	// │       │   │   ├── .eslintrc
	// │       │   │   └── CHANGELOG.md
	// │       │   ├── postcss
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── lib
	// │       │   │       ├── processor.d.ts
	// │       │   │       ├── no-work-result.d.ts
	// │       │   │       ├── stringifier.d.ts
	// │       │   │       ├── lazy-result.d.ts
	// │       │   │       ├── stringify.d.ts
	// │       │   │       ├── list.js
	// │       │   │       ├── result.js
	// │       │   │       ├── parser.js
	// │       │   │       ├── tokenize.js
	// │       │   │       ├── declaration.js
	// │       │   │       ├── postcss.js
	// │       │   │       ├── root.js
	// │       │   │       ├── comment.d.ts
	// │       │   │       ├── fromJSON.js
	// │       │   │       ├── previous-map.d.ts
	// │       │   │       ├── comment.js
	// │       │   │       ├── container.js
	// │       │   │       ├── fromJSON.d.ts
	// │       │   │       ├── postcss.d.ts
	// │       │   │       ├── node.js
	// │       │   │       ├── input.js
	// │       │   │       ├── parse.js
	// │       │   │       ├── warn-once.js
	// │       │   │       ├── at-rule.d.ts
	// │       │   │       ├── document.d.ts
	// │       │   │       ├── list.d.ts
	// │       │   │       ├── symbols.js
	// │       │   │       ├── document.js
	// │       │   │       ├── result.d.ts
	// │       │   │       ├── container.d.ts
	// │       │   │       ├── postcss.mjs
	// │       │   │       ├── map-generator.js
	// │       │   │       ├── rule.d.ts
	// │       │   │       ├── postcss.d.mts
	// │       │   │       ├── parse.d.ts
	// │       │   │       ├── root.d.ts
	// │       │   │       ├── declaration.d.ts
	// │       │   │       ├── css-syntax-error.js
	// │       │   │       ├── rule.js
	// │       │   │       ├── css-syntax-error.d.ts
	// │       │   │       ├── previous-map.js
	// │       │   │       ├── input.d.ts
	// │       │   │       ├── processor.js
	// │       │   │       ├── at-rule.js
	// │       │   │       ├── lazy-result.js
	// │       │   │       ├── stringify.js
	// │       │   │       ├── stringifier.js
	// │       │   │       ├── node.d.ts
	// │       │   │       ├── terminal-highlight.js
	// │       │   │       ├── no-work-result.js
	// │       │   │       ├── warning.js
	// │       │   │       └── warning.d.ts
	// │       │   ├── to-regex-range
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── index.js
	// │       │   └── postcss-nested
	// │       │       ├── package.json
	// │       │       ├── LICENSE
	// │       │       ├── README.md
	// │       │       ├── index.js
	// │       │       └── index.d.ts
	// │       ├── package-lock.json
	// │       ├── .gitignore
	// │       ├── package.json
	// │       ├── README.md
	// │       ├── main.css
	// │       ├── tailwind.config.js
	// │       ├── robots.txt
	// │       ├── wasm_exec.js
	// │       ├── main.js
	// │       ├── main.wasm
	// │       ├── sitemap.xml
	// │       ├── main.go
	// │       ├── tailwind_base.css
	// │       ├── toast.css
	// │       ├── service_worker.js
	// │       ├── Makefile
	// │       ├── confirm.sh
	// │       ├── index.html
	// │       ├── tab.js
	// │       └── toast.js
	// ├── process.svg
	// ├── node_test.go
	// ├── go.mod
	// ├── .vscode
	// │   ├── settings.json
	// │   ├── dryrun.log
	// │   ├── configurationCache.log
	// │   └── targets.log
	// ├── Dockerfile
	// ├── process.pu
	// ├── wasm_root_generator.go
	// ├── Makefile
	// ├── wasm_tree_grower.go
	// ├── .goreleaser.yml
	// ├── web.png
	// ├── go.sum
	// ├── wasm_tree_spreader.go
	// ├── tree_handler_programmably_output_test.go
	// ├── tree_handler_output_test.go
	// ├── doc.go
	// └── export_test.go
}
