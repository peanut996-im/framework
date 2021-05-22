// Package model
// @Title  friend.go
// @Description
// @Author  peanut996
// @Update  peanut996  2021/5/22 10:05
package model

import (
	"framework/api"
	"framework/db"
	"framework/logger"
	"go.mongodb.org/mongo-driver/bson"
)

type Friend struct {
	FriendA string `json:"friendA" bson:"userA"`
	FriendB string `json:"friendB" bson:"userB"`
	RoomID  string `json:"roomID" bson:"roomID"`
}

func NewFriend() *Friend {
	return &Friend{}
}

func insertFriend(friend *Friend) error {
	mongo := db.GetLastMongoClient()
	r := NewFriendRoom()
	friend.RoomID = r.RoomID
	// First try to insert the room
	err := insertRoom(r)
	if nil != err {
		logger.Error("InsertFriendRoom err: %v", err)
		return err
	}
	// Second try to insert the friend relationship
	if _, err := mongo.InsertOne("Friend", friend); err != nil {
		logger.Error("mongo insert friend err: %v", err)
		return err
	}
	//Symmetrical insertion
	oppositeFriend := NewFriend()
	oppositeFriend.FriendA = friend.FriendB
	oppositeFriend.FriendB = friend.FriendA
	if _, err = mongo.InsertOne("Friend", oppositeFriend); err != nil {
		logger.Error("mongo insert oppositefriend err: %v", err)
		return err
	}
	return nil
}

//AddNewFriend Add friends by UID
func AddNewFriend(friendA, friendB string) error {
	f := NewFriend()
	f.FriendA = friendA
	f.FriendB = friendB
	if err := insertFriend(f); nil != err {
		return err
	}
	return nil
}

func DeleteFriend(origin, target string) error {
	mongo := db.GetLastMongoClient()
	filter := bson.D{{
		"userA",
		bson.D{{
			"$in",
			bson.A{origin, target},
		}},
	}}
	_, err := mongo.DeleteMany("Friend", filter)
	if err != nil {
		return nil
	}
	return nil
}

func GetAllFriends(user string) ([]string, error) {
	mongo := db.GetLastMongoClient()
	friends := make([]string, 0)
	filterA := bson.M{"userA": user}
	friendsA := []Friend{}
	if err := mongo.Find("Friend", &friendsA, filterA); nil != err {
		logger.Debug("Find friendB err: %v", err)
		return nil, err
	}
	for _, friend := range friendsA {
		friends = append(friends, friend.FriendB)
	}
	filterB := bson.M{"userB": user}
	friendsB := []Friend{}
	if err := mongo.Find("Friend", &friendsB, filterB); nil != err {
		logger.Debug("Find friendA err: %v", err)
		return nil, err
	}
	for _, friend := range friendsB {
		friends = append(friends, friend.FriendA)
	}
	return api.RemoveDuplicateString(friends), nil
}
