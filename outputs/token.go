package outputs

type Token struct {
	Object             string  `json:"object"`
	ID                 string  `json:"id"`
	Used               bool    `json:"used"`
	IP                 *string `json:"ip"`
	TokenizationMethod *string `json:"tokenization_method"`
	CreatedAt          int     `json:"created_at"`
	UpdatedAt          *int    `json:"updated_at"`
	Card               struct {
		Data Card `json:"data"`
	} `json:"card"`
}

type TokenResponse struct {
	Data Token `json:"data"`
	Meta Meta  `json:"meta"`
}
