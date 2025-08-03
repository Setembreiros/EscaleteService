package like_post

import (
	common_data "escalateservice/internal/common/data"
	model "escalateservice/internal/model/domain"
	event_model "escalateservice/internal/model/event"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=handler.go -destination=test/mock/handler.go

type Service interface {
	AddLikePost(data *model.LikePost)
}

type UserLikedPostEventHandler struct {
	service Service
}

func NewUserLikedPostEventHandler(service Service) *UserLikedPostEventHandler {
	return &UserLikedPostEventHandler{
		service: service,
	}
}

func (handler *UserLikedPostEventHandler) Handle(event []byte) {
	var userLikedPostEvent event_model.UserLikedPostEvent
	log.Info().Msg("Handling LikeWasRegisteredEvent")

	err := common_data.DeserializeData(event, &userLikedPostEvent)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Invalid event data")
		return
	}

	data := mapData(userLikedPostEvent)
	handler.service.AddLikePost(data)
}

func mapData(event event_model.UserLikedPostEvent) *model.LikePost {
	return &model.LikePost{
		PostId:   event.PostId,
		Username: event.Username,
	}
}
