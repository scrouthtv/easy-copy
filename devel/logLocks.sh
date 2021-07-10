#!/bin/bash

for f in $(ls *.go); do
	printf "%s\n" $f
	perl -p \
		-e 's/^(([[:space:]\t]*)lock\.((Lock)|(RLock))\(\);)$/\2fmt.Println("[$ARGV:$.] lock");\n\1/g' -i $f
	perl -p \
		-e 's/^(([[:space:]\t]*)lock\.((Unlock)|(RUnlock))\(\);)$/\2fmt.Println("[$ARGV:$.] unlock");\n\1/g' -i $f
done

# sed -E 's/^([[:space:]\t]*)filesLock\.((RLock)|(Lock))\(\);$/\1/g'
