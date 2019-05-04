package dictionary

// The AlternateHeadwordInformation is documented at https://dictionaryapi.com/products/json#sec-2.ahws
type AlternateHeadwordInformation struct {
	Headword             string           `json:"hw"`
	Pronunciations       []*Pronunciation `json:"prs,omitempty"`
	ParenthesizedSubject string           `json:"psl,omitempty"`
}
