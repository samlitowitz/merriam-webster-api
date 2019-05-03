package collegiate_dictionary

type Variant struct {
	Variant        string           `json:"va"`
	VariantLabel   string           `json:"vl,omitempty"`
	Pronunciations []*Pronunciation `json:"prs,omitempty"`
	SPL            string           `json:"spl,omitempty"` // This label provides information on the grammatical number (eg, singular, plural) of an inflection in a particular sense. A sense-specific inflection plural label is contained in an spl.
}
