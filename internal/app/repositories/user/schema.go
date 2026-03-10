package userRepo

import (
	"strings"
	"time"
)

const tableUsers = "public.users"

const (
	columnID         = "id"
	columnEmail      = "email"
	columnPassword   = "password"
	columnName       = "name"
	columnIsVerified = "is_verified"
	columnCreatedAt  = "created_at"
	columnUpdatedAt  = "updated_at"
)

type user struct {
	Email      string    `db:"email"`
	Password   string    `db:"password"`
	Name       string    `db:"name"`
	IsVerified bool      `db:"is_verified"`
	CreatedAt  time.Time `db:"created_at"`
	UpdatedAt  time.Time `db:"updated_at"`
}

var (
	tableUsersColumns = []string{
		columnID,
		columnEmail,
		columnPassword,
		columnName,
		columnIsVerified,
		columnCreatedAt,
		columnUpdatedAt,
	}

	createUserColumns = []string{
		columnEmail,
		columnPassword,
		columnName,
	}
)

func (o *user) mapFields() map[string]any {
	return map[string]any{
		columnEmail:      o.Email,
		columnPassword:   o.Password,
		columnName:       o.Name,
		columnIsVerified: o.IsVerified,
		columnCreatedAt:  o.CreatedAt,
		columnUpdatedAt:  o.UpdatedAt,
	}
}

func (o *user) Values(columns ...string) []any {
	mapFields := o.mapFields()
	values := make([]any, 0, len(columns))
	for i := range columns {
		if v, ok := mapFields[columns[i]]; ok {
			values = append(values, v)
		} else {
			values = append(values, nil)
		}
	}
	return values
}

func (o *user) ReturningValues(columns ...string) string {
	return strings.Join(columns, ", ")
}