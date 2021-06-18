package language

import (
	"context"
)

type contextKey string

const languageKey contextKey = "language"

// FromContext returns language tag. If language tag does not exist it returns language.Und value.
func FromContext(ctx context.Context) Tag {
	tag, _ := ctx.Value(languageKey).(Tag)

	return tag
}

// WithContext adds language tag to context.
func WithContext(ctx context.Context, tag Tag) context.Context {
	return context.WithValue(ctx, languageKey, tag)
}
