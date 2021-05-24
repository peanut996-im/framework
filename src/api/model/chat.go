package model

import (
	"framework/db"
	"framework/tool"
)

type ChatMessage struct {
	From    string `json:"from" bson:"from"`
	To      string `json:"to,omitempty" bson:"to"`
	Content string `json:"content" bson:"content"`
	Time    string `json:"time,omitempty" bson:"time"`
	Type    string `json:"type" bson:"type"`
}

func NewChatMessage() *ChatMessage {
	return &ChatMessage{
		Time: tool.GetNowUnixNanoSecond(),
	}
}

func ChatMessageFrom(from, to, content, Type string) *ChatMessage {
	c := NewChatMessage()
	c.From = from
	c.To = to
	c.Content = content
	c.Type = Type
	return c
}

func InsertChatMessage(c *ChatMessage) error {
	mongo := db.GetLastMongoClient()
	if _, err := mongo.InsertOne("Chat", c); err != nil {
		return err
	}
	return nil
}
