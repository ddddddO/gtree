Output demo.gif

Set Theme "IC_Green_PPL"
Set Shell "bash"
Set FontSize 32
Set Width 1500
Set Height 1100
Set TypingSpeed 75ms

Type "gtree template" Enter
Sleep 2s
Type "gtree template | gtree output" Enter
Sleep 5s
Hide
Type "clear" Enter
Show

Type "cat testdata/sample4.md" Enter
Sleep 2s
Type "cat testdata/sample4.md | gtree output --two-spaces" Enter
Sleep 5s
Hide
Type "clear" Enter
Show

Type "cat testdata/sample5.md" Enter
Sleep 2s
Type "cat testdata/sample5.md | gtree output --four-spaces" Enter
Sleep 5s
Hide
Type "clear" Enter
Show

Type "gtree template" Enter
Sleep 2s
Type "gtree template | gtree mkdir --dry-run" Enter
Sleep 5s
Hide
Type "clear" Enter
Show

Type "gtree template | gtree mkdir --dry-run -e .go -e .md -e Makefile" Enter
Sleep 2s

Hide
Type "gtree template | gtree mkdir --dry-run -e .go -e .md -e Makefile"
Show
Sleep 2s
Left 26
Sleep 2s
Backspace 10
Sleep 2s
Enter
Sleep 2s

Type "tree gtree" Enter
Sleep 8s
Hide
Type "rm -rf gtree/" Enter
Show