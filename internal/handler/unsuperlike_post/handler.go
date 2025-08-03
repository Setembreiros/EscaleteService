package unsuperlike_post

import (
	common_data "escalateservice/internal/common/data"
	model "escalateservice/internal/model/domain"
	event_model "escalateservice/internal/model/event"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=handler.go -destination=test/mock/handler.go

type Service interface {
	RemoveSuperlikePost(data *model.SuperlikePost)
}

type UserUnsuperlikedPostEventHandler struct {
	service Service
}

func NewUserUnsuperlikedPostEventHandler(service Service) *UserUnsuperlikedPostEventHandler {
	return &UserUnsuperlikedPostEventHandler{
		service: service,
	}
}

func (handler *UserUnsuperlikedPostEventHandler) Handle(event []byte) {
	var userUnsuperlikedPostEvent event_model.UserUnsuperlikedPostEvent
	log.Info().Msg("Handling UnsuperlikeWasRegisteredEvent")

	err := common_data.DeserializeData(event, &userUnsuperlikedPostEvent)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Invalid event data")
		return
	}

	data := mapData(userUnsuperlikedPostEvent)
	handler.service.RemoveSuperlikePost(data)
}

func mapData(event event_model.UserUnsuperlikedPostEvent) *model.SuperlikePost {
	return &model.SuperlikePost{
		PostId:   event.PostId,
		Username: event.Username,
	}
}
