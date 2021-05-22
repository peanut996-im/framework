package model

import (
	"framework/db"
	"framework/encoding"
	"framework/logger"
	"go.mongodb.org/mongo-driver/bson"
)

//User means a people who use the system.
type User struct {
	UID      string `json:"uid,omitempty"`
	Account  string `json:"account"`
	Password string `json:"-"`
}

//NewUser returns a User who UID generate by snowlake Algorithm
func NewUser(account string, password string) *User {
	return &User{
		UID:      encoding.NewSnowFlakeID(),
		Account:  account,
		Password: password,
	}
}

func GetUserByAccount(account string) (*User, error) {
	mongo := db.GetLastMongoClient()
	filter := bson.M{"account": account}
	user := &User{}
	err := mongo.FindOne("User", user, filter)
	if err != nil {
		logger.Info("mongo get user from account err: %v, uid: %v", err, account)
		return nil, err
	}
	return user, nil
}

func GetUserByUID(uid string) (*User, error) {
	mongo := db.GetLastMongoClient()
	filter := bson.M{"uid": uid}
	user := &User{}
	err := mongo.FindOne("User", user, filter)
	if err != nil {
		logger.Info("mongo get user from uid err: %v, uid: %v", err, uid)
		return nil, err
	}
	return user, nil
}

func InsertUser(user *User) error {
	mongo := db.GetLastMongoClient()
	res, err := mongo.InsertOne("User", user)
	if err != nil {
		logger.Error("mongo insert user err: %v", err)
		return err
	}
	logger.Info("Mongo insert user success, id: %v", res.InsertedID)
	return nil
}
