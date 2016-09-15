package model

import (
	"github.com/gorilla/sessions"
)

type AuthRequest struct {
	IsAuthenticated bool
	Username string
	Email string
	Domains []string
	Error error
	SessionProject string
	SessionDomain string
}

func (ar AuthRequest) ValidateDomain(domain string) bool {
	for _ , dom := range ar.Domains {
		if domain == dom {
			return true
		}
	}
	return false
}

func (ar *AuthRequest) DeserializeFromSession(session *sessions.Session){
	var ok bool
	ar.Username ,ok = session.Values["Username"].(string)
	ar.Email ,ok = session.Values["Email"].(string)
	ar.SessionProject ,ok = session.Values["SessionProject"].(string)
	ar.SessionDomain ,ok = session.Values["SessionDomain"].(string)
	ar.Domains,ok = session.Values["Domains"].([]string)
	if ok && len(ar.Domains)>0 {
		ar.IsAuthenticated = true
	}else {
		ar.IsAuthenticated = false
	}
}

func (ar *AuthRequest) SerializeToSession(session *sessions.Session){
	session.Values["Username"] = ar.Username
	session.Values["Email"] = ar.Email
	session.Values["SessionProject"] = ar.SessionProject
	session.Values["SessionDomain"] = ar.SessionDomain
}

//type UserAuth struct {
//	ID     string `json:"id"`
//	Login  string `json:"login"`
//	Token  string `json:"-"`
//	Secret string `json:"-"`
//	Expiry int64  `json:"-"`
//	Email  string `json:"email"`
//	Avatar string `json:"avatar_url"`
//	Active bool   `json:"active,"`
//	Admin  bool   `json:"admin,"`
//	DomainId string
//	DomainName string
//	Hash   string `json:"-"`
//}
