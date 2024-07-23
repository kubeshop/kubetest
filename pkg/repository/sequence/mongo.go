package sequence

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var _ Repository = (*MongoRepository)(nil)

const (
	CollectionSequences = "sequences"
)

type ExecutionType string

const (
	ExecutionTypeTest         ExecutionType = "t"
	ExecutionTypeTestSuite    ExecutionType = "ts"
	ExecutionTypeTestWorkflow ExecutionType = "tw"
)

func NewMongoRepository(db *mongo.Database, opts ...Opt) *MongoRepository {
	r := &MongoRepository{
		Coll: db.Collection(CollectionSequences),
	}

	for _, opt := range opts {
		opt(r)
	}

	return r
}

type Opt func(*MongoRepository)

func WithMongoRepositoryCollection(collection *mongo.Collection) Opt {
	return func(r *MongoRepository) {
		r.Coll = collection
	}
}

type MongoRepository struct {
	Coll *mongo.Collection
}

type oldExecutionNumber struct {
	Name        string `json:"name"`
	Number      int    `json:"number"`
	IsTestSuite bool   `json:"isTestSuite"`
}

type executionNumber struct {
	Id            string        `bson:"_id"`
	Number        int           `bson:"number"`
	ExecutionType ExecutionType `bson:"executionType"`
}

// GetNextExecutionNumber gets next execution number by name and type
func (r *MongoRepository) GetNextExecutionNumber(ctx context.Context, name string, executionType ExecutionType) (number int32, err error) {
	oldName := getOldName(name, executionType)
	number, err = r.getOldNumber(ctx, oldName)
	if err != nil {
		return 0, err
	}

	id := getMongoId(name, executionType)
	executionNumber := executionNumber{
		Id:            id,
		Number:        int(number),
		ExecutionType: executionType,
	}

	opts := options.FindOneAndUpdate()
	opts.SetUpsert(true)

	err = r.Coll.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$set": executionNumber}, opts).Err()
	if err != nil && !mongo.IsDuplicateKeyError(err) {
		return 0, err
	}

	if number != 0 {
		_, err = r.Coll.DeleteOne(ctx, bson.M{"name": oldName})
		if err != nil {
			return 0, err
		}
	}

	opts.SetUpsert(false)
	opts.SetReturnDocument(options.After)

	err = r.Coll.FindOneAndUpdate(ctx, bson.M{"_id": id}, bson.M{"$inc": bson.M{"number": 1}}, opts).Decode(&executionNumber)
	if err == nil {
		return 0, err
	}

	return int32(executionNumber.Number), nil
}

// DeleteExecutionNumber deletes execution number by name and type
func (r *MongoRepository) DeleteExecutionNumber(ctx context.Context, name string, executionType ExecutionType) (err error) {
	_, err = r.Coll.DeleteOne(ctx, bson.M{"name": getOldName(name, executionType)})
	if err != nil {
		return err
	}

	_, err = r.Coll.DeleteOne(ctx, bson.M{"_id": getMongoId(name, executionType)})
	return err
}

// DeleteExecutionNumbers deletes multiple execution numbers by names and type
func (r *MongoRepository) DeleteExecutionNumbers(ctx context.Context, names []string, executionType ExecutionType) (err error) {
	ids := make([]string, len(names))
	for i := range names {
		ids[i] = getOldName(names[i], executionType)
	}

	_, err = r.Coll.DeleteMany(ctx, bson.M{"name": bson.M{"$in": ids}})
	if err != nil {
		return err
	}

	for i := range names {
		ids[i] = getMongoId(names[i], executionType)
	}

	_, err = r.Coll.DeleteMany(ctx, bson.M{"_id": bson.M{"$in": ids}})
	return err
}

// DeleteAllExecutionNumbers deletes all execution numbers by type
func (r *MongoRepository) DeleteAllExecutionNumbers(ctx context.Context, executionType string) (err error) {
	isTestSuite := false
	if executionType == "testsuite" {
		isTestSuite = true
	}

	_, err = r.Coll.DeleteMany(ctx, bson.M{"istestsuite": isTestSuite})
	if err != nil {
		return err
	}

	_, err = r.Coll.DeleteMany(ctx, bson.M{"executionType": executionType})
	return err
}

func (r *MongoRepository) getOldNumber(ctx context.Context, name string) (int32, error) {
	var executionNumber oldExecutionNumber

	err := r.Coll.FindOne(ctx, bson.M{"name": name}).Decode(&executionNumber)
	if err != nil && err != mongo.ErrNoDocuments {
		return 0, err
	}

	return int32(executionNumber.Number), nil
}

func getMongoId(name string, executionType ExecutionType) string {
	return fmt.Sprintf("%s-%s", name, executionType)
}

func getOldName(name string, executionType ExecutionType) string {
	oldPrefix := ""
	if executionType == "testsuite" {
		oldPrefix = "ts-"
	}

	return fmt.Sprintf("%s%s", oldPrefix, name)
}
