package schemas

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type CreateUserRequest struct {
	ID             bson.ObjectID `bson:"-" json:"-" `
	UserID         string        `bson:"user_id" json:"user_id" binding:"required,uuid4"`
	FirstName      string        `bson:"first_name" json:"first_name" binding:"required,min=2,max=30"`
	LastName       string        `bson:"last_name" json:"last_name"  binding:"required,min=2,max=30"`
	PrimaryEmail   string        `bson:"email" json:"email" binding:"email,required"`
	SecondaryEmail string        `bson:"backup_email" json:"backup_email,omitempty" binding:"omitempty,email"`
	Password       string        `bson:"password" json:"password"  binding:"required"`
	CreatedAt      time.Time     `bson:"created_at" json:"created_at" `
	UpdatedAt      time.Time     `bson:"updated_at" json:"updated_at,omitempty"`
}

type PatchUserRequest struct {
	FirstName      string    `bson:"first_name" json:"first_name" binding:"required,min=2,max=30"`
	LastName       string    `bson:"last_name" json:"last_name"  binding:"required,min=2,max=30"`
	PrimaryEmail   string    `bson:"email" json:"email" binding:"email"`
	SecondaryEmail string    `bson:"backup_email" json:"backup_email,omitempty" binding:"omitempty,email"`
	UpdatedAt      time.Time `bson:"updated_at" json:"updated_at,omitempty"`
}
