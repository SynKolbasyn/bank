package repository

import (
	"context"
	"database/sql"
	"errors"
	"net/http"

	"github.com/SynKolbasyn/bank/gateway/internal/domain"
	"github.com/google/uuid"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

type User struct {
	pool *pgxpool.Pool
}

func NewUser(pool *pgxpool.Pool) *User {
	return &User{
		pool: pool,
	}
}

func (u *User) Create(ctx context.Context, email, password string) (uuid.UUID, error) {
	query := `
		INSERT INTO users (email, password_hash)
		VALUES ($1::VARCHAR(256), $2::VARCHAR(256))
		RETURNING id;
	`

	var userID uuid.UUID

	err := u.pool.QueryRow(ctx, query, email, password).Scan(&userID)
	if err != nil {
		var pgErr *pgconn.PgError
		statusCode := http.StatusInternalServerError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			statusCode = http.StatusConflict
		}
		return uuid.UUID{}, domain.NewAppError(statusCode, err)
	}

	return userID, nil
}

func (u *User) Get(ctx context.Context, email string) (uuid.UUID, string, error) {
	query := `
		SELECT id, password_hash
		FROM users
		WHERE email = $1::VARCHAR(256);
	`

	var userID uuid.UUID
	var password string

	err := u.pool.QueryRow(ctx, query, email).Scan(&userID, &password)
	if err != nil {
		statusCode := http.StatusInternalServerError
		if errors.Is(err, sql.ErrNoRows) {
			statusCode = http.StatusNotFound
		}
		return uuid.UUID{}, "", domain.NewAppError(statusCode, err)
	}

	return userID, password, nil
}
