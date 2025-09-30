package storage

import (
	"context"
	"os"
	"testing"
)

func pgURL() string {
	u := os.Getenv("PGURL")
	if u == "" {
		// локальный дефолт для удобства
		u = "postgres://app:app@127.0.0.1:5432/petdb?sslmode=disable"
	}
	return u
}

func TestRepo_NotesLifecycle(t *testing.T) {
	r, err := New(pgURL())
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	ctx := context.Background()

	uid, err := r.CreateUser(ctx, "test@example.com")
	if err != nil {
		t.Fatalf("create user: %v", err)
	}

	if _, err := r.AddNote(ctx, uid, "hello"); err != nil {
		t.Fatalf("add note: %v", err)
	}
	if _, err := r.AddNote(ctx, uid, "world"); err != nil {
		t.Fatalf("add note: %v", err)
	}

	notes, err := r.GetNotes(ctx, uid)
	if err != nil {
		t.Fatalf("get notes: %v", err)
	}
	if len(notes) != 2 || notes[0] != "hello" || notes[1] != "world" {
		t.Fatalf("unexpected notes: %#v", notes)
	}
}
