package model

type AuthRequest struct {
	IsAuthenticated bool
	Username string
	Email string
	DomainId string
	DomainName string
	Error error
}
