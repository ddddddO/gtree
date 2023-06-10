package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/ddddddO/gtree"
)

// Exapmle:
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
func main() {
	var (
		root *gtree.Node
		node *gtree.Node
	)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		splited := strings.Split(line, "/")

		for i, s := range splited {
			if i == 0 {
				if root == nil {
					root = gtree.NewRoot(s)
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
	// │       │   ├── path-is-absolute
	// │       │   │   ├── license
	// │       │   │   ├── package.json
	// │       │   │   ├── index.js
	// │       │   │   └── readme.md
	// │       │   ├── readdirp
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   └── index.d.ts
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
	// │       │   ├── once
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── once.js
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
	// │       │   ├── glob
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── glob.js
	// │       │   │   ├── sync.js
	// │       │   │   └── common.js
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
	// │       │   ├── is-number
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── index.js
	// │       │   ├── clean-css-cli
	// │       │   │   ├── History.md
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── index.js
	// │       │   │   └── bin
	// │       │   │       └── cleancss
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
	// │       │   ├── wrappy
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   └── wrappy.js
	// │       │   ├── inherits
	// │       │   │   ├── package.json
	// │       │   │   ├── LICENSE
	// │       │   │   ├── README.md
	// │       │   │   ├── inherits_browser.js
	// │       │   │   └── inherits.js
	// │       │   └── to-regex-range
	// │       │       ├── package.json
	// │       │       ├── LICENSE
	// │       │       ├── README.md
	// │       │       └── index.js
	// │       ├── package-lock.json
	// │       ├── .gitignore
	// │       ├── package.json
	// │       ├── README.md
	// │       ├── main.css
	// │       ├── tailwind.config.js
	// │       ├── robots.txt
	// │       ├── wasm_exec.js
	// │       ├── main.js
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
