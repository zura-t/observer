package token

type Maker interface {
	CreateToken(payload *Payload) (string, *Payload, error)

	VerifyToken(token string) (*Payload, error)
}
