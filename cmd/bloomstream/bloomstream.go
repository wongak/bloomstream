package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	stdlog "log"

	"github.com/go-kit/kit/log"
)

var (
	// Version is the bloomstream version
	Version = "dev"
)

// os interfaces
var (
	stdout io.Writer
	stderr io.Writer
)

// globals
var (
	logger *log.Context
)

func init() {
	stdout = os.Stdout
	stderr = os.Stderr
}

func usage() {
	fmt.Fprintf(stderr, "USAGE\n")
	fmt.Fprintf(stderr, "\t%s <mode> [flags]\n", os.Args[0])
	fmt.Fprintf(stderr, "\n")
	fmt.Fprintf(stderr, "MODES\n")
	fmt.Fprintf(stderr, "\teditor  Runs and serves the editor\n")
	fmt.Fprintf(stderr, "\n")
	fmt.Fprintf(stderr, "VERSION\n")
	fmt.Fprintf(stderr, "\t%s\n", Version)
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	logger = log.NewContext(log.NewLogfmtLogger(log.NewSyncWriter(stdout))).WithPrefix(
		"t", log.Timestamp,
		"version", Version,
	)
	stdlog.SetOutput(log.NewStdlibAdapter(logger,
		log.TimestampKey("t"),
	))

	var run func([]string) error
	switch strings.ToLower(os.Args[1]) {
	case "editor":
		run = runEditor
	default:
		usage()
		os.Exit(1)
	}

	if err := run(os.Args[2:]); err != nil {
		fmt.Fprintf(stderr, "%v\n", err)
		os.Exit(1)
	}
}

func usageFor(fs *flag.FlagSet, short string) func() {
	return func() {
		fmt.Fprintf(os.Stderr, "USAGE\n")
		fmt.Fprintf(os.Stderr, "  %s\n", short)
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "FLAGS\n")
		fs.PrintDefaults()
		fmt.Fprintf(os.Stderr, "\n")
	}
}
