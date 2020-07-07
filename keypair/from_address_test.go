package keypair

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("keypair.FromAddress", func() {
	var subject KP

	JustBeforeEach(func() {
		subject = &FromAddress{address}
	})

	ItBehavesLikeAKP(&subject)

	Describe("Sign()", func() {
		It("fails", func() {
			_, err := subject.Sign(mdfcge)
			Expect(err).To(HaveOccurred())
		})

	})
	Describe("SignDecorated()", func() {
		It("fails", func() {
			_, err := subject.SignDecorated(mdfcge)
			Expect(err).To(HaveOccurred())
		})
	})
})
