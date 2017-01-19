package flow_test

import (
	. "github.com/wongak/bloomstream/pkg/flow"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("A Flow", func() {
	var f Flow

	It("has a zero identity", func() {
		Expect(f.ID).To(Equal(int64(0)))
	})
	It("can be identified", func() {
		f.ID = 123
		Expect(f.ID).To(Equal(int64(123)))
	})
})
