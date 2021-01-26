package strkey_test

import (
	. "gitlab.com/tokend/go/strkey"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("strkey.Decode", func() {
	validAccount := "GA3D5KRYM6CB7OWQ6TWYRR3Z4T7GNZLKERYNZGGA5SOAOPIFY6YQHES5"
	validAccountPayload := []byte{
		0x36, 0x3e, 0xaa, 0x38, 0x67, 0x84, 0x1f, 0xba,
		0xd0, 0xf4, 0xed, 0x88, 0xc7, 0x79, 0xe4, 0xfe,
		0x66, 0xe5, 0x6a, 0x24, 0x70, 0xdc, 0x98, 0xc0,
		0xec, 0x9c, 0x07, 0x3d, 0x05, 0xc7, 0xb1, 0x03,
	}
	validSeed := "SBU2RRGLXH3E5CQHTD3ODLDF2BWDCYUSSBLLZ5GNW7JXHDIYKXZWHOKR"
	validSeedPayload := []byte{
		0x69, 0xa8, 0xc4, 0xcb, 0xb9, 0xf6, 0x4e, 0x8a,
		0x07, 0x98, 0xf6, 0xe1, 0xac, 0x65, 0xd0, 0x6c,
		0x31, 0x62, 0x92, 0x90, 0x56, 0xbc, 0xf4, 0xcd,
		0xb7, 0xd3, 0x73, 0x8d, 0x18, 0x55, 0xf3, 0x63,
	}

	It("decodes valid values", func() {
		payload, err := Decode(VersionByteAccountID, validAccount)
		Expect(err).To(BeNil())
		Expect(payload).To(Equal(validAccountPayload))

		payload, err = Decode(VersionByteSeed, validSeed)
		Expect(err).To(BeNil())
		Expect(payload).To(Equal(validSeedPayload))
	})

	Context("the expected version byte doesn't match the actual version byte", func() {
		It("fails", func() {
			_, err := Decode(VersionByteAccountID, validSeed)
			Expect(err).To(HaveOccurred())
			_, err = Decode(VersionByteSeed, validAccount)
			Expect(err).To(HaveOccurred())
		})
	})

	Context("the expected version byte isn't a valid constant", func() {
		It("fails", func() {
			_, err := Decode(VersionByte(2), validAccount)
			Expect(err).To(HaveOccurred())
		})
	})

	Context("the checksum has been corrupted", func() {
		It("fails", func() {
			corrupted := "GB" + validAccount[2:]
			_, err := Decode(VersionByteAccountID, corrupted)
			Expect(err).To(HaveOccurred())
		})
	})

	Context("the payload has been corrupted", func() {
		It("fails", func() {
			corrupted := validAccount[0:len(validAccount)-2] + "Z5"
			_, err := Decode(VersionByteAccountID, corrupted)
			Expect(err).To(HaveOccurred())
		})
	})

	Context("the input is empty", func() {
		It("fails", func() {
			_, err := Decode(VersionByteAccountID, "")
			Expect(err).To(HaveOccurred())
		})
	})

})
