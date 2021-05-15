package repository

import (
	"app/domain"
	"context"
	"database/sql"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]domain.User, error)
}

type userRepository struct {
	Conn *sql.DB
}

func NewUserRepository(Conn *sql.DB) UserRepository {
	return &userRepository{Conn}
}

// ここは共通化してーなあ
func (r *userRepository) userQuery(ctx context.Context, query string, args ...interface{}) ([]domain.User, error) {
	rows, err := r.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users := make([]domain.User, 0)
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(
			&user.Id,
			&user.LastName,
			&user.FirstName,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *userRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
	query := `SELECT * FROM users`
	return r.userQuery(ctx, query)
}
