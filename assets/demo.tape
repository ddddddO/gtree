Output assets/demo.gif

Set Theme "IC_Green_PPL"
Set Shell "bash"
Set FontSize 32
Set Width 1500
Set Height 1100
Set TypingSpeed 150ms

Type "gtree template"
Sleep 2s
Enter
Sleep 2s
Hide
Type "gtree template"
Show
Sleep 2s
Type " | gtree output"
Sleep 2s
Enter
Sleep 5s
Hide
Type "clear" Enter
Show

Type "cat testdata/sample2.md"
Sleep 2s
Enter
Sleep 4s
Hide
Type "cat testdata/sample2.md"
Show
Sleep 2s
Type " | gtree output"
Sleep 2s
Enter
Sleep 6s
Hide
Type "clear" Enter
Show

Type "gtree template"
Sleep 2s
Enter
Sleep 2s
Hide
Type "gtree template"
Show
Sleep 2s
Type " | gtree mkdir --dry-run"
Sleep 2s
Enter
Sleep 6s
Hide
Type "clear" Enter
Show

Type "gtree template | gtree mkdir --dry-run -e .go -e .md -e Makefile"
Sleep 2s
Enter
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

Type "tree gtree --dirsfirst"
Sleep 2s
Enter
Sleep 8s
Hide
Type "rm -rf gtree/" Enter
Show