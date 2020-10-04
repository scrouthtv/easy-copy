#!/bin/bash

SRCDIR=$(dirname $0)/src
BINLOC=$(dirname $0)/pkg/easycopy
PKGDIR=$(dirname $0)/pkg
OUTDIR=$(dirname $0)/build

WIN_ARCHES=("386" amd64)
LINUX_ARCHES=("386" amd64)
DARWIN_ARCHES=("386" amd64 arm)
for GOARCH in ${WIN_ARCHES[@]}; do
	echo "compiling for windows $GOARCH"
	GOOS=windows go build -o "$BINLOC-windows-$GOARCH.exe" "$SRCDIR"
	zip "$OUTDIR/windows-$GOARCH.zip" "$PKGDIR/"*
done
for GOARCH in ${LINUX_ARCHES[@]}; do
	echo "compiling for linux $GOARCH"
	GOOS=linux go build -o "$BINLOC-linux-$GOARCH" "$SRCDIR"
	zstd -o "$OUTDIR/windows-$GOARCH.zip" "$PKGDIR/"*
done
for GOARCH in ${DARWIN_ARCHES[@]}; do
	echo "compiling for darwin $GOARCH"
	GOOS=darwin go build -o "$BINLOC-darwin-$GOARCH.app" "$SRCDIR"
	zip "$OUTDIR/linux-$GOARCH.zip" "$PKGDIR/"*
done
