Output assets/demo_ogp.gif

Set Theme "IC_Green_PPL"
Set Shell "bash"
Set FontSize 18
Set Width 1280
Set Height 640
Set TypingSpeed 40ms

Type "gtree template"
Sleep 1s
Enter
Sleep 1s
Hide
Type "gtree template"
Show
Sleep 1s
Type " | gtree output"
Sleep 1s
Enter
Sleep 2s
Hide
Type "clear" Enter
Show

Type "cat testdata/sample2.md"
Sleep 1s
Enter
Sleep 2s
Hide
Type "cat testdata/sample2.md"
Show
Sleep 1s
Type " | gtree output"
Sleep 1s
Enter
Sleep 2s
Hide
Type "clear" Enter
Show

Type "gtree template"
Sleep 1s
Enter
Sleep 1s
Hide
Type "gtree template"
Show
Sleep 1s
Type " | gtree mkdir --dry-run"
Sleep 1s
Enter
Sleep 2s
Hide
Type "clear" Enter
Show

Type "gtree template | gtree mkdir --dry-run -e .go -e .md -e Makefile"
Sleep 1s
Enter
Sleep 1s

Hide
Type "gtree template | gtree mkdir --dry-run -e .go -e .md -e Makefile"
Show
Sleep 1s
Left 26
Sleep 1s
Backspace 10
Sleep 1s
Enter
Sleep 1s

Type "tree gtree --dirsfirst"
Sleep 1s
Enter
Sleep 2.5s
Hide
Type "rm -rf gtree/" Enter
Show