// Package model
// @Title  friend.go
// @Description
// @Author  peanut996
// @Update  peanut996  2021/5/22 10:05
package model

import (
	"framework/db"
	"framework/logger"
)

type Friend struct {
	FriendA string `json:"friend_a"`
	FriendB string `json:"friend_b"`
	RoomID  string `json:"room_id"`
}

func NewFriend() *Friend {
	return &Friend{}
}

func InsertFriend(friendA, friendB string) error {
	f := NewFriend()
	f.FriendA = friendA
	f.FriendB = friendB
	r := NewFriendRoom()
	f.RoomID = r.RoomID
	// First try to insert the room
	err := InsertRoom(r)
	if nil != err {
		logger.Error("InsertFriendRoom err: %v", err)
		return err
	}
	// Second try to insert the friend relationship
	err = insertFriend(f)
	if nil != err {
		logger.Error("InsertFriend err: %v", err)
		return err
	}

	return nil
}

func insertFriend(friend *Friend) error {
	mongo := db.GetLastMongoClient()
	res, err := mongo.InsertOne("Friend", friend)
	if err != nil {
		logger.Error("mongo insert friend err: %v", err)
		return err
	}
	logger.Info("Mongo insert friend success, id: %v", res.InsertedID)
	return nil
}
