package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Username  string             `bson:"username"`
	Password  string             `bson:"password"`
	RoleID    primitive.ObjectID `bson:"role_id"`
	SubRoleID primitive.ObjectID `bson:"sub_role_id"`
}
