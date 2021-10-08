package result

import (
	"context"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const CollectionName = "executions"

// NewMongoRespository creates new result repository with db setup for given collection
// use empty collection name as param for default "executions" collection name
func NewMongoRespository(db *mongo.Database, collection string) *MongoRepository {
	if collection == "" {
		collection = CollectionName
	}

	return &MongoRepository{
		Coll: db.Collection(CollectionName),
	}
}

type MongoRepository struct {
	Coll *mongo.Collection
}

func (r *MongoRepository) Get(ctx context.Context, id string) (result testkube.Execution, err error) {
	err = r.Coll.FindOne(ctx, bson.M{"id": id}).Decode(&result)
	return
}

func (r *MongoRepository) Insert(ctx context.Context, result testkube.Execution) (err error) {
	_, err = r.Coll.InsertOne(ctx, result)
	return
}

func (r *MongoRepository) Update(ctx context.Context, result testkube.Execution) (err error) {
	_, err = r.Coll.ReplaceOne(ctx, bson.M{"id": result.Id}, result)
	return
}

func (r *MongoRepository) UpdateResult(ctx context.Context, id string, result testkube.ExecutionResult) (err error) {
	_, err = r.Coll.UpdateOne(ctx, bson.M{"id": id}, bson.M{"$set": bson.M{"executionresult": result}})
	return
}

func (r *MongoRepository) QueuePull(ctx context.Context) (result testkube.Execution, err error) {
	returnDocument := options.After
	err = r.Coll.FindOneAndUpdate(
		ctx,
		bson.M{"executionresult.status": testkube.QUEUED_ExecutionStatus},
		bson.M{"$set": bson.M{"executionresult.status": testkube.PENDING_ExecutionStatus}},
		&options.FindOneAndUpdateOptions{ReturnDocument: &returnDocument},
	).Decode(&result)
	return
}
