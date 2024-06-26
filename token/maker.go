package token

import "time"

// Maker is an interface for mannaging tokens.
type Maker interface {
	// CreateToken creates a new token for the given username with the given duration.
	CreateToken(username string, role string, duration time.Duration) (string, *Payload, error)

	// VerifyTokenchecks if the token is vaild or not.
	VerifyToken(token string) (*Payload, error)
}
