package unlike_post

import (
	common_data "escalateservice/internal/common/data"
	model "escalateservice/internal/model/domain"
	event_model "escalateservice/internal/model/event"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=handler.go -destination=test/mock/handler.go

type Service interface {
	RemoveLikePost(data *model.LikePost)
}

type UserUnlikedPostEventHandler struct {
	service Service
}

func NewUserUnlikedPostEventHandler(service Service) *UserUnlikedPostEventHandler {
	return &UserUnlikedPostEventHandler{
		service: service,
	}
}

func (handler *UserUnlikedPostEventHandler) Handle(event []byte) {
	var userUnlikedPostEvent event_model.UserUnlikedPostEvent
	log.Info().Msg("Handling UnlikeWasRegisteredEvent")

	err := common_data.DeserializeData(event, &userUnlikedPostEvent)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Invalid event data")
		return
	}

	data := mapData(userUnlikedPostEvent)
	handler.service.RemoveLikePost(data)
}

func mapData(event event_model.UserUnlikedPostEvent) *model.LikePost {
	return &model.LikePost{
		PostId:   event.PostId,
		Username: event.Username,
	}
}
