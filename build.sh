#!/bin/bash

ROOTDIR="."
echo "root is $ROOTDIR"
SRCDIR="./"
BINLOC="pkg/easycopy"
PKGDIR="pkg/"
OUTDIR="build"
DOCDIR="doc"
SHRDIR="shared"
LICENSE="LICENSE"
VERSION="$(grep 'EASYCOPY_VERSION' $SRCDIR/meta.go | grep -oE '[0-9]+.[0-9]+.[0-9]+')"
OUTPREFIX="easycopy-$VERSION"

# using https://github.com/tpoechtrager/osxcross
CC_FOR_darwin=( ["386"]="/usr/bin/gcc" ["amd64"]="/usr/bin/gcc" )

CC_FOR_windows=( ["amd64"]="/usr/bin/x86_64-w64-mingw32-gcc" ["386"]="/usr/bin/i686-w64-mingw32-gcc" )

CGO_ENABLED=1

mkdir -p "$ROOTDIR" "$SRCDIR" "$PKGDIR" "$OUTDIR"
rm -rf "$OUTDIR/"*

WIN_ARCHES=( "386" amd64 )
LINUX_ARCHES=( "386" amd64 )
FREEBSD_ARCHES=( "386" amd64 arm )
OPENBSD_ARCHES=( "386" amd64 arm )
NETBSD_ARCHES=( "386" amd64 arm )
DARWIN_ARCHES=( amd64 )

cp "$LICENSE"* "$PKGDIR/"
cd "$SRCDIR"

GOOS=windows 
for GOARCH in ${WIN_ARCHES[@]}; do
	echo ""
	echo "compiling for $GOOS $GOARCH:"
	CC=${CC_FOR_windows[$GOARCH]} go build -o "$BINLOC-windows-$GOARCH.exe" .
	strip "$BINLOC-windows-$GOARCH.exe"

	echo "packaging for windows $GOARCH:"
	ls "$PKGDIR/"
	# -j: relative paths -q: quiet
	zip -jq "$OUTDIR/$OUTPREFIX-windows-$GOARCH.zip" "$PKGDIR/"*
	rm "$BINLOC-windows-$GOARCH.exe"
done

GOOS=linux
for GOARCH in ${LINUX_ARCHES[@]}; do
	echo ""
	echo "compiling for linux $GOARCH:"
	go build -o "$BINLOC-linux-$GOARCH" .
	strip "$BINLOC-linux-$GOARCH"

	echo "packaging for linux $GOARCH:"
	cp "$DOCDIR/easycopy.1" "$DOCDIR/ec.conf.5" "$PKGDIR/"
	cp "$DOCDIR/INSTALL" "$PKGDIR/"
	cp "$SHRDIR/_ec" "$SHRDIR/ec-completion.bash" "$SHRDIR/ec-completion.fish" "$PKGDIR/"
	ls "$PKGDIR/"
	tar --zstd -cf "$OUTDIR/$OUTPREFIX-linux-$GOARCH.tar.zst" -C "pkg/" .
	rm "$PKGDIR/easycopy.1" "$PKGDIR/ec.conf.5" "$BINLOC-linux-$GOARCH"
done

GOOS=darwin 
for GOARCH in ${DARWIN_ARCHES[@]}; do
	echo ""
	echo "compiling for darwin $GOARCH:"
	CC=${CC_FOR_darwin[$GOARCH]} go build -o "$BINLOC-darwin-$GOARCH.app" .
	strip "$BINLOC-darwin-$GOARCH.app"

	echo "packaging for darwin $GOARCH:"
	ls "$PKGDIR/"
	zip -jq "$OUTDIR/$OUTPREFIX-darwin-$GOARCH.zip" "$PKGDIR/"*
	rm "$BINLOC-darwin-$GOARCH.app"
done
rm "$PKGDIR/"*

cd "$ROOTDIR"
