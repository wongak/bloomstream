package main

import (
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestBloomstream(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Bloomstream Suite")
	var oldArgs []string
	BeforeSuite(func() {
		oldArgs = os.Args
	})
	AfterSuite(func() {
		os.Args = oldArgs
		stderr = os.Stderr
		stdout = os.Stdout
	})
}
