package auth

import "context"

// ctxKey is a private type to avoid collisions
type ctxKey string

var (
	AdminClaimsKey = ClaimNewKey("adminClaims")
	UserClaimsKey  = ClaimNewKey("userClaims")
	AdminJwtKey    = []byte("Admin_secret_key")
	UserjwtKey     = []byte("my_secret_key")
)

// 1️⃣ create base context

// NewKey creates a new unique key
func ClaimNewKey(name string) ctxKey {
	return ctxKey(name)
}

// Set claims (pointer)
func ClaimSet[T any](ctx context.Context, key ctxKey, claims *T) context.Context {
	return context.WithValue(ctx, key, claims)
}

// Get claims (pointer)
func ClaimGet[T any](ctx context.Context, key ctxKey) (*T, bool) {
	val := ctx.Value(key)
	if val == nil {
		return nil, false
	}
	c, ok := val.(*T)
	return c, ok
}

// Update claims safely
func ClaimUpdate[T any](ctx context.Context, key ctxKey, updateFunc func(c *T)) context.Context {
	claims, ok := ClaimGet[T](ctx, key)
	if !ok || claims == nil {
		// initialize if missing
		claims = new(T)
	}
	updateFunc(claims)
	return ClaimSet(ctx, key, claims)
}
