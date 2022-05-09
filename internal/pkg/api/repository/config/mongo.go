package config

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/kubeshop/testkube/pkg/analytics"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
)

const CollectionName = "config"
const Id = "api"

func NewMongoRespository(db *mongo.Database) *MongoRepository {
	return &MongoRepository{
		Coll: db.Collection(CollectionName),
	}
}

type MongoRepository struct {
	Coll *mongo.Collection
}

func (r *MongoRepository) GetUniqueClusterId(ctx context.Context) (clusterId string, err error) {
	config := testkube.Config{}
	_ = r.Coll.FindOne(ctx, bson.M{"id": Id}).Decode(&config)

	// generate new cluster Id and save if there is not already
	if config.ClusterId == "" {
		config.ClusterId = fmt.Sprintf("cluster%s", analytics.MachineID())
		err := r.Upsert(ctx, config)
		return config.ClusterId, err
	}

	return config.ClusterId, nil
}

func (r *MongoRepository) Get(ctx context.Context) (result testkube.Config, err error) {
	err = r.Coll.FindOne(ctx, bson.M{"id": Id}).Decode(&result)
	return
}

func (r *MongoRepository) Upsert(ctx context.Context, result testkube.Config) (err error) {
	upsert := true
	result.Id = Id
	_, err = r.Coll.ReplaceOne(ctx, bson.M{"id": Id}, result, &options.ReplaceOptions{Upsert: &upsert})
	return
}
