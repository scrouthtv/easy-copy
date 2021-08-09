package fs

import (
	"errors"
	"io"
	"testing"
)

func TestRead(t *testing.T) {
	f := NewFile("test")
	f.contents = []byte("foobar")

	buf := make([]byte, 3)

	n, err := f.Read(buf)
	if err != nil {
		t.Error(err)
	}

	if n != 3 {
		t.Error("Read wrong number of bytes:", n, "expected: 3")
	}

	if string(buf) != "foo" {
		t.Error("Read wrong contents:", string(buf), "expected: foo")
	}

	_, err = f.ReadAt(buf, 1)
	if err != nil {
		t.Error(err)
	}

	if string(buf) != "oob" {
		t.Error("ReadAt wrong contents:", string(buf), "expected: oob")
	}

	buf = make([]byte, 8)

	_, err = f.Read(buf)
	if err != nil {
		t.Error(err)
	}

	if string(buf[:3]) != "bar" {
		t.Error("Read wrong contents:", string(buf), "expected: bar")
	}

	n, err = f.Read(buf)
	if !errors.Is(err, io.EOF) {
		t.Error("Read wrong error:", err, "expected: io.EOF")
	}

	if n != 0 {
		t.Error("Read wrong number of bytes:", n, "expected: 0")
	}
}

func TestSeek(t *testing.T) {
	f := NewFile("test")
	f.contents = []byte("foobarbaz")

	buf := make([]byte, 3)

	_, err := f.Read(buf)
	if err != nil {
		t.Error(err)
	}

	if string(buf) != "foo" {
		t.Error("Read wrong contents:", string(buf), "expected: foo")
	}

	n, err := f.Seek(1, io.SeekStart)
	if n != 1 {
		t.Error("Seek wrong number of bytes:", n, "expected: 1")
	}

	if err != nil {
		t.Error(err)
	}

	_, err = f.Read(buf)
	if err != nil {
		t.Error(err)
	}

	if string(buf) != "oob" {
		t.Error("Read wrong contents:", string(buf), "expected: oob")
	}

	n, err = f.Seek(2, io.SeekCurrent)
	if n != 6 {
		t.Error("Seek wrong number of bytes:", n, "expected: 6")
	}

	if err != nil {
		t.Error(err)
	}

	_, err = f.Read(buf)
	if err != nil {
		t.Error(err)
	}

	if string(buf) != "baz" {
		t.Error("Read wrong contents:", string(buf), "expected: baz")
	}

	n, err = f.Seek(-2, io.SeekEnd)
	if n != 7 {
		t.Error("Seek wrong number of bytes:", n, "expected: 7")
	}

	if err != nil {
		t.Error(err)
	}

	read, err := f.Read(buf)
	if err != nil {
		t.Error(err)
	}

	if string(buf[:read]) != "az" {
		t.Error("Read wrong contents:", string(buf[:read]), "expected: az")
	}

	_, err = f.Seek(0, io.SeekStart+io.SeekCurrent+io.SeekEnd)

	var errOpNotSupported error = &ErrOperationNotSupported{}
	if !errors.As(err, &errOpNotSupported) {
		t.Error("Wrong error:", err)
	}
}

func TestWrite(t *testing.T) {
	f := NewFile("test")

	n, err := f.Write([]byte("qwertz"))
	if n != 6 {
		t.Error("Write wrong number of bytes:", n, "expected: 6")
	}

	if err != nil {
		t.Error(err)
	}

	if string(f.contents) != "qwertz" {
		t.Error("Write wrong contents:", string(f.contents), "expected: qwertz")
	}

	_, err = f.WriteAt([]byte("bar"), 1)
	if err != nil {
		t.Error(err)
	}

	if string(f.contents) != "qbartz" {
		t.Error("WriteAt wrong contents:", string(f.contents), "expected: qbartz")
	}

	_, err = f.Write([]byte("hehe"))
	if err != nil {
		t.Error(err)
	}

	if string(f.contents) != "qbartzhehe" {
		t.Error("Write wrong contents:", string(f.contents), "expected: qbartzhehe")
	}

	seek, err := f.Seek(-3, io.SeekCurrent)
	if seek != 7 {
		t.Error("Seek wrong number of bytes:", seek, "expected: 7")
	}

	if err != nil {
		t.Error(err)
	}

	_, err = f.Write([]byte("foo"))
	if err != nil {
		t.Error(err)
	}

	if string(f.contents) != "qbartzhfoo" {
		t.Error("Write wrong contents:", string(f.contents), "expected: qbartzhfoo")
	}

	err = f.Truncate(3)
	if err != nil {
		t.Error(err)
	}

	if string(f.contents) != "qba" {
		t.Error("Truncate wrong contents:", string(f.contents), "expected: qba")
	}
}

func TestWriteString(t *testing.T) {
	f := NewFile("test")

	n, err := f.WriteString("qwertz")
	if n != 6 {
		t.Error("Write wrong number of bytes:", n, "expected: 6")
	}

	if err != nil {
		t.Error(err)
	}

	if string(f.contents) != "qwertz" {
		t.Error("Write wrong contents:", string(f.contents), "expected: qwertz")
	}
}

func TestFolderRW(t *testing.T) {
	folder := NewFolder("test")
	var errNotAFile error = &ErrNotAFile{}

	n, err := folder.Write([]byte("qwertz"))
	if n != 0 {
		t.Error("Write wrong number of bytes:", n, "expected: 0")
	}

	if !errors.As(err, &errNotAFile) {
		t.Error("Wrong error:", err)
	}

	if err.Error() != "not a file: test" {
		t.Error("Wrong error message:", err, "expected not a file: test")
	}

	n, err = folder.Read(make([]byte, 8))
	if n != 0 {
		t.Error("Read wrong number of bytes:", n, "expected: 0")
	}

	if !errors.As(err, &errNotAFile) {
		t.Error("Wrong error:", err)
	}
}
