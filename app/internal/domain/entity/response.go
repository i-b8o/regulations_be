package entity

type Response struct {
	ID       string   `json:"id,omitempty"`
	Err      error    `json:"err,omitempty"`
	Errors   []string `json:"errors,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	Message  string   `json:"message,omitempty"`
}
