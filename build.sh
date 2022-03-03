#!/bin/sh

VERSION="1.0.0-RC3"
BINNAME="easycopy"

# Version check
truever="$(go run . --version)"
if [[ $truever != "EasyCopy v"$VERSION ]]; then
	echo "Version mismatch, should be \"EasyCopy v$VERSION\", but is \"$truever\""
	echo "Aborting"
	exit
fi

mkdir out pkg -p
rm out/* -f

echo "linux/amd64"
rm pkg/* -f
go build -v -o pkg/$BINNAME .
cp shared/* pkg/
cp doc/* pkg/
cp LICENSE pkg/
cd pkg/
tar cf "easycopy-$VERSION-linux-amd64.tar" *
zstd "easycopy-$VERSION-linux-amd64.tar"
cd ..
cp pkg/"easycopy-$VERSION-linux-amd64.tar.zst" out/


echo ""
echo "linux/arm"
rm pkg/* -f
GOOS=linux GOARCH=arm go build -v -o pkg/$BINNAME .
cp shared/* pkg/
cp doc/* pkg/
cp LICENSE pkg/
cd pkg/
tar cf "easycopy-$VERSION-linux-arm.tar" *
zstd "easycopy-$VERSION-linux-arm.tar"
cd ..
cp pkg/"easycopy-$VERSION-linux-arm.tar.zst" out/


echo ""
echo "windows/amd64"
rm pkg/* -f
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build -v -o pkg/$BINNAME.exe .
cp LICENSE pkg/
cd pkg/
zip "easycopy-$VERSION-windows-amd64.zip" *
cd ..
cp pkg/"easycopy-$VERSION-windows-amd64.zip" out/

rm pkg/* 
