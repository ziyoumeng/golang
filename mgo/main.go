package main

import (
	"strings"
	"time"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	. "gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2/txn"
)

type MgoDialInfo struct {
	MongodbAddr      string
	MongodbUser      string
	MongodbPassword  string
	MongodbPoolLimit int
}

func main() {
	dialInfo := &mgo.DialInfo{
		Addrs:     strings.Split("192.168.202.125:5000", ","),
		Direct:    false,
		Timeout:   time.Second * 3,
		PoolLimit: 10,
		Username:  "",
		Password:  "",
	}
	session, err := mgo.DialWithInfo(dialInfo)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	tcollection := session.DB("test").C("mytest1")
	runner := txn.NewRunner(tcollection)
	ops := []txn.Op{{
		C:      "accounts",
		Id:     "aram",
		Assert: M{"balance": M{"$gte": 100}},
		Update: M{"$inc": M{"balance": -100}},
	}, {
		C:      "accounts",
		Id:     "ben",
		Assert: M{"valid": true},
		Update: M{"$inc": M{"balance": 100}},
	}}
	id := bson.NewObjectId() // Optional
	err = runner.Run(ops, id, nil)
	if err != nil {
		panic(err)
	}

}
