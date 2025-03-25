package db

import (
	"context"
	"database/sql"
	"math/rand"
	"strconv"
	"time"

	"auth-service-2.0/internal/repository"
)

func Seed(db *sql.DB, count int) error {
	ctx := context.Background()

	userStore := &repository.UserStore{Db: db}
	sessionStore := &repository.SessionStore{Db: db}
	logStore := &repository.LoginLogsStore{Db: db}

	for i := 1; i <= count; i++ {
		user := repository.User{
			Username:  "user" + randomString(4),
			Password:  "pass" + randomString(4),
			CreatedAt: time.Now().Format(time.RFC3339),
			Email:     randomString(5) + "@example.com",
		}

		// assume Insert adds user.ID after insert
		if err := userStore.Create(ctx, &user); err != nil {
			return err
		}

		session := repository.Session{
			UserID:    user.ID,
			Token:     "token-" + randomString(10),
			CreatedAt: time.Now().Format(time.RFC3339),
			ExpiresAt: time.Now().Add(24 * time.Hour).Format(time.RFC3339),
		}

		if err := sessionStore.Create(ctx, &session); err != nil {
			return err
		}

		for j := 0; j < count; j++ {
			log := repository.Logs{
				UserID:    user.ID,
				IPAddress: "192.168.1." + randomInt(2, 254),
				Tags:      []string{"login", "seed"},
				CreatedAt: time.Now(),
			}
			if err := logStore.Create(ctx, &log); err != nil {
				return err
			}
		}
	}

	return nil
}

func randomString(n int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyz")
	s := make([]rune, n)
	for i := range s {
		s[i] = letters[rand.Intn(len(letters))]
	}
	return string(s)
}

func randomInt(min, max int) string {
	return strconv.Itoa(rand.Intn(max-min) + min)
}
