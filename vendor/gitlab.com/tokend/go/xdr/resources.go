package xdr

// package contains helper struct for xdr enums marshaling

type value struct {
	Value int32 `json:"value"`
}

type enum struct {
	Value  int32  `json:"value"`
	String string `json:"name"`
}

type flag struct {
	Value int32       `json:"value"`
	Flags []flagValue `json:"flags"`
}

type flagValue struct {
	Name  string `json:"name"`
	Value int32  `json:"value"`
}
