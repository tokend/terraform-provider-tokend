/*
 * GENERATED. Do not modify. Your changes might be overwritten!
 */

package resources

type Flagger interface {
	IsFlag() bool
}

type Flag struct {
	Name  string `json:"name"`
	Value int32  `json:"value"`
}

type Flags struct {
	Mask   int32  `json:"mask"`
	Values []Flag `json:"flags"`
}

func FlagsFromMask(mask int32, allFlags map[int32]string) Flags {
	values := []Flag{}

	for value, name := range allFlags {
		if (value & mask) == value {
			values = append(values, Flag{
				Value: value,
				Name:  name,
			})
		}
	}

	return Flags{
		Mask:   mask,
		Values: values,
	}
}
