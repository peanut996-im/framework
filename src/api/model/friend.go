// Package model
// @Title  friend.go
// @Description
// @Author  peanut996
// @Update  peanut996  2021/5/22 10:05
package model

import (
	"errors"
	"framework/db"
	"framework/logger"
	"framework/tool"
	"go.mongodb.org/mongo-driver/bson"
)

type Friend struct {
	FriendA    string `json:"friendA" bson:"userA"`
	FriendB    string `json:"friendB" bson:"userB"`
	RoomID     string `json:"roomID" bson:"roomID"`
	CreateTime string `json:"-" bson:"createTime"`
}

func NewFriend() *Friend {
	return &Friend{
		CreateTime: tool.GetNowUnixMilliSecond(),
	}
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
	if _, err := mongo.InsertOne(MongoCollectionFriend, friend); err != nil {
		logger.Error("mongo insert friend err: %v", err)
		return err
	}
	//Symmetrical insertion
	oppositeFriend := NewFriend()
	oppositeFriend.FriendA = friend.FriendB
	oppositeFriend.FriendB = friend.FriendA
	oppositeFriend.RoomID = r.RoomID
	if _, err = mongo.InsertOne(MongoCollectionFriend, oppositeFriend); err != nil {
		logger.Error("mongo insert oppositefriend err: %v", err)
		return err
	}
	return nil
}

//AddNewFriend Add friends by UID
func AddNewFriend(friendA, friendB string) error {
	if _, err := FindFriend(friendA, friendB); nil == err {
		// already exists or error
		return errors.New("friend already exists or find error")
	}
	f := NewFriend()
	f.FriendA = friendA
	f.FriendB = friendB

	if err := insertFriend(f); nil != err {
		return err
	}
	return nil
}

func DeleteFriend(friendA, friendB string) error {
	mongo := db.GetLastMongoClient()
	// find room and delete
	friend, err := FindFriend(friendA, friendB)
	if err != nil {
		return err
	}
	if err := deleteRoom(friend.RoomID); nil != err {
		return err
	}
	filter := bson.M{"userA": friendA, "userB": friendB}
	if _, err = mongo.DeleteMany(MongoCollectionFriend, filter); err != nil {
		return err
	}
	filter = bson.M{"userB": friendA, "userA": friendB}
	if _, err = mongo.DeleteMany(MongoCollectionFriend, filter); err != nil {
		return err
	}
	return nil
}

func GetAllFriends(user string) ([]string, error) {
	mongo := db.GetLastMongoClient()
	friends := make([]string, 0)
	filterA := bson.M{"userA": user}
	friendsA := []Friend{}
	if err := mongo.Find(MongoCollectionFriend, &friendsA, filterA); nil != err {
		logger.Debug("Find friendB err: %v", err)
		return nil, err
	}
	for _, friend := range friendsA {
		friends = append(friends, friend.FriendB)
	}
	filterB := bson.M{"userB": user}
	friendsB := []Friend{}
	if err := mongo.Find(MongoCollectionFriend, &friendsB, filterB); nil != err {
		logger.Debug("Find friendA err: %v", err)
		return nil, err
	}
	for _, friend := range friendsB {
		friends = append(friends, friend.FriendA)
	}
	return tool.RemoveDuplicateString(friends), nil
}

func FindFriend(friendA, friendB string) (*Friend, error) {
	mongo := db.GetLastMongoClient()
	friend := &Friend{}
	filter := bson.M{
		"userA": friendA,
		"userB": friendB,
	}

	if err := mongo.FindOne(MongoCollectionFriend, friend, filter); err != nil {
		return nil, err
	}
	return friend, nil
}
