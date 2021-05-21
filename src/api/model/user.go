package model

import (
	"github.com/bwmarrin/snowflake"
)

//User means a people who use the system.
type User struct {
	UID      string
	Account  string
	Password string
}

//NewUser returns a User who UID generate by snowlake Algorithm
func NewUser(account string, password string) *User {
	node, _ := snowflake.NewNode(1)
	return &User{
		UID:      node.Generate().String(),
		Account:  account,
		Password: password,
	}
}
