// Package model
// @Title  group.go
// @Description
// @Author  peanut996
// @Update  peanut996  2021/5/22 17:23
package model

import (
	"framework/db"
	"framework/tool"
	"go.mongodb.org/mongo-driver/bson"
)

type Group struct {
	RoomID     string `json:"roomID" bson:"roomID"`
	GroupID    string `json:"groupID" bson:"groupID"`
	GroupName  string `json:"groupName" bson:"groupName"`
	GroupAdmin string `json:"groupAdmin" bson:"groupAdmin"`
	CreateTime string `json:"-" bson:"createTime"`
}

type GroupData struct {
	Group    `json:",inline"`
	Members  []*User        `json:"members"`
	Messages []*ChatMessage `json:"messages"`
}

func NewGroup() *Group {
	return &Group{
		GroupID:    tool.NewSnowFlakeID(),
		CreateTime: tool.GetNowUnixMilliSecond(),
	}
}

func insertGroup(g *Group) error {
	mongo := db.GetLastMongoClient()
	r := NewGroupRoom()
	g.RoomID = r.RoomID
	g.GroupID = r.RoomID
	// First try to insert the room
	if err := insertRoom(r); nil != err {
		return err
	}
	// Second try to insert group
	_, err := mongo.InsertOne(MongoCollectionGroup, g)
	if err != nil {
		return err
	}
	return nil
}

func CreateGroup(name, admin string) error {
	g := NewGroup()
	g.GroupAdmin = admin
	g.GroupName = name
	if err := insertGroup(g); nil != err {
		return err
	}
	// create admin
	if err := CreateGroupUser(g.GroupID, admin); nil != err {
		return err
	}
	return nil
}

func GetGroupsByUID(uid string) ([]*Group, error) {
	groupIDs, err := GetGroupIDsByUID(uid)
	if err != nil {
		return nil, err
	}
	groups := make([]*Group, 0)
	for _, groupID := range groupIDs {
		group, err := GetGroupByGroupID(groupID)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}
	return groups, nil
}

func GetGroupByGroupID(groupID string) (*Group, error) {
	mongo := db.GetLastMongoClient()
	filter := bson.M{"groupID": groupID}
	group := &Group{}
	if err := mongo.FindOne(MongoCollectionGroup, group, filter); nil != err {
		return nil, err
	}
	return group, nil
}

func GetGroupDatasByUID(uid string) ([]*GroupData, error) {
	groupDatas := make([]*GroupData, 0)
	groupIDs, err := GetGroupIDsByUID(uid)
	if err != nil {
		return nil, err
	}
	for _, groupID := range groupIDs {
		groupData, err := GetGroupDataByGroupID(groupID)
		if err != nil {
			return nil, err
		}
		groupDatas = append(groupDatas, groupData)
	}
	return groupDatas, nil
}

func GetGroupDataByGroupID(groupID string) (*GroupData, error) {
	groupData := &GroupData{}
	group, err := GetGroupByGroupID(groupID)
	if nil != err {
		return nil, err
	}
	groupData.Group = *group
	users, err := GetUsersByGroup(group)
	if err != nil {
		return nil, err
	}
	groupData.Members = users
	messages, err := GetGroupMessageWithPage(group, 0, DefaultFriendPageSize)
	if err != nil {
		return nil, err
	}
	groupData.Messages = messages
	return groupData, nil
}

func FindGroupsByGroupName(groupName string) ([]*Group, error) {
	mongo := db.GetLastMongoClient()
	filter := bson.M{
		"groupName": bson.M{
			"$regex": primitive.Regex{Pattern: ".*" + groupName + ".*", Options: "i"},
		},
	}
	groups := make([]*Group, 0)
	err := mongo.Find(MongoCollectionGroup, &groups, filter)
	if err != nil {
		logger.Debug("Mongo Error error: %v", err)
		return nil, err
	}
	return groups, nil
}
