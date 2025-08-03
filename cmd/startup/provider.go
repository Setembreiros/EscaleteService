package startup

import (
	"escalateservice/infrastructure/database/migrator"
	"escalateservice/infrastructure/database/sql_db"
	"escalateservice/infrastructure/kafka"
	"escalateservice/internal/api"
	"escalateservice/internal/bus"
	database "escalateservice/internal/db"
	"escalateservice/internal/handler/like_post"
	"escalateservice/internal/handler/post_created"
	"escalateservice/internal/handler/review_created"
	"escalateservice/internal/handler/superlike_post"
	"escalateservice/internal/handler/unlike_post"
	"escalateservice/internal/handler/unsuperlike_post"
	"escalateservice/internal/handler/user_created"
	"escalateservice/internal/model/event"
)

type Provider struct {
	env     string
	connStr string
}

func NewProvider(env, connStr string) *Provider {
	return &Provider{
		env:     env,
		connStr: connStr,
	}
}

func (p *Provider) ProvideGooseCLient() (*migrator.GooseClient, error) {
	return migrator.NewGooseClient(p.connStr)
}

func (p *Provider) ProvideDb() (*sql_db.SqlDatabase, error) {
	return sql_db.NewDatabase(p.connStr)
}

func (p *Provider) ProvideEventBus() (*bus.EventBus, error) {
	kafkaProducer, err := kafka.NewKafkaProducer(p.kafkaBrokers())
	if err != nil {
		return nil, err
	}

	return bus.NewEventBus(kafkaProducer), nil
}

func (p *Provider) ProvideApiEndpoint(sqlClient *sql_db.SqlDatabase, bus *bus.EventBus) *api.Api {
	return api.NewApiEndpoint(p.env, p.ProvideApiControllers(sqlClient, bus))
}

func (p *Provider) ProvideApiControllers(sqlClient *sql_db.SqlDatabase, bus *bus.EventBus) []api.Controller {
	return []api.Controller{
		api.NewPingController(),
	}
}

func (p *Provider) ProvideSubscriptions(sqlClient *sql_db.SqlDatabase) *[]bus.EventSubscription {
	return &[]bus.EventSubscription{
		{
			EventType: event.UserWasRegisteredEventName,
			Handler:   user_created.NewUserWasRegisteredEventHandler(user_created.NewUserCreatedService(user_created.NewUserCreatedRepository(database.NewDatabase(sqlClient)))),
		},
		{
			EventType: event.PostWasCreatedEventName,
			Handler:   post_created.NewPostWasCreatedEventHandler(post_created.NewPostCreatedService(post_created.NewPostCreatedRepository(database.NewDatabase(sqlClient)))),
		},
		{
			EventType: event.ReviewWasCreatedEventName,
			Handler:   review_created.NewReviewWasCreatedEventHandler(review_created.NewReviewCreatedService(review_created.NewReviewCreatedRepository(database.NewDatabase(sqlClient)))),
		},
		{
			EventType: event.UserLikedPostEventName,
			Handler:   like_post.NewUserLikedPostEventHandler(like_post.NewLikePostService(like_post.NewLikePostRepository(database.NewDatabase(sqlClient)))),
		},
		{
			EventType: event.UserUnlikedPostEventName,
			Handler:   unlike_post.NewUserUnlikedPostEventHandler(unlike_post.NewUnlikePostService(unlike_post.NewUnlikePostRepository(database.NewDatabase(sqlClient)))),
		},
		{
			EventType: event.UserSuperlikedPostEventName,
			Handler:   superlike_post.NewUserSuperlikedPostEventHandler(superlike_post.NewSuperlikePostService(superlike_post.NewSuperlikePostRepository(database.NewDatabase(sqlClient)))),
		},
		{
			EventType: event.UserUnsuperlikedPostEventName,
			Handler:   unsuperlike_post.NewUserUnsuperlikedPostEventHandler(unsuperlike_post.NewUnsuperlikePostService(unsuperlike_post.NewUnsuperlikePostRepository(database.NewDatabase(sqlClient)))),
		},
	}
}

func (p *Provider) ProvideKafkaConsumer(eventBus *bus.EventBus) (*kafka.KafkaConsumer, error) {
	brokers := p.kafkaBrokers()

	return kafka.NewKafkaConsumer(brokers, eventBus)
}

func (p *Provider) kafkaBrokers() []string {
	if p.env == "development" || p.env == "test" {
		return []string{
			"localhost:9093",
		}
	} else {
		return []string{
			"172.31.0.242:9092",
			"172.31.7.110:9092",
		}
	}
}
