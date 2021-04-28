package api

import (
	"crypto/sha1"
	"fmt"
	"framework/db"
	"framework/logger"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

//GenerateToken 根据uid和时间戳生成token
func GenerateToken(uid string) string {
	ts := time.Now().Unix()
	origin := fmt.Sprintf("%v_%v", uid, ts)
	return EncryptBySha1(origin)
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

func CheckToken(token string) (*User, error) {
	rds := db.GetLastRedisClient()
	mongo := db.GetLastMongoClient()

	uid, err := rds.Get(TokenToUIDFormat(token))

	if err != nil {
		logger.Info("redis get uid from token err: %v", err)
		return nil, err
	}

	logger.Debug("mongo filter uid: %v", uid)

	filter := bson.M{"uid": uid}
	user := &User{}
	err = mongo.FindOne("user", user, filter)
	if err != nil {
		logger.Info("mongo get user from uid err: %v, uid: %v", err, uid)
		return nil, err
	}

	return user, nil
}

func EncryptBySha1(plain string) string {
	h := sha1.New()
	h.Write([]byte(plain))

	return fmt.Sprintf("%X", h.Sum(nil))
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
