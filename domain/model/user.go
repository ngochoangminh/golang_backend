
package model

import (
	"time"
	"github.com/google/uuid"
)

type User struct {
	ID        	uuid.UUID `gorm:"primary_key" json:"id"`
	FirstName	string    `json:"first_name"`
	LastName  	string    `json:"last_name"`
	Email 	  	string	  `json:"email"`
	PhoneNumber string 	  `json:"phone_number"`
	Username 	string	  `json:"username"`
	Role       	string    `json:"Role"`
	CreatedAt 	time.Time `json:"created_at"`
	CreatedBy	uuid.UUID `json:"created_by"`
	UpdatedAt 	time.Time `json:"updated_at"`
	UpdatedBy	uuid.UUID `json:"updated_by"`
	DeletedAt 	time.Time `json:"deleted_at"`
	DeletedBy	uuid.UUID `json:"deleteted_by"`
	IsDeleted	bool	  `json:"is_deleted"`
}

func (User) TableName() string { return "users" }