package models

// APIResponse structure used as api response.
type APIResponse struct {
	Name      string `json:"name"`
	Kind      string `json:"kind"`
	Date      string `json:"date"`
	Originapi string `json:"originapi"`
}
