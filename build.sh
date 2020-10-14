#!/bin/bash

ROOTDIR="$(realpath $(dirname $0))"
echo "root is $ROOTDIR"
SRCDIR="$ROOTDIR/src"
BINLOC="$ROOTDIR/pkg/easycopy"
PKGDIR="$ROOTDIR/pkg"
OUTDIR="$ROOTDIR/build"
DOCDIR="$ROOTDIR/doc"
LICENSE="$ROOTDIR/LICENSE"
VERSION="$(grep 'EASYCOPY_VERSION' $SRCDIR/meta.go | grep -oE '[0-9]+.[0-9]+.[0-9]+')"
OUTPREFIX="easycopy-$VERSION"

mkdir -p "$ROOTDIR" "$SRCDIR" "$PKGDIR" "$OUTDIR"
rm -rf "$OUTDIR/"*

WIN_ARCHES=("386" amd64)
LINUX_ARCHES=("386" amd64)
#DARWIN_ARCHES=("386" amd64 arm arm64)

cp "$LICENSE" "$PKGDIR/"
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
	cp "$DOCDIR/easycopy.1" "$DOCDIR/ec.conf.5" "$PKGDIR/"
	zstd -v -o "$OUTDIR/$OUTPREFIX-linux-$GOARCH.zst" "$PKGDIR/"* 2>&1 | sed 's/(.*)//g'
	rm "$PKGDIR/easycopy.1" "$PKGDIR/ec.conf.5" "$BINLOC-linux-$GOARCH"
done
for GOARCH in ${DARWIN_ARCHES[@]}; do
	echo "compiling for darwin $GOARCH"
	GOOS=darwin go build -o "$BINLOC-darwin-$GOARCH.app" .
	zip -j "$OUTDIR/$OUTPREFIX-darwin-$GOARCH.zip" "$PKGDIR/"*
	rm "$BINLOC-darwin-$GOARCH.app"
done
rm "$PKGDIR/LICENSE"

cd "$ROOTDIR"
