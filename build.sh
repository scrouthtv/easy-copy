#!/bin/bash

ROOTDIR="$(realpath $(dirname $0))"
echo "root is $ROOTDIR"
SRCDIR="$ROOTDIR/src"
BINLOC="$ROOTDIR/pkg/easycopy"
PKGDIR="$ROOTDIR/pkg"
OUTDIR="$ROOTDIR/build"
DOCDIR="$ROOTDIR/doc"
SHRDIR="$ROOTDIR/shared"
LICENSE="$ROOTDIR/LICENSE"
VERSION="$(grep 'EASYCOPY_VERSION' $SRCDIR/meta.go | grep -oE '[0-9]+.[0-9]+.[0-9]+')"
OUTPREFIX="easycopy-$VERSION"

# using https://github.com/tpoechtrager/osxcross
DARWINCC="$ROOTDIR/osxcross/bin"
CC_FOR_darwin_amd64="$DARWINCC/o64-clang"
CC_FOR_darwin_arm="$DARWINCC/oa64-clang"

mkdir -p "$ROOTDIR" "$SRCDIR" "$PKGDIR" "$OUTDIR"
rm -rf "$OUTDIR/"*

WIN_ARCHES=("386" amd64)
LINUX_ARCHES=("386" amd64)
#DARWIN_ARCHES=("386" amd64 arm arm64)

cp "$LICENSE" "$PKGDIR/"
cd "$SRCDIR"
export CGO_ENABLED=1
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
	cp "$DOCDIR/INSTALL" "$PKGDIR/"
	cp "$SHRDIR/_ec" "$SHRDIR/ec-completion.bash" "$SHRDIR/ec-completion.fish" "$PKGDIR/"
	zstd -v -o "$OUTDIR/$OUTPREFIX-linux-$GOARCH.zst" "$PKGDIR/"* 2>&1 | sed 's/(.*)//g'
	rm "$PKGDIR/easycopy.1" "$PKGDIR/ec.conf.5" "$BINLOC-linux-$GOARCH"
done
for GOARCH in ${DARWIN_ARCHES[@]}; do
	echo "compiling for darwin $GOARCH"
	GOOS=darwin go build -o "$BINLOC-darwin-$GOARCH.app" .
	zip -j "$OUTDIR/$OUTPREFIX-darwin-$GOARCH.zip" "$PKGDIR/"*
	rm "$BINLOC-darwin-$GOARCH.app"
done
rm "$PKGDIR/"*

cd "$ROOTDIR"
