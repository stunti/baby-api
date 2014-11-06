package global

import r "github.com/dancannon/gorethink"

var (
	PrivateKey []byte
	PublicKey  []byte
	Session    *r.Session
)
