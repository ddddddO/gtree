#!/bin/bash

cat > mk.sh <<EOF
#!/bin/bash

cat > a.txt<<EOFF
aaaa
$HOME
gggg
EOFF

cat > b.txt<<'EOFF'
bbbb
$HOME
gggg
EOFF
EOF
