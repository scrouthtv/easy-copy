#!/bin/bash

ROOTDIR="$(realpath $(dirname $0))"
echo "root is $ROOTDIR"
SRCDIR="$ROOTDIR/src"
BINLOC="$ROOTDIR/pkg/easycopy"
PKGDIR="$ROOTDIR/pkg"
OUTDIR="$ROOTDIR/build"
VERSION="$(grep 'EASYCOPY_VERSION' $SRCDIR/meta.go | grep -oE '[0-9]+.[0-9]+.[0-9]+')"
OUTPREFIX="easycopy-$VERSION"

mkdir -p "$ROOTDIR" "$SRCDIR" "$PKGDIR" "$OUTDIR"
rm -rf "$OUTDIR/"*

WIN_ARCHES=("386" amd64)
LINUX_ARCHES=("386" amd64)
DARWIN_ARCHES=("386" amd64 arm arm64)

cd "$SRCDIR"
for GOARCH in ${WIN_ARCHES[@]}; do
	echo "compiling for windows $GOARCH to $BINLOC-windows-$GOARCH.exe"
	GOOS=windows go build -o "$BINLOC-windows-$GOARCH.exe" .
	zip -j "$OUTDIR/$OUTPREFIX-windows-$GOARCH.zip" "$PKGDIR/"*
	rm "$BINLOC-windows-$GOARCH.exe"
done
for GOARCH in ${LINUX_ARCHES[@]}; do
	echo "compiling for linux $GOARCH"
	GOOS=linux go build -o "$BINLOC-linux-$GOARCH" .
	zstd -v -o "$OUTDIR/$OUTPREFIX-linux-$GOARCH.zstd" "$PKGDIR/"*
	rm "$BINLOC-linux-$GOARCH"
done
for GOARCH in ${DARWIN_ARCHES[@]}; do
	echo "compiling for darwin $GOARCH"
	GOOS=darwin go build -o "$BINLOC-darwin-$GOARCH.app" .
	zip -j "$OUTDIR/$OUTPREFIX-darwin-$GOARCH.zip" "$PKGDIR/"*
	rm "$BINLOC-darwin-$GOARCH.app"
done

cd "$ROOTDIR"
