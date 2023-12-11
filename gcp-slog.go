package gcp_slog

import (
	"log/slog"
	"strings"
)

func NewGCPHandlerOptions() (*slog.HandlerOptions, *slog.LevelVar) {
	var programLevel = new(slog.LevelVar)
	return &slog.HandlerOptions{
		ReplaceAttr: gcpSlogReplaceAttr,
		Level:       programLevel,
	}, programLevel
}

func gcpSlogReplaceAttr(group []string, a slog.Attr) slog.Attr {
	if a.Key == "level" {
		return slog.Attr{Key: "severity", Value: slog.StringValue(strings.ToLower(a.Value.String()))}
	}

	if a.Key == "time" {
		return slog.Attr{Key: a.Key, Value: slog.AnyValue(a.Value.Time().UnixMilli())}
	}

	if a.Key == "msg" {
		return slog.Attr{Key: "message", Value: a.Value}
	}

	return slog.Attr{Key: a.Key, Value: a.Value}
}
