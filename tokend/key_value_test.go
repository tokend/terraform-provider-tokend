package tokend

import (
	"testing"

	"github.com/tokend/terraform-provider-tokend/tokend/data/mocks"

	"gitlab.com/tokend/go/xdr"
	"gitlab.com/tokend/regources/generated"

	mocks2 "github.com/tokend/terraform-provider-tokend/tokend/connector/keyvalues/mocks"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/stretchr/testify/assert"
)

func TestKeyValue_Read(t *testing.T) {
	connector := mocks.Connector{}
	keyValues := mocks2.KeyValues{}
	m := Meta{Horizon: &connector}

	t.Run("new type", func(t *testing.T) {
		c := map[string]interface{}{
			"key":        "asset_update_tasks",
			"value_type": "uint32",
			"value":      "1",
		}
		newValue := "new_value"
		d := schema.TestResourceDataRaw(t, keyValueSchema, c)
		connector.On("KeyValues").Return(&keyValues).Once()
		keyValues.On("Value", "asset_update_tasks").Return(
			&regources.KeyValueEntryValue{
				Type: xdr.KeyValueEntryTypeString,
				Str:  &newValue,
			}, nil).Once()
		defer connector.AssertExpectations(t)
		defer keyValues.AssertExpectations(t)

		err := resourceKeyValueRead(d, m)
		assert.NoError(t, err)

		gotValueKey := d.Get("key").(string)
		assert.Equal(t, "asset_update_tasks", gotValueKey)

		gotValue := d.Get("value").(string)
		assert.Equal(t, "new_value", gotValue)

		gotValueType := d.Get("value_type").(string)
		assert.Equal(t, "string", gotValueType)
	})

	t.Run("new value", func(t *testing.T) {
		c := map[string]interface{}{
			"key":        "asset_update_tasks",
			"value_type": "uint32",
			"value":      "1",
		}
		newValue := uint32(12)
		d := schema.TestResourceDataRaw(t, keyValueSchema, c)
		connector.On("KeyValues").Return(&keyValues).Once()
		keyValues.On("Value", "asset_update_tasks").Return(
			&regources.KeyValueEntryValue{
				Type: xdr.KeyValueEntryTypeUint32,
				U32:  &newValue,
			}, nil).Once()
		defer connector.AssertExpectations(t)
		defer keyValues.AssertExpectations(t)

		err := resourceKeyValueRead(d, m)
		assert.NoError(t, err)

		gotValueKey := d.Get("key").(string)
		assert.Equal(t, "asset_update_tasks", gotValueKey)

		gotValue := d.Get("value").(string)
		assert.Equal(t, "12", gotValue)

		gotValueType := d.Get("value_type").(string)
		assert.Equal(t, "uint32", gotValueType)
	})
}
