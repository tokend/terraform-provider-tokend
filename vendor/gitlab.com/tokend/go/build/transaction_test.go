package build

import (
	"gitlab.com/tokend/go/xdr"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Transaction Mutators:", func() {

	var (
		subject *TransactionBuilder
		mut     TransactionMutator
	)

	BeforeEach(func() { subject = &TransactionBuilder{} })
	JustBeforeEach(func() { subject.Mutate(mut) })

	Describe("MemoHash", func() {
		BeforeEach(func() { mut = MemoHash{[32]byte{0x01}} })
		It("sets a Hash memo on the transaction", func() {
			Expect(subject.TX.Memo.Type).To(Equal(xdr.MemoTypeMemoHash))
			Expect(subject.TX.Memo.MustHash()).To(Equal(xdr.Hash([32]byte{0x01})))
		})
	})

	Describe("MemoID", func() {
		BeforeEach(func() { mut = MemoID{123} })
		It("sets an ID memo on the transaction", func() {
			Expect(subject.TX.Memo.Type).To(Equal(xdr.MemoTypeMemoId))
			Expect(subject.TX.Memo.MustId()).To(Equal(xdr.Uint64(123)))
		})
	})

	Describe("MemoReturn", func() {
		BeforeEach(func() { mut = MemoReturn{[32]byte{0x01}} })
		It("sets a Hash memo on the transaction", func() {
			Expect(subject.TX.Memo.Type).To(Equal(xdr.MemoTypeMemoReturn))
			Expect(subject.TX.Memo.MustRetHash()).To(Equal(xdr.Hash([32]byte{0x01})))
		})
	})

	Describe("MemoText", func() {
		BeforeEach(func() { mut = MemoText{"hello"} })
		It("sets a TEXT memo on the transaction", func() {
			Expect(subject.TX.Memo.Type).To(Equal(xdr.MemoTypeMemoText))
			Expect(subject.TX.Memo.MustText()).To(Equal("hello"))
		})

		Context("a string longer than 28 bytes", func() {
			BeforeEach(func() { mut = MemoText{"12345678901234567890123456789"} })
			It("sets an error", func() {
				Expect(subject.Err).ToNot(BeNil())
			})
		})
	})

	Describe("SourceAccount", func() {
		Context("with a valid address", func() {
			address := "GAXEMCEXBERNSRXOEKD4JAIKVECIXQCENHEBRVSPX2TTYZPMNEDSQCNQ"
			BeforeEach(func() { mut = SourceAccount{address} })
			It("sets the AccountId correctly", func() {
				var aid xdr.AccountId
				aid.SetAddress(address)
				Expect(subject.TX.SourceAccount.MustEd25519()).To(Equal(aid.MustEd25519()))
			})
		})

		Context("with bad address", func() {
			BeforeEach(func() { mut = SourceAccount{"foo"} })
			It("fails", func() { Expect(subject.Err).To(HaveOccurred()) })
		})
	})

})
