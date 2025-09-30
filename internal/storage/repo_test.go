// internal/storage/repo_test.go
package storage

import (
	"context"
	"testing"

	appdb "github.com/abdullinmm/pet-analytics/internal/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestRepo_NotesLifecycle(t *testing.T) {
	r, err := New(pgURL())
	if err != nil {
		t.Fatalf("connect: %v", err)
	}
	ctx := context.Background()

	// отдельный пул для sqlc на том же URL
	pool, err := pgxpool.New(ctx, pgURL())
	if err != nil {
		t.Fatalf("pool: %v", err)
	}
	defer pool.Close()
	q := appdb.New(pool)

	u, err := q.CreateUser(ctx, appdb.CreateUserParams{
		Email: "test@example.com",
		Name:  "Test User",
	})
	if err != nil {
		t.Fatalf("create user via sqlc: %v", err)
	}
	uid := u.ID

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
