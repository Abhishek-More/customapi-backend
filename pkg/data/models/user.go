package models

type User struct {
	ID       int     `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string  `json:"name,omitempty" bson:"name,omitempty"`
	Username string  `json:"username,omitempty" bson:"username,omitempty"`
	Address  Address `json:"address,omitempty" bson:"address,omitempty"`
}

type Address struct {
	Street string `json:"street,omitempty" bson:"street,omitempty"`
	City string `json:"city,omitempty" bson:"city,omitempty"`
	State string `json:"state,omitempty" bson:"state,omitempty"`
}
