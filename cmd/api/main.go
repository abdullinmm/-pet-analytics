package main

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	appdb "github.com/abdullinmm/pet-analytics/internal/db"
	"github.com/abdullinmm/pet-analytics/internal/platform/dbpool"
)

type health struct {
	Status string `json:"status"`
	Time   string `json:"time"`
}

type createUserReq struct {
	Email string `json:"email"`
	Name  string `json:"name"`
}

func main() {
	mux := http.NewServeMux()

	// healthz
	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(health{
			Status: "ok",
			Time:   time.Now().UTC().Format(time.RFC3339),
		})
	})

	// metrics
	mux.Handle("/metrics", promhttp.Handler())

	// DB connect
	ctx := context.Background()
	pool, err := dbpool.Connect(ctx)
	if err != nil {
		log.Fatalf("db connect: %v", err)
	}
	defer pool.Close()

	q := appdb.New(pool)

	// GET /users/{id}
	mux.HandleFunc("/users", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodPost:
			var in createUserReq
			if err := json.NewDecoder(r.Body).Decode(&in); err != nil {
				http.Error(w, "invalid json", http.StatusBadRequest)
				return
			}
			if !strings.Contains(in.Email, "@") || len(in.Name) == 0 || len(in.Name) > 100 {
				http.Error(w, "validation error", http.StatusBadRequest)
				return
			}
			u, err := q.CreateUser(ctx, appdb.CreateUserParams{
				Email: in.Email,
				Name:  in.Name,
			})
			if err != nil {
				http.Error(w, "conflict or db error", http.StatusConflict)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusCreated)
			_ = json.NewEncoder(w).Encode(u)
		case http.MethodGet:
			limit := int64(20)
			offset := int64(0)
			if v := r.URL.Query().Get("limit"); v != "" {
				if n, err := strconv.ParseInt(v, 10, 64); err == nil && n > 0 && n <= 1000 {
					limit = n
				}
			}
			if v := r.URL.Query().Get("offset"); v != "" {
				if n, err := strconv.ParseInt(v, 10, 64); err == nil && n >= 0 {
					offset = n
				}
			}
			users, err := q.ListUsers(ctx, appdb.ListUsersParams{
				Limit:  int32(limit),
				Offset: int32(offset),
			})
			if err != nil {
				http.Error(w, "db error", http.StatusInternalServerError)
				return
			}
			w.Header().Set("Content-Type", "application/json")
			_ = json.NewEncoder(w).Encode(users)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	addr := getenv("APP_ADDR", ":2112")
	log.Printf("listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, mux))
}

func getenv(k, def string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return def
}
