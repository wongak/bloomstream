package main

import (
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
	exit   func(int)
)

// globals
var (
	logger *log.Context
)

func init() {
	stdout = os.Stdout
	stderr = os.Stderr
	exit = os.Exit
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
		exit(1)
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
		exit(1)
	}

	if err := run(os.Args[2:]); err != nil {
		fmt.Fprintf(stderr, "%v\n", err)
		exit(1)
	}
	exit(0)
}
