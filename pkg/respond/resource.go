package respond

// Resource describes a JSON-LD resource header
type Resource struct {
	Context string `json:"@context,omitempty"`
	Type    string `json:"@type,omitempty"`
	ID      string `json:"@id,omitempty"`
}

// Status describe a resource to represent an operational status
type Status struct {
	*Resource `json:",inline"`
	Code      int    `json:"code"`
	Message   string `json:"message"`
}
