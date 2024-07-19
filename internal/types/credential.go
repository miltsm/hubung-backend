package types

type PublicKeyCred struct {
	ID            uint
	PublicKey     string
	PasskeyUserId uint
	BackedUp      bool
	Name          string
}
