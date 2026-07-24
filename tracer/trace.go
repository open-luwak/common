package tracer

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"time"
)

type traceKey struct{}

type TraceMeta struct {
	TraceID  string
	SpanID   string
	ParentID string
}

type Span struct {
	ctx       context.Context
	name      string
	startTime time.Time
	duration  time.Duration
	meta      TraceMeta
}

func Start(ctx context.Context, name string) (context.Context, *Span) {
	parentMeta, hasParent := FromContext(ctx)

	var traceID string
	var parentID string

	if hasParent && parentMeta.TraceID != "" {
		traceID = parentMeta.TraceID
		parentID = parentMeta.SpanID
	} else {
		traceID = generateID(16)
		parentID = ""
	}

	currentSpanID := generateID(8)

	meta := TraceMeta{
		TraceID:  traceID,
		SpanID:   currentSpanID,
		ParentID: parentID,
	}

	newCtx := context.WithValue(ctx, traceKey{}, meta)

	span := &Span{
		ctx:       newCtx,
		name:      name,
		startTime: time.Now(),
		meta:      meta,
	}

	return newCtx, span
}

func (s *Span) End() time.Duration {
	if s.duration == 0 {
		s.duration = time.Since(s.startTime)
	}
	return s.duration
}

func (s *Span) TraceMeta() TraceMeta {
	return s.meta
}

func (s *Span) Name() string {
	return s.name
}

func (s *Span) StartTime() time.Time {
	return s.startTime
}

func FromContext(ctx context.Context) (TraceMeta, bool) {
	if ctx == nil {
		return TraceMeta{}, false
	}
	sc, ok := ctx.Value(traceKey{}).(TraceMeta)
	return sc, ok
}

func WithTraceID(ctx context.Context, traceID string) context.Context {
	if len(traceID) > 36 {
		traceID = traceID[:36]
	}
	meta := TraceMeta{
		TraceID:  traceID,
		SpanID:   "",
		ParentID: "",
	}
	return context.WithValue(ctx, traceKey{}, meta)
}

func generateID(bytes int) string {
	b := make([]byte, bytes)
	_, _ = rand.Read(b)
	return hex.EncodeToString(b)
}
