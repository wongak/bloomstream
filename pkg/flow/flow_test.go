package flow_test

import (
	. "github.com/wongak/bloomstream/pkg/flow"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("A Flow", func() {
	var f Flow

	It("has a zero ID", func() {
		Expect(f.ID).To(Equal(int64(0)))
	})
	It("has a modifieable ID", func() {
		f.ID = 123
		Expect(f.ID).To(Equal(int64(123)))
	})
	It("has a modifieable title", func() {
		f.Title = "Hello, world!"
		Expect(f.Title).To(Equal("Hello, world!"))
	})
})
