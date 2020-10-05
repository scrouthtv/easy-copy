#!/bin/bash

for f in $(ls *.go); do
	perl -p \
		-e 's/^(([[:space:]\t]*)filesLock\.((Lock)|(RLock))\(\);)$/\2fmt.Println("[$ARGV:$.] lock");\n\1/g' -i $f
	perl -p \
		-e 's/^(([[:space:]\t]*)filesLock\.((Unlock)|(RUnlock))\(\);)$/\2fmt.Println("[$ARGV:$.] unlock");\n\1/g' -i $f
done

# sed -E 's/^([[:space:]\t]*)filesLock\.((RLock)|(Lock))\(\);$/\1/g'
