package dictionary

// The Inflection is documented at https://dictionaryapi.com/products/json#sec-2.ins
type Inflection struct {
	Full           string           `json:"if,omitempty"`  // A fully spelled-out inflection
	Ending         string           `json:"ifc,omitempty"` // An inflection ending (eg, "-ing")
	Label          string           `json:"il,omitempty"`  // Label, such as “also”, “plural”, “or”
	Pronunciations []*Pronunciation `json:"prs,omitempty"`
	SPL            string           `json:"spl,omitempty"` // This label provides information on the grammatical number (eg, singular, plural) of an inflection in a particular sense. A sense-specific inflection plural label is contained in an spl.
}
