package dictionary

// The Metadata describes information about the entry.
// See https://dictionaryapi.com/products/json#sec-2.meta for more information
type Metadata struct {
	ID        string   `json:"id,omitempty"`        // Unique entry identifier within a particular dictionary data set, used in cross-references pointing to the entry. It consists of the headword, and in homograph entries, an appended colon and homograph number.
	UUID      string   `json:"uuid,omitempty"`      // Universally unique identifier
	Sort      string   `json:"sort,omitempty"`      // a 9-digit code which may be used to sort entries in the proper dictionary order if alphabetical display is needed
	Src       string   `json:"src,omitempty"`       // Source data set for entry—ignore
	Section   string   `json:"section,omitempty"`   // Indicates the section the entry belongs to in print, where "alpha" indicates the main alphabetical section, "biog" biographical, "geog" geographical, and "fw&p" the foreign words & phrases section.
	Stems     []string `json:"stems,omitempty"`     // Lists all of the entry's headwords, variants, inflections, undefined entry words, and defined run-on phrases. Each stem string is a valid search term that should match this entry.
	Offensive bool     `json:"offensive,omitempty"` // True if there is a label containing "offensive" in the entry; otherwise, false.
}
