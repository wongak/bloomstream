package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/pkg/errors"
)

func runEditor(args []string) error {
	flagset := flag.NewFlagSet("editor", flag.ExitOnError)
	var (
		address = flagset.String("a", ":8080", "listen address for editor")
	)
	flagset.Usage = usageFor(flagset, fmt.Sprintf("%s editor [flags]", os.Args[0]))
	if err := flagset.Parse(args); err != nil {
		return errors.Wrap(err, "flag error")
	}
	_ = address
	return nil
}
