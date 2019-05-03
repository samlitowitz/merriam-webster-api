package collegiate_dictionary

type CognateCrossReference struct {
	Label   string                         `json:"cxl,omitempty"`
	Targets []*CognateCrossReferenceTarget `json:"cxtis,omitempty"` // Of one or more cognate cross-reference targets
}

type CognateCrossReferenceTarget struct {
	Label       string `json:"cxl,omitempty"`
	ReferenceID string `json:"cxr,omitempty"` // When present, use as cross-reference target ID for immediately preceding cxt
	URI         string `json:"cxt,omitempty"` // Provides hyperlink text in all cases, and cross-reference target ID when no immediately following cxr
	Target      string `json:"cxn,omitempty"` // Sense number of cross-reference target
}
