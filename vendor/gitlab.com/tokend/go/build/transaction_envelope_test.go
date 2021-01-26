package build

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("TransactionEnvelope Mutators:", func() {

	var (
		subject TransactionEnvelopeBuilder
		mut     TransactionEnvelopeMutator
	)

	BeforeEach(func() { subject = TransactionEnvelopeBuilder{} })
	JustBeforeEach(func() { subject.Mutate(mut) })

	Describe("TransactionBuilder", func() {

		Context("with an error set on it", func() {
			err := errors.New("busted!")
			BeforeEach(func() { mut = &TransactionBuilder{Err: err} })
			It("propagates the error upwards", func() {
				Expect(subject.Err).To(HaveOccurred())
				Expect(subject.Err.Error()).To(ContainSubstring("busted!"))
			})
		})

	})

	Describe("Sign", func() {
		Context("with a valid key", func() {
			BeforeEach(func() {
				subject.MutateTX(SourceAccount{"SDOTALIMPAM2IV65IOZA7KZL7XWZI5BODFXTRVLIHLQZQCKK57PH5F3H"}, TestNetwork)
				mut = Sign{"SDOTALIMPAM2IV65IOZA7KZL7XWZI5BODFXTRVLIHLQZQCKK57PH5F3H"}
			})

			It("succeeds", func() { Expect(subject.Err).NotTo(HaveOccurred()) })
			It("adds a signature to the envelope", func() {
				Expect(subject.E.Signatures).To(HaveLen(1))
			})
		})

		Context("with an invalid key", func() {
			BeforeEach(func() { mut = Sign{""} })

			It("fails", func() {
				Expect(subject.Err).To(HaveOccurred())
			})
		})
	})

})
