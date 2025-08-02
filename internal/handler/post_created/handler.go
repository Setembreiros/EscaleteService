package post_created

import (
	common_data "escalateservice/internal/common/data"
	model "escalateservice/internal/model/domain"
	event_model "escalateservice/internal/model/event"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=handler.go -destination=test/mock/handler.go

type Service interface {
	AddPost(data *model.Post)
}

type PostWasCreatedEventHandler struct {
	service Service
}

func NewPostWasCreatedEventHandler(service Service) *PostWasCreatedEventHandler {
	return &PostWasCreatedEventHandler{
		service: service,
	}
}

func (handler *PostWasCreatedEventHandler) Handle(event []byte) {
	var postWasCreatedEvent event_model.PostWasCreatedEvent
	log.Info().Msg("Handling PostWasRegisteredEvent")

	err := common_data.DeserializeData(event, &postWasCreatedEvent)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Invalid event data")
		return
	}

	data := mapData(postWasCreatedEvent)
	handler.service.AddPost(data)
}

func mapData(event event_model.PostWasCreatedEvent) *model.Post {
	return &model.Post{
		PostId:   event.PostId,
		Username: event.Metadata.Username,
	}
}
