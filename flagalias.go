// Package flagalias allows you to rename flags and deprecated the old name.
package flagalias

import (
	"flag"
	"fmt"
)

// Alias defines an alternate name for a flag.
func Alias(old, new string) {
	if err := AliasOnFlagSet(flag.CommandLine, old, new); err != nil {
		panic(err)
	}
}

// Deprecate defines an alias which will emit a deprecation warning if used.
// The warning will be written to the output defined by flag.Commandline.Output().
func Deprecated(old, new string) {
	if err := DeprecatedOnFlagSet(flag.CommandLine, old, new); err != nil {
		panic(err)
	}
}

// AliasOnFlagSet defines an alternate name for a flag.
func AliasOnFlagSet(fs *flag.FlagSet, old, new string) error {
	return registerAlias(fs, old, new, "")
}

// DeprecatedOnFlagSet defines an alias which will emit a deprecation warning to fs.Output() if used.
func DeprecatedOnFlagSet(fs *flag.FlagSet, old, new string) error {
	return registerAlias(fs, old, new, fmt.Sprintf("Flag %q is deprecated, please use %q instead", old, new))
}

func registerAlias(fs *flag.FlagSet, old, new, msg string) error {
	of := fs.Lookup(new)
	if of == nil {
		return fmt.Errorf("cannot find flag -%s", new)
	}

	a := alias{of.Value, func() { fmt.Fprintln(fs.Output(), msg) }}
	var v flag.Value = a
	if b, ok := of.Value.(boolFlag); ok {
		v = boolAlias{a, b}
	}

	fs.Var(v, old, of.Usage)
	return nil
}

type alias struct {
	flag.Value
	report func()
}

func (a alias) Set(s string) error {
	a.report()
	return a.Value.Set(s)
}

type boolAlias struct {
	alias
	boolFlag
}

func (b boolAlias) Set(s string) error {
	return b.alias.Set(s)
}

type boolFlag interface {
	flag.Value
	IsBoolFlag() bool
}
