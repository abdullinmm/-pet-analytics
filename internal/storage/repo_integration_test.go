package storage

import (
	"context"
	"testing"
)

func TestRepo_SeededData(t *testing.T) {
	r, err := New(pgURL())
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	ctx := context.Background()

	// Сид положил две заметки пользователю 1: "hello", "world"
	got, err := r.GetNotes(ctx, 1)
	if err != nil {
		t.Fatalf("get notes: %v", err)
	}
	if len(got) < 2 || got[0] != "hello" || got[1] != "world" {
		t.Fatalf("unexpected seeded notes: %#v", got)
	}
}
