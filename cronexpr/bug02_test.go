package cronexpr

import (
	"context"
	"testing"
	"time"
)

func TestBug02_NextNRespectsCancel(t *testing.T) {
	e, err := Parse("*/1 * * * *")
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()
	out, err := e.NextNWithContext(ctx, time.Date(2024, 6, 1, 0, 0, 0, 0, time.UTC), 5)
	if err == nil {
		t.Fatal("expected cancel error")
	}
	if out != nil {
		t.Fatalf("partial results leaked: %v", out)
	}
}
