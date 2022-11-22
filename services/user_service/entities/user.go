package entities

import (
	"time"

	"github.com/google/uuid"
	"github.com/ngochoangminh/golang_backend/core/util"
)

type User struct {
	id				uuid.UUID
	first_name		string
	last_name		string
	username		string
	email			string
	phone_number	string
	role			string
	created_at		time.Time
	created_by		uuid.UUID
	updated_at		time.Time
	updated_by		uuid.UUID
	deleted_at		time.Time
	deleted_by		uuid.UUID
	is_deleted		bool
}

func NewUser (first_name string, last_name string, username string, email string, phone_number string, role string) (*User, error) {
	u := &User{
		id: uuid.New(),
		first_name: first_name,
		last_name: last_name,
		username: username,
		email: email,
		phone_number: phone_number,
		role: role,
		created_at: time.Now().UTC(),
		is_deleted: false,
	}

	return u, nil
} 

func (u *User) Validate() error {
	if u.email == "" || u.first_name == "" || u.last_name == "" || u.username == "" {
		return util.ErrInvalidEntity
	}
	
	return nil
}