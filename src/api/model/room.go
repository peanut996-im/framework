package model

import (
	"framework/db"
	"framework/logger"
	"framework/tool"
	"time"
)

type Room struct {
	RoomID     string `json:"roomID" bson:"roomID"`
	OneToOne   bool   `json:"oneToOne" bson:"oneToOne"`
	Status     string `json:"status" bson:"status"`
	CreateTime int64  `json:"createTime,omitempty" bson:"createTime"`
}

func newRoom() *Room {
	return &Room{
		RoomID:     tool.NewSnowFlakeID(),
		CreateTime: time.Now().Unix(),
		OneToOne:   false,
	}
}

func NewGroupRoom() *Room {
	r := newRoom()
	r.OneToOne = false
	return r
}

func NewFriendRoom() *Room {
	r := newRoom()
	r.OneToOne = true
	return r
}

func insertRoom(room *Room) error {
	mongo := db.GetLastMongoClient()
	res, err := mongo.InsertOne("Room", room)
	if err != nil {
		logger.Error("mongo insert room err: %v", err)
		return err
	}
	logger.Info("Mongo insert room success, id: %v", res.InsertedID)
	return nil
}
