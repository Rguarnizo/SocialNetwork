package auth

import "net/http"

type contextKey string

const UserCtxKey contextKey = "auth_user"

func UserFromRequest(r *http.Request) User {
	return r.Context().Value(UserCtxKey).(User)
}
