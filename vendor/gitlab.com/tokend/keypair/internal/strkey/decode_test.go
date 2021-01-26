package strkey_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	. "gitlab.com/tokend/keypair/internal/strkey"
)

func TestDecode(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		t.Run("address", func(t *testing.T) {
			expected := []byte{
				0x36, 0x3e, 0xaa, 0x38, 0x67, 0x84, 0x1f, 0xba,
				0xd0, 0xf4, 0xed, 0x88, 0xc7, 0x79, 0xe4, 0xfe,
				0x66, 0xe5, 0x6a, 0x24, 0x70, 0xdc, 0x98, 0xc0,
				0xec, 0x9c, 0x07, 0x3d, 0x05, 0xc7, 0xb1, 0x03,
			}
			got, err := Decode(VersionByteAccountID, "GA3D5KRYM6CB7OWQ6TWYRR3Z4T7GNZLKERYNZGGA5SOAOPIFY6YQHES5")
			assert.NoError(t, err)
			assert.Equal(t, expected, got)
		})

		t.Run("seed", func(t *testing.T) {
			expected := []byte{
				0x69, 0xa8, 0xc4, 0xcb, 0xb9, 0xf6, 0x4e, 0x8a,
				0x07, 0x98, 0xf6, 0xe1, 0xac, 0x65, 0xd0, 0x6c,
				0x31, 0x62, 0x92, 0x90, 0x56, 0xbc, 0xf4, 0xcd,
				0xb7, 0xd3, 0x73, 0x8d, 0x18, 0x55, 0xf3, 0x63,
			}
			got, err := Decode(VersionByteSeed, "SBU2RRGLXH3E5CQHTD3ODLDF2BWDCYUSSBLLZ5GNW7JXHDIYKXZWHOKR")
			assert.NoError(t, err)
			assert.Equal(t, expected, got)
		})

		t.Run("balance", func(t *testing.T) {
			expected := []byte{
				0x21, 0x3f, 0xbc, 0x9b, 0xe7, 0x4c, 0x44, 0xd5,
				0xce, 0xb3, 0xc9, 0x65, 0x4a, 0x39, 0x89, 0x27,
				0x51, 0xa0, 0xf4, 0x68, 0x4f, 0x61, 0xd3, 0x65,
				0x86, 0x2c, 0x4c, 0x2e, 0x6e, 0x5f, 0xb9, 0x13,
			}
			got, err := Decode(VersionByteBalanceID, "BAQT7PE345GEJVOOWPEWKSRZRETVDIHUNBHWDU3FQYWEYLTOL64RH75C")
			assert.NoError(t, err)
			assert.Equal(t, expected, got)
		})
	})

	t.Run("fails", func(t *testing.T) {
		t.Run("incorrect version byte", func(t *testing.T) {
			valid := "GA3D5KRYM6CB7OWQ6TWYRR3Z4T7GNZLKERYNZGGA5SOAOPIFY6YQHES5"
			_, err := Decode(VersionByteBalanceID, valid)
			assert.Error(t, err)
		})

		t.Run("invalid version byte", func(t *testing.T) {
			valid := "GA3D5KRYM6CB7OWQ6TWYRR3Z4T7GNZLKERYNZGGA5SOAOPIFY6YQHES5"
			_, err := Decode(VersionByte(2), valid)
			assert.Error(t, err)
		})

		t.Run("corrupted checksum", func(t *testing.T) {
			valid := "GA3D5KRYM6CB7OWQ6TWYRR3Z4T7GNZLKERYNZGGA5SOAOPIFY6YQHES5"
			corrupted := "GB" + valid[2:]
			_, err := Decode(VersionByteAccountID, corrupted)
			assert.Error(t, err)
		})

		t.Run("corrupted payload", func(t *testing.T) {
			valid := "GA3D5KRYM6CB7OWQ6TWYRR3Z4T7GNZLKERYNZGGA5SOAOPIFY6YQHES5"
			corrupted := valid[0:len(valid)-2] + "Z5"
			_, err := Decode(VersionByteAccountID, corrupted)
			assert.Error(t, err)
		})

		t.Run("empty input", func(t *testing.T) {
			_, err := Decode(VersionByteBalanceID, "")
			assert.Error(t, err)
		})
	})
}
