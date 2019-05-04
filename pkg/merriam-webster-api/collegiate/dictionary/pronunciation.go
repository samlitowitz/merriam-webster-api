package dictionary

// The Pronunciation is documented at https://dictionaryapi.com/products/json#sec-2.prs
type Pronunciation struct {
	Written     string `json:"mw,omitempty"`  // Written pronunciation in Merriam-Webster format
	BeforeLabel string `json:"l,omitempty"`   // Pronunciation label before pronunciation
	AfterLabel  string `json:"l2,omitempty"`  // Pronunciation label after pronunciation
	Punctuation string `json:"pun,omitempty"` // Punctuation to separate pronunciation objects
	// Missing sound for now
}
