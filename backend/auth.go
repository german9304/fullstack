package fullstack_backend

import (
	"net/http"
)

type Auth struct {
	RW http.ResponseWriter
	RQ *http.Request
}

func (a *Auth) setSession(name, value string) {
}

func (a *Auth) clearSession() {

}
