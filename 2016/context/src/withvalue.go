package main

import (
	"context"
	"net/http/httptest"
)

// START OMIT
type key int // unexported なキー型

var locationKey = 0

func locationMiddleware(h http.HandlerFunc, w http.RewsponseWriter, r *http.Request) {
	location := getLocation(r)
	r = r.WithContext(context.WithValue(r.Context(), locationKey, location)) //値の格納 // HL

	h(w, r)
}

func handle(w http.ResponseWriter, r *http.Request) {
	l := r.Value(locationKey).(Location) // 値の取得 // HL
	// ...
}

// END OMIT

var ts = httptest.NewServer(locationMiddleware(handle))

type User struct {
	name string
	age  int
}

type key int // unexportedなcontext key型

var usreKey key = 0

func NewContext(ctx context.Context, u *User) context.Context {
	return context.WithValue(ctx, userKey, u) // contextに値を格納 // HL
}

func FromContext(ctx context.Context) (*User, bool) {
	u, ok := ctx.Value(userKey).(*User) // contextから値を取得 // HL
	return u, ok
}

// END OMIT
