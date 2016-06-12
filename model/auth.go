package model

type AuthRequest struct {
	IsAuthenticated bool
	Username string
	Email string
	DomainId string
	DomainName string
	Error error
}

type UserAuth struct {
	ID     string `json:"id"`
	Login  string `json:"login"`
	Token  string `json:"-"`
	Secret string `json:"-"`
	Expiry int64  `json:"-"`
	Email  string `json:"email"`
	Avatar string `json:"avatar_url"`
	Active bool   `json:"active,"`
	Admin  bool   `json:"admin,"`
	DomainId string
	DomainName string
	Hash   string `json:"-"`
}
