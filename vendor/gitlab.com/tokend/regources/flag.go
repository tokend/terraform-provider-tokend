package regources

// Flag represents one value of binary mask
type Flag struct {
	Name  string `json:"name,omitempty"`
	Value int32  `json:"value"`
}
