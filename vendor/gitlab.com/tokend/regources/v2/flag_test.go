package regources

import (
	"gotest.tools/assert"
	"testing"
)

func TestFlagFromMask(t *testing.T) {
	t.Run("", func(t *testing.T) {
		var allFlags = map[int]string{
			1: "some-value",
			2: "some-another-value",
			4: "and-one-more-different-value",
		}

		assert.DeepEqual(t,
			FlagsFromMask(1, allFlags),
			Flags{
				Mask: 1,
				Values: []Flag{
					{
						Value: 1,
						Name:  "some-value",
					},
				},
			})

		assert.DeepEqual(t,
			FlagsFromMask(3, allFlags),
			Flags{
				Mask: 3,
				Values: []Flag{
					{
						Value: 1,
						Name:  "some-value",
					},
					{
						Value: 2,
						Name:  "some-another-value",
					},
				},
			},
		)
	})
}
