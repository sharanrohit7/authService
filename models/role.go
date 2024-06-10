package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Role struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	RoleName    string             `bson:"role_name"`
	SubRole     string             `bson:"sub_role"`
	Permissions []string           `bson:"permissions"`
}
