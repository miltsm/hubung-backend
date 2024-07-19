package types

import "github.com/go-webauthn/webauthn/webauthn"

type UserEntity struct {
	UserID        string
	Email         string
	Username      string
	PasskeyUserId uint
}

type User struct {
	Id          []byte
	Name        string
	DisplayName string
	Credentials []webauthn.Credential
}

func (u *User) WebAuthnID() []byte {
	return u.Id
}

func (u *User) WebAuthnName() string {
	return u.Name
}

func (u *User) WebAuthnDisplayName() string {
	return u.DisplayName
}

func (u *User) WebAuthnCredentials() []webauthn.Credential {
	return u.Credentials
}

func (u *User) WebAuthnIcon() string {
	return ""
}
