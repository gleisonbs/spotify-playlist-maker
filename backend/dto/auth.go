package dto

type Request struct {
	Form Form `json:"form"`
}

type Form struct {
	Code        string `json:"code"`
	RedirectUri string `json:"redirect_uri"`
	GrantType   string `json:"grant_type"`
}

type AuthResponse struct {
	AccessToken string `json:"access_token"`
}
