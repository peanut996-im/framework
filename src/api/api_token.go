// Package api
// @Title  api_token.go
// @Description  提供对token的常用操作
// @Author  peanut996
// @Update  peanut996  2021/5/22 1:43
package api

import (
	"fmt"
	"framework/api/model"
	"framework/db"
	"framework/encoding"
	"framework/logger"
	"time"
)

const (
	//UID_TO_TOKEN_FORMAT redis key name for value: token
	UID_TO_TOKEN_FORMAT = "%v_to_token"
	//TOKEN_TO_UID_FORMAT redis key name for value: uid
	TOKEN_TO_UID_FORMAT = "%v_to_uid"
	//DEFAULT_TOKEN_EXPIRE_TIME Token default expiration time
	DEFAULT_TOKEN_EXPIRE_TIME = 6 * 60 * 60
)

func CheckToken(token string) (*model.User, error) {
	rds := db.GetLastRedisClient()

	uid, err := rds.Get(TokenToUIDFormat(token))

	if err != nil {
		logger.Info("redis get uid from token err: %v", err)
		return nil, err
	}

	logger.Debug("mongo filter uid: %v", uid)
	return model.GetUserByUID(uid)
}

//InsertToken token插入数据库，若已存在则直接返回已存在的token
func InsertToken(uid string) (string, error) {
	// 先查询 已存在则直接获取数据库token 并重新设置过期时间
	rds := db.GetLastRedisClient()
	tokenKey := UIDToTokenFormat(uid)

	token, err := rds.Get(tokenKey)
	if err != nil {
		if db.IsNotExistError(err) {
			token = GenerateToken(uid)
		} else {
			// redis 有问题
			logger.Info("redis get token err", err)
			return "", err
		}
	}

	// redis uid => token
	_, err = rds.Set(tokenKey, token, DEFAULT_TOKEN_EXPIRE_TIME)
	if nil != err {
		logger.Info("redis set uid => token err: %v", err)
		return "", err
	}
	uidKey := TokenToUIDFormat(token)
	// redis token => uid
	_, err = rds.Set(uidKey, uid, DEFAULT_TOKEN_EXPIRE_TIME)
	if nil != err {
		logger.Info("redis set token => uid err: %v", err)
		return "", err
	}
	return token, nil
}

//GenerateToken 根据uid和时间戳生成token
func GenerateToken(uid string) string {
	ts := time.Now().Unix()
	origin := fmt.Sprintf("%v_%v", uid, ts)
	return encoding.EncryptBySha1(origin)
}

func DeleteToken(token string) error {
	rds := db.GetLastRedisClient()
	// 只需要删除token => uid 即可使token失效 uid=>token可复用
	_, err := rds.DelOne(TokenToUIDFormat(token))

	if nil != err {
		return err
	}
	return nil
}

func UIDToTokenFormat(uid string) string {
	return fmt.Sprintf(UID_TO_TOKEN_FORMAT, uid)
}

func TokenToUIDFormat(token string) string {
	return fmt.Sprintf(TOKEN_TO_UID_FORMAT, token)
}
