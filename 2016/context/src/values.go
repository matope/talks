package main

import "context"

// START OMIT
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
