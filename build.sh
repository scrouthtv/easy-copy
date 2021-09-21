#!/bin/sh

VERSION="1.0.0-RC1"

mkdir out -p
rm out/* -f

echo "linux/amd64"
rm pkg/* -f
go build -v .
cp easy-copy pkg/
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
GOOS=linux GOARCH=arm go build .
cp easy-copy pkg/
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
GOOS=windows GOARCH=amd64 CGO_ENABLED=1 CC=x86_64-w64-mingw32-gcc go build .
cp easy-copy.exe pkg/
cp LICENSE pkg/
cd pkg/
zip "easycopy-$VERSION-windows-amd64.zip" *
cd ..
cp pkg/"easycopy-$VERSION-windows-amd64.zip" out/
