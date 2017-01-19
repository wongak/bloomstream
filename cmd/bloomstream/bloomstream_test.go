package main

import (
	"bytes"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type testExit struct{}

var _ = Describe("The bloomstream command", func() {
	var args = []string{
		"bloomstream",
	}
	var stdErr *bytes.Buffer
	var stdOut *bytes.Buffer

	var exitCode int

	BeforeEach(func() {
		stdErr = bytes.NewBuffer(make([]byte, 0, 128))
		stdOut = bytes.NewBuffer(make([]byte, 0, 128))
		stderr = stdErr
		stdout = stdOut
	})
	JustBeforeEach(func() {
		os.Args = args
		exit = func(i int) {
			exitCode = i
			panic(testExit{})
		}
		defer func() {
			if p := recover(); p != nil {
				if _, ok := p.(testExit); !ok {
					panic(p)
				}
			}
			exit = os.Exit
		}()
		main()
	})
	Context("when calling it without a mode", func() {
		It("should display the usage on stderr", func() {
			Expect(stdErr.String()).To(ContainSubstring("USAGE"))
		})
		It("should exit with code 1", func() {
			Expect(exitCode).To(Equal(1))
		})
	})
	Context("when calling it with an unknown mode", func() {
		BeforeEach(func() {
			args = []string{"a", "asdf"}
		})
		It("should display the usage on stderr", func() {
			Expect(stdErr.String()).To(ContainSubstring("USAGE"))
		})
		It("should exit with code 1", func() {
			Expect(exitCode).To(Equal(1))
		})
	})

	Context("when running the editor mode", func() {
		BeforeEach(func() {
			args = append(args, "editor")
		})
	})
})
