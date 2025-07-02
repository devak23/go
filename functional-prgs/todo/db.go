package todo

import "os"

type authorizationFunc func() bool

type Db struct {
	AuthorizationFn authorizationFunc
}

func (db *Db) IsAuthorized() bool {
	return db.AuthorizationFn()
}

func argsAuthorization() bool {
	user := os.Args[1]
	if user == "admin" {
		return true
	}
	return false
}

func NewDb() *Db {
	return &Db{
		AuthorizationFn: argsAuthorization,
	}
}
