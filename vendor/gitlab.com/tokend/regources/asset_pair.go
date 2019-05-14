package regources

type AssetPair struct {
	Base                    string   `json:"base"`
	Quote                   string   `json:"quote"`
	CurrentPrice            Amount   `json:"current_price"`
	PhysicalPrice           Amount   `json:"physical_price"`
	PhysicalPriceCorrection Amount   `json:"physical_price_correction"`
	MaxPriceStep            Amount   `json:"max_price_step"`
	Policy                  int32    `json:"policy"`
	Policies                []Policy `json:"policies"`
}

func (p AssetPair) GetLoganFields() map[string]interface{} {
	return map[string]interface{}{
		"base":                      p.Base,
		"quote":                     p.Quote,
		"current_price":             p.CurrentPrice,
		"physical_price":            p.PhysicalPrice,
		"physical_price_correction": p.PhysicalPriceCorrection,
		"max_price_step":            p.MaxPriceStep,
		"policy":                    p.Policy,
		"policies":                  p.Policies,
	}
}

type Policy struct {
	Name  string `json:"name"`
	Value int32  `json:"value"`
}

func (p Policy) GetLoganFields() map[string]interface{} {
	return map[string]interface{}{
		"name":  p.Name,
		"value": p.Value,
	}
}
