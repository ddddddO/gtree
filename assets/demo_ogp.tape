Output assets/demo_ogp.gif

Set Theme "IC_Green_PPL"
Set Shell "bash"
Set FontSize 18
Set Width 1280
Set Height 640
Set TypingSpeed 40ms

Type "gtree template"
Sleep 1.5s
Enter
Sleep 1.5s
Hide
Type "gtree template"
Show
Sleep 1.5s
Type " | gtree output"
Sleep 1.5s
Enter
Sleep 3s
Hide
Type "clear" Enter
Show

Type "cat testdata/sample2.md"
Sleep 1.5s
Enter
Sleep 2.5s
Hide
Type "cat testdata/sample2.md"
Show
Sleep 1.5s
Type " | gtree output"
Sleep 1.5s
Enter
Sleep 3s
Hide
Type "clear" Enter
Show

Type "gtree template"
Sleep 1.5s
Enter
Sleep 1.5s
Hide
Type "gtree template"
Show
Sleep 1.5s
Type " | gtree mkdir --dry-run"
Sleep 1.5s
Enter
Sleep 3s
Hide
Type "clear" Enter
Show

Type "gtree template | gtree mkdir --dry-run -e .go -e .md -e Makefile"
Sleep 1.5s
Enter
Sleep 1.5s

Hide
Type "gtree template | gtree mkdir --dry-run -e .go -e .md -e Makefile"
Show
Sleep 1.5s
Left 26
Sleep 2s
Backspace 10
Sleep 1.5s
Enter
Sleep 1.5s

Type "tree gtree --dirsfirst"
Sleep 1.5s
Enter
Sleep 4s
Hide
Type "rm -rf gtree/" Enter
Show