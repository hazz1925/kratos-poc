package data

import (
	"context"
	"time"

	"kratos-poc/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
  collection *mongo.Collection
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
    collection: data.database.Collection("greeter"),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
  g.CreatedAt = time.Now()
  g.UpdatedAt = time.Now()
  res, err := r.collection.InsertOne(ctx, g)
  if err != nil {
    log.Fatal(err)
  }
  g.ID = res.InsertedID.(primitive.ObjectID)
  log.Info(res.InsertedID)
	return g, nil
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(context.Context, int64) (*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(context.Context, string) ([]*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(context.Context) ([]*biz.Greeter, error) {
	return nil, nil
}
