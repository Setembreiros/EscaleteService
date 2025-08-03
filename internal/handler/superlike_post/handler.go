package superlike_post

import (
	common_data "escalateservice/internal/common/data"
	model "escalateservice/internal/model/domain"
	event_model "escalateservice/internal/model/event"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=handler.go -destination=test/mock/handler.go

type Service interface {
	AddSuperlikePost(data *model.SuperlikePost)
}

type UserSuperlikedPostEventHandler struct {
	service Service
}

func NewUserSuperlikedPostEventHandler(service Service) *UserSuperlikedPostEventHandler {
	return &UserSuperlikedPostEventHandler{
		service: service,
	}
}

func (handler *UserSuperlikedPostEventHandler) Handle(event []byte) {
	var userSuperlikedPostEvent event_model.UserSuperlikedPostEvent
	log.Info().Msg("Handling SuperlikeWasRegisteredEvent")

	err := common_data.DeserializeData(event, &userSuperlikedPostEvent)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Invalid event data")
		return
	}

	data := mapData(userSuperlikedPostEvent)
	handler.service.AddSuperlikePost(data)
}

func mapData(event event_model.UserSuperlikedPostEvent) *model.SuperlikePost {
	return &model.SuperlikePost{
		PostId:   event.PostId,
		Username: event.Username,
	}
}
