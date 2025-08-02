package provider

import (
	"escalateservice/infrastructure/database/migrator"
	"escalateservice/infrastructure/database/sql_db"
	"escalateservice/infrastructure/kafka"
	"escalateservice/internal/api"
	"escalateservice/internal/bus"
	database "escalateservice/internal/db"
	"escalateservice/internal/handler/post_created"
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
	return []api.Controller{}
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
	}
}

func (p *Provider) ProvideKafkaConsumer(eventBus *bus.EventBus) (*kafka.KafkaConsumer, error) {
	brokers := p.kafkaBrokers()

	return kafka.NewKafkaConsumer(brokers, eventBus)
}

func (p *Provider) kafkaBrokers() []string {
	if p.env == "development" {
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
