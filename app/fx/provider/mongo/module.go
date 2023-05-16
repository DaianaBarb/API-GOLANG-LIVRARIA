package mongo

import (
	dao "api/app/provider/mongodb/dao"
	"context"
	"fmt"
	"sync"

	"github.com/americanas-go/ignite/go.mongodb.org/mongo-driver.v1"
	health "github.com/americanas-go/ignite/go.mongodb.org/mongo-driver.v1/plugins/contrib/americanas-go/health.v1"
	mongoDriver "go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"
)

func NewLivroRepositoryModule() fx.Option {
	return fx.Options(
		mongoModule(),
		fx.Provide(
			dao.NewOptions,
			dao.NewLivroRepository,
		),
	)
}

var once sync.Once

func mongoModule() fx.Option {
	options := fx.Options()
	once.Do(func() {
		options = fx.Options(
			fx.Provide(
				health.NewHealth,
			),
			fx.Provide(
				func(healthPlugin *health.Health) []mongo.Plugin {
					return []mongo.Plugin{
						healthPlugin.Register,
					}
				},
			),
			fx.Provide(

				func(ctx context.Context, plugins []mongo.Plugin) (*mongoDriver.Database, error) {

					connection, err := mongo.NewConn(ctx, plugins...)
					if err != nil {
						return nil, err
					}
					fmt.Println("DB successfully checked and authorized")
					return connection.Database, nil
				},
			),
		)
	})
	return options

}
