package db

import (
	"github.com/rogercoll/ipgeo"
	"go.mongodb.org/mongo-driver/mongo"
)

type Client interface {
	Unseen() (*[]IPColletion, error)
	Store(*[]ipgeo.IPStack) (int, error)
	DbType() string
}

type MongoClient struct {
	client         *mongo.Client
	AtlasAPI       string
	DbName         string
	FromCollection string
	ToCollection   string
}

type IPColletion struct {
	Ip   string
	Seen bool
}

func GetDbType(c Client) string {
	return c.DbType()
}
