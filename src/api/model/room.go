package model

import (
	"framework/db"
	"framework/encoding"
	"framework/logger"
	"time"
)

type Room struct {
	RoomID     string `json:"room_id"`
	OneToOne   bool   `json:"one_to_one"`
	Status     string `json:"status"`
	CreateTime int64  `json:"create_time,omitempty"`
}

func newRoom() *Room {
	return &Room{
		RoomID:     encoding.NewSnowFlakeID(),
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

func InsertRoom(room *Room) error {
	mongo := db.GetLastMongoClient()
	res, err := mongo.InsertOne("Room", room)
	if err != nil {
		logger.Error("mongo insert room err: %v", err)
		return err
	}
	logger.Info("Mongo insert room success, id: %v", res.InsertedID)
	return nil
}
