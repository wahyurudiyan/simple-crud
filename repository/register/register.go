package register

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
	"github.com/wahyurudiyan/simple-crud/entity/register"
)

type repository struct {
	db *gorm.DB
}

type Repository interface {
	GetAccount(ctx context.Context) ([]*register.User, error)
	GetAccountByUniqueID(ctx context.Context, uniqueId string) ([]*register.User, error)
	CreateAccount(ctx context.Context, records []*register.User) error
	UpdateAccount(ctx context.Context, record *register.User) error
	DeleteAccount(ctx context.Context, record *register.User) error
}

func NewRepository(conn *gorm.DB) Repository {
	return &repository{conn}
}

func (r *repository) GetAccount(ctx context.Context) ([]*register.User, error) {
	var users []*register.User
	query := `SELECT * FROM account`

	rows, err := r.db.Exec(query).Rows()
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var user *register.User
		err := rows.Scan(user)
		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	defer rows.Close()

	return users, nil
}

func (r *repository) GetAccountByUniqueID(ctx context.Context, uniqueId string) (*register.User, error) {
	var user *register.User

	query := `SELECT * FROM account WHERE unique_id`
	row := r.db.Exec(query).Row()

	err := row.Scan(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *repository) CreateAccount(ctx context.Context, records []*register.User) error {
	var args []string
	query := `INSERT INTO simple_crud.account (unique_id, firstname, lastname, address, email, created_at, updated_at VALUES %s`

	for _, v := range records {
		format := fmt.Sprintf(`('%v', '%v', '%v', '%v', '%v', '%v', '%v')`,
			v.UniqueID, v.Firstname, v.Lastname, v.Address, v.Email, time.Now().Format(time.RFC3339), time.Now().Format(time.RFC3339))
		args = append(args, format)
	}

	arg := strings.Join()

	return nil
}

// func (r *repository) UpdateAccount(ctx context.Context, record *register.User) error {
// 	return
// }
// func (r *repository) DeleteAccount(ctx context.Context, record *register.User) error {
// 	return
// }
