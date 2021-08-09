package fs

import "testing"

func TestEqualFS(t *testing.T) {
	a := CreateFS([]string{
		"foo/empty/",
		"bar/a",
		"bar/b",
		"bar/c",
		"foo/x",
		"test.txt",
	})
	b := CreateFS([]string{ // same as a
		"foo/empty/",
		"bar/a",
		"bar/b",
		"bar/c",
		"foo/x",
		"test.txt",
	})
	c := CreateFS([]string{ // missing /foo/empty/
		"bar/a",
		"bar/b",
		"bar/c",
		"foo/x",
		"test.txt",
	})
	d := CreateFS([]string{ // missing /bar/b
		"foo/empty/",
		"bar/a",
		"bar/c",
		"foo/x",
		"test.txt",
	})
	e := CreateFS([]string{ // missing /test.txt
		"foo/empty/",
		"bar/a",
		"bar/b",
		"bar/c",
		"foo/x",
	})
	f := CreateFS([]string{ // extra /bar/quz/
		"foo/empty/",
		"bar/quz/",
		"bar/a",
		"bar/b",
		"bar/c",
		"foo/x",
		"test.txt",
	})
	g := CreateFS([]string{ // renamed bar/c to bar/d
		"foo/empty/",
		"bar/a",
		"bar/b",
		"bar/d",
		"foo/x",
		"test.txt",
	})
	h := CreateFS([]string{ // renamed foo/empty to foo/empti
		"foo/empti/",
		"bar/a",
		"bar/b",
		"bar/c",
		"foo/x",
		"test.txt",
	})
	i := CreateFS([]string{ // all of the above
		"bar/quz/",
		"bar/a",
		"bar/c",
		"foo/x",
	})

	ok, bad := a.Equal(b)
	if !ok {
		t.Error("a != b, failed at", bad)
	}

	ok, bad = a.Equal(c)
	if ok {
		t.Error("a == c")
	}
	if bad != "/foo" {
		t.Error("wrong bad position:", bad, "should be /foo")
	}

	ok, bad = a.Equal(d)
	if ok {
		t.Error("a == d")
	}
	if bad != "/bar" {
		t.Error("wrong bad position:", bad, "should be /bar")
	}

	ok, bad = a.Equal(e)
	if ok {
		t.Error("a == e")
	}
	if bad != "/" {
		t.Error("wrong bad position:", bad, "should be /")
	}

	ok, bad = a.Equal(f)
	if ok {
		t.Error("a == f")
	}
	if bad != "/bar" {
		t.Error("wrong bad position:", bad, "should be /bar")
	}

	ok, bad = a.Equal(g)
	if ok {
		t.Error("a == g")
	}
	if bad != "/bar/c" {
		t.Error("wrong bad position:", bad, "should be /bar/c")
	}

	ok, bad = a.Equal(h)
	if ok {
		t.Error("a == h")
	}
	if bad != "/foo/empty" {
		t.Error("wrong bad position:", bad, "should be /foo/empty")
	}

	ok, _ = a.Equal(i)
	if ok {
		t.Error("a == i")
	}

}
