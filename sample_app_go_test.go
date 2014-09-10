package main

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SampleAppGoTest", func() {
	Context("top()", func() {
		It("return 200 Status", func() {
			RequestToRoot("GET", top)
			Expect(recorder.Code).To(Equal(200))
			Expect(recorder.Body).To(ContainSubstring("Hello World!"))
		})
	})
})
