package collegiate_dictionary

type AlternateHeadwordInformation struct {
	Headword             string           `json:"hw"`
	Pronunciations       []*Pronunciation `json:"prs,omitempty"`
	ParenthesizedSubject string           `json:"psl,omitempty"`
}
