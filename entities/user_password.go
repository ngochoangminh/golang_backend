package entities

import (
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserPassword struct {
	id					uuid.UUID
	user_id				uuid.UUID
	password			string
	one_time_password	string
	status				string
	created_at			time.Time
	created_by			uuid.UUID
	updated_at			time.Time
	updated_by			uuid.UUID
	deleted_at			time.Time
	deleted_by			uuid.UUID
	is_deleted			bool
}

func CreateUserPassword(user_id uuid.UUID, password string) (*UserPassword, error) {
	pass_hashed, _ := HashPassword(password)
	up := &UserPassword{
		id: uuid.New(),
		user_id: user_id,
		password: pass_hashed,
		status: "ACTIVE",
		created_at: time.Now().UTC(),
		created_by: user_id,
		is_deleted: false,
	}

	return up, nil
}

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func (u *UserPassword) ValidatePassword(p string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.password), []byte(p))
	return err == nil
}