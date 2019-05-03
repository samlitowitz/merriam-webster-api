package merriam_webster_api

import cd "github.com/samlitowitz/merriam-webster-api/pkg/merriam-webster-api/collegiate-dictionary"

type CollegiateDictionary struct {
	Metadata                     *cd.Meta                           `json:"meta,omitempty"` // Entry metadata is information about the entry, as opposed to the actual content of the entry.
	HomographNumber              int                                `json:"hom,omitempty"`  // Homographs are headwords with identical spellings but distinct meanings and origins. The hom member contains a homograph number used to mark and distinguish the identically-spelled headwords.
	HeadwordInformation          *cd.HeadwordInformation            `json:"hwi,omitempty"`  // The headword is the word being defined or translated in a dictionary entry. Key headword information is grouped in an hwi object. This includes the headword itself in the hw member, and may include one or more pronunciations in prs.
	AlternateHeadwordInformation []*cd.AlternateHeadwordInformation `json:"ahws,omitempty"` // An alternate headword is a regional or less common spelling of a headword. A set of one or more alternate headwords is contained in an ahws.
	Variants                     []*cd.Variant                      `json:"vrs,omitempty"`  // A variant is a different spelling or styling of a headword, defined run-on phrase, or undefined entry word. A set of one or more variants is contained in a vrs.
	Pronuncians                  []*cd.Pronunciation                `json:"prs,omitempty"`  // A pronunciation specifies how a written word is pronounced aloud. A set of pronunciations is encoded in a prs array. Each element represents a distinct pronunciation of a headword or other term.
	// Labels
	FunctionLabel string                      `json:"fl,omitempty"`    // The functional label describes the grammatical function of a headword or undefined entry word, such as "noun" or "adjective".
	GeneralLabels []string                    `json:"lbs,omitempty"`   // General labels provide information such as whether a headword is typically capitalized, used as an attributive noun, etc. A set of one or more such labels is contained in an lbs.
	SLS           []string                    `json:"sls,omitempty"`   // A subject/status label describes the subject area (eg, "computing") or regional/usage status (eg, "British", "formal", "slang") of a headword or a particular sense of a headword. A set of one or more subject/status labels is contained in an sls.
	PSL           string                      `json:"psl,omitempty"`   // A parenthesized subject/status label describes the subject area or regional/usage status (eg, "British") of a headword or other defined term, and is displayed in parentheses. The parenthesized subject/status label is contained in a psl.
	SPL           string                      `json:"spl,omitempty"`   // This label provides information on the grammatical number (eg, singular, plural) of an inflection in a particular sense. A sense-specific inflection plural label is contained in an spl.
	SGRAM         string                      `json:"sgram,omitempty"` // This label notes whether a particular sense of a verb is transitive (T) or intransitive (I). The sense-specific grammatical label is contained in an sgram.
	Inflections   []*cd.Inflection            `json:"ins,omitempty"`   // Inflection is the change of form that words undergo in different grammatical contexts, such as tense or number. A set of one or more inflections is contained in an ins.
	CXS           []*cd.CognateCrossReference `json:"cxs,omitempty"`   // When a headword is a less common spelling of another word with the same meaning, there will be a cognate cross-reference pointing to the headword with the more common spelling. A set of cognate cross-references is contained in a cxs.
	// Sense
	// https://dictionaryapi.com/products/json#sec-2.sense-struct
}
