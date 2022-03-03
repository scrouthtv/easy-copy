package flags_test

import "easy-copy/flags"
import "testing"

func TestSimpleSep(t *testing.T) {
	l := "ec cp foo bar"
	sha := []string{"ec", "cp", "foo", "bar"}

	isa, err := flags.Sep(l)
	snf(t, err)

	compSlice(t, isa, sha)
}

func TestBackslash(t *testing.T) {
	l := "ec cp file\\ with\\ space bar"
	sha := []string{"ec", "cp", "file with space", "bar"}

	isa, err := flags.Sep(l)
	snf(t, err)

	compSlice(t, isa, sha)
}

func TestQuotes(t *testing.T) {
	l := "ec cp \"file a\" 'file b'"
	sha := []string{"ec", "cp", "file a", "file b"}

	isa, err := flags.Sep(l)
	snf(t, err)

	compSlice(t, isa, sha)
}

// snf fails if the error is not nil
func snf(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("Shouldn't have failed, but did: %s", err)
	}
}

func compSlice(t *testing.T, is, should []string) {
	t.Helper()

	t.Logf("Is: %v", is)
	t.Logf("Should: %v", should)

	if len(is) == 0 {
		t.Errorf("Slice should be empty")
		return
	} else if len(is) == 0 {
		t.Errorf("Slice is empty")
		return
	}

	if len(is) != len(should) {
		t.Errorf("Wrong slice length: got %d, expected %d", len(is), len(should))
		return
	}

	for i := range is {
		if is[i] != should[i] {
			t.Errorf("Non-matching element at pos %d: got %s, expected %s",
						i, is[i], should[i])
		}
	}
}

