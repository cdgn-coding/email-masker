package entities

type Email struct {
	From    string `validate:"required" json:"from"`
	To      string `validate:"required" json:"to,omitempty"`
	Subject string `validate:"required" json:"subject,omitempty"`
	Content string `validate:"required" json:"content,omitempty"`
	HTML    string `json:"HTML,omitempty"`
}
