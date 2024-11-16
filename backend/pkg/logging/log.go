package logging

import (
	"context"
	"fmt"
	"log/slog"
)

func additionalFields(ctx context.Context) []slog.Attr {
	var fields []slog.Attr
	for _, tag := range dynamicTags {
		value := ctx.Value(tag)
		if value != nil {
			fields = append(fields, slog.String(string(tag), fmt.Sprintf("%v", value)))
		}
	}
	return fields
}

func Debugf(ctx context.Context, format string, args ...interface{}) {
	fields := additionalFields(ctx)
	slog.DebugContext(ctx, fmt.Sprintf(format, args...), fields)
}

func Infof(ctx context.Context, format string, args ...interface{}) {
	fields := additionalFields(ctx)
	slog.InfoContext(ctx, fmt.Sprintf(format, args...), fields)
}

func Warnf(ctx context.Context, format string, args ...interface{}) {
	fields := additionalFields(ctx)
	slog.WarnContext(ctx, fmt.Sprintf(format, args...), fields)
}

func Errorf(ctx context.Context, format string, args ...interface{}) {
	fields := additionalFields(ctx)
	slog.ErrorContext(ctx, fmt.Sprintf(format, args...), fields)
}
