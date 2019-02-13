package flagalias

import (
	"flag"
	"strings"
	"testing"
)

func TestAlias(t *testing.T) {
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	foo := fs.Bool("foo", false, "")
	if err := AliasOnFlagSet(fs, "bar", "foo"); err != nil {
		t.Fatal(err)
	}

	fs.Parse([]string{"--bar"})

	if got, want := *foo, true; got != want {
		t.Fatalf("got: %v, want: %v", got, want)
	}
}

func TestDeprecated(t *testing.T) {
	var buf strings.Builder
	fs := flag.NewFlagSet("test", flag.ContinueOnError)
	fs.SetOutput(&buf)

	foo := fs.Bool("foo", false, "")
	if err := DeprecatedOnFlagSet(fs, "bar", "foo"); err != nil {
		t.Fatal(err)
	}

	fs.Parse([]string{"--bar"})

	if got, want := *foo, true; got != want {
		t.Fatalf("got: %v, want: %v", got, want)
	}

	if got, want := buf.String(), "Flag \"bar\" is deprecated, please use \"foo\" instead\n"; got != want {
		t.Fatalf("got: %q, want: %q", got, want)
	}
}
