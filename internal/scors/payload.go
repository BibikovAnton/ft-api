package scors

type AccountsCreateRequest struct {
	Name     string  `json:"name" validate:"required"`
	Type     string  `json:"type" validate:"required"`
	Balance  float64 `json:"balance" validate:"required"`
	Currency string  `json:"currency" validate:"required"`
}

type AccountsUpdateRequest struct {
	Name     string  `json:"name" `
	Type     string  `json:"type" `
	Balance  float64 `json:"balance" `
	Currency string  `json:"currency" `
}
