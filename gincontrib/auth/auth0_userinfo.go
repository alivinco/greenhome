package auth

import "time"

type Auth0UserInfo struct {
	Email string `json:"email"`
	EmailVerified bool `json:"email_verified"`
	ClientID string `json:"clientID"`
	UserID string `json:"user_id"`
	Picture string `json:"picture"`
	Nickname string `json:"nickname"`
	Identities []struct {
		UserID string `json:"user_id"`
		Provider string `json:"provider"`
		Connection string `json:"connection"`
		IsSocial bool `json:"isSocial"`
	} `json:"identities"`
	UpdatedAt time.Time `json:"updated_at"`
	CreatedAt time.Time `json:"created_at"`
	Name string `json:"name"`
	LastPasswordReset time.Time `json:"last_password_reset"`
	AppMetadata struct {
		BhubRole string `json:"bhub_role"`
		Domains []string `json:"domains"`
	} `json:"app_metadata"`
}
