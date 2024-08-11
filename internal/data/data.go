package data

import (
	"context"
	"kratos-poc/internal/conf"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
  database *mongo.Database
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
  ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
  defer cancel()
  client, err := mongo.Connect(ctx, options.Client().ApplyURI(c.Database.Source))

	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
    if err = client.Disconnect(ctx); err != nil {
        panic(err)
    }
	}
	return &Data{
    database: client.Database("kratos-poc"),
  }, cleanup, nil
}
