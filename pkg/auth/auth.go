package auth

import "context"

type contextKey string

const (
	authorizationContextKey = contextKey("Authorization")
)

func getContextValueString(ctx context.Context, key contextKey) string {
	var value string
	if ctx != nil {
		raw := ctx.Value(key)
		if raw != nil {
			parsed, ok := raw.(string)
			if ok {
				value = parsed
			}
		}
	}
	return value
}

func SetAuthorizationToken(ctx context.Context, token string) context.Context {
	return context.WithValue(ctx, authorizationContextKey, token)
}

func GetAuthorizationToken(ctx context.Context) string {
	return getContextValueString(ctx, authorizationContextKey)
}
