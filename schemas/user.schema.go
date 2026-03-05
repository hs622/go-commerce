package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreateUserRequest struct {
	ID             bson.ObjectID `bson:"-" json:"-" `
	UserID         string        `bson:"user_id" json:"user_id,omitempty" binding:"omitempty,uuid4"`
	FirstName      string        `bson:"first_name" json:"first_name,omitempty" binding:"required,min=2,max=30"`
	LastName       string        `bson:"last_name,omitempty" json:"last_name,omitempty" binding:"omitempty,min=2,max=30"`
	PrimaryEmail   string        `bson:"email" json:"email,omitempty" binding:"email,required"`
	SecondaryEmail string        `bson:"backup_email,omitempty" json:"backup_email,omitempty" binding:"omitempty,email"`
	Password       string        `bson:"password" json:"password,omitempty"  binding:"required"`
	CreatedAt      *time.Time    `bson:"created_at" json:"created_at,omitempty" `
	UpdatedAt      *time.Time    `bson:"updated_at" json:"updated_at,omitempty"`
}

type PatchUserRequest struct {
	FirstName      string     `bson:"first_name,omitempty" json:"first_name" binding:"omitempty,min=2,max=30"`
	LastName       string     `bson:"last_name,omitempty" json:"last_name,omitempty"  binding:"omitempty,min=2,max=30"`
	PrimaryEmail   string     `bson:"email,omitempty" json:"email" binding:"omitempty,email"`
	SecondaryEmail string     `bson:"backup_email,omitempty" json:"backup_email,omitempty" binding:"omitempty,email"`
	UpdatedAt      *time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}
