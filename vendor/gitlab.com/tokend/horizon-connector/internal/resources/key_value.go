package resources

type KeyValue struct {
	Key  string `json:"key"`
	Type struct {
		Name  string `json:"name"`
		Value int    `json:"value"`
	} `json:"type"`
	UI32Value   int    `json:"ui32_value"`
	StringValue string `json:"string_value"`
}

func (v KeyValue) MustUint32() int {
	if &v.UI32Value == nil {
		panic("value is not uint32")
	}
	return v.UI32Value
}

func (v KeyValue) MustString() string {
	if &v.StringValue == nil {
		panic("value is not string")
	}
	return v.StringValue
}
