package collegiate_dictionary

type HeadwordInformation struct {
	Headword       string           `json:"hw"`
	Pronunciations []*Pronunciation `json:"prs,omitempty"`
}
