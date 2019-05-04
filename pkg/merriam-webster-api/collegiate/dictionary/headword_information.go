package dictionary

// The HeadwordInformation is documented at https://dictionaryapi.com/products/json#sec-2.hwi
type HeadwordInformation struct {
	Headword       string           `json:"hw"`
	Pronunciations []*Pronunciation `json:"prs,omitempty"`
}
