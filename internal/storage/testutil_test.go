package storage

import "os"

func pgURL() string {
	u := os.Getenv("PGURL")
	if u == "" {
		u = "postgres://app:app@127.0.0.1:5432/petdb?sslmode=disable"
	}
	return u
}
