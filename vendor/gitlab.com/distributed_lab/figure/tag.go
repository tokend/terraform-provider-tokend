package figure

import (
	"reflect"

	"strings"

	"gitlab.com/distributed_lab/logan/v3/errors"
)

var (
	ErrUnknownAttribute      = errors.New("Unknown syntax of tag")
	ErrConflictingAttributes = errors.New("Conflict attributes")
)

type Tag struct {
	Key        string
	Required   bool
	NonZero    bool
}

func parseFieldTag(field reflect.StructField, tagKey string) (*Tag, error) {
	tag := &Tag{}

	fieldTag := field.Tag.Get(tagKey)
	splitedTag := strings.Split(fieldTag, `,`)

	if len(splitedTag) == 1 && splitedTag[0] == ignore {
		return nil, nil
	}

	if len(splitedTag) == 0 {
		tag.Key = ""
	} else {
		tag.Key = splitedTag[0]
	}

	if tag.Key == "" {
		tag.Key = toSnakeCase(field.Name)
	}

	if len(splitedTag) > 1 {
		if contains(splitedTag, ignore) {
			return nil, ErrConflictingAttributes
		}

		for _, rule := range splitedTag[1:] {
			switch rule {
			case required:
				tag.Required = true
			case nonZero:
				tag.NonZero = true
			default:
				return nil, ErrUnknownAttribute
			}
		}
	}

	return tag, nil
}

func contains(slice []string, value string) bool {
	for _, item := range slice {
		if item == value {
			return true
		}
	}
	return false
}
