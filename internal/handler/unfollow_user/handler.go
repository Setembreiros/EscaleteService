package unfollow_user

import (
	common_data "escalateservice/internal/common/data"
	model "escalateservice/internal/model/domain"
	event_model "escalateservice/internal/model/event"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=handler.go -destination=test/mock/handler.go

type Service interface {
	RemoveFollow(data *model.Follow)
}

type UserAUnfollowedUserBEventHandler struct {
	service Service
}

func NewUserAUnfollowedUserBEventHandler(service Service) *UserAUnfollowedUserBEventHandler {
	return &UserAUnfollowedUserBEventHandler{
		service: service,
	}
}

func (handler *UserAUnfollowedUserBEventHandler) Handle(event []byte) {
	var userAUnfollowedUserBEvent event_model.UserAUnfollowedUserBEvent
	log.Info().Msg("Handling LikeWasRegisteredEvent")

	err := common_data.DeserializeData(event, &userAUnfollowedUserBEvent)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Invalid event data")
		return
	}

	data := mapData(userAUnfollowedUserBEvent)
	handler.service.RemoveFollow(data)
}

func mapData(event event_model.UserAUnfollowedUserBEvent) *model.Follow {
	return &model.Follow{
		Follower: event.Follower,
		Followee: event.Followee,
	}
}
