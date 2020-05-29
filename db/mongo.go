package db

import (
	"context"
	"time"

	"github.com/rogercoll/ipgeo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMongoClient(atlasAPI, dbName, fColl, tColl string) (*MongoClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(
		atlasAPI,
	))
	if err != nil {
		return nil, err
	}

	return &MongoClient{client: client, AtlasAPI: atlasAPI,
		DbName:         dbName,
		FromCollection: fColl,
		ToCollection:   tColl}, nil
}

func (m MongoClient) Unseen() (*[]IPColletion, error) {

	var unseenIPs []IPColletion
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	mdb := m.client.Database(m.DbName)
	mip := mdb.Collection(m.FromCollection)
	//filter bson.D{{"coin", coin}}
	cursor, err := mip.Find(ctx, bson.D{{"seen", false}})
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &unseenIPs); err != nil {
		return nil, err
	}
	return &unseenIPs, nil
}

//We must update the other db too
func (m MongoClient) Store(ipsInfo *[]ipgeo.IPStack) (int, error) {
	for i, ipInfo := range *ipsInfo {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()
		mdb := m.client.Database(m.DbName)
		mipInf := mdb.Collection(m.ToCollection)
		_, err := mipInf.InsertOne(ctx, ipInfo)
		if err != nil {
			return i, err
		}
		mseen := mdb.Collection(m.FromCollection)
		_, err = mseen.UpdateOne(ctx,
			bson.D{{"ip", ipInfo.Ip}},
			bson.D{
				{"$set", bson.D{{"seen", true}}},
			},
		)
		if err != nil {
			return i + 1, err
		}
	}
	return len(*ipsInfo), nil
}

func (m MongoClient) DbType() string {
	return "MongoDB"
}
