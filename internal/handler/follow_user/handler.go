package follow_user

import (
	common_data "escalateservice/internal/common/data"
	model "escalateservice/internal/model/domain"
	event_model "escalateservice/internal/model/event"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=handler.go -destination=test/mock/handler.go

type Service interface {
	AddFollow(data *model.Follow)
}

type UserAFollowedUserBEventHandler struct {
	service Service
}

func NewUserAFollowedUserBEventHandler(service Service) *UserAFollowedUserBEventHandler {
	return &UserAFollowedUserBEventHandler{
		service: service,
	}
}

func (handler *UserAFollowedUserBEventHandler) Handle(event []byte) {
	var userAFollowedUserBEvent event_model.UserAFollowedUserBEvent
	log.Info().Msg("Handling LikeWasRegisteredEvent")

	err := common_data.DeserializeData(event, &userAFollowedUserBEvent)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Invalid event data")
		return
	}

	data := mapData(userAFollowedUserBEvent)
	handler.service.AddFollow(data)
}

func mapData(event event_model.UserAFollowedUserBEvent) *model.Follow {
	return &model.Follow{
		Follower: event.Follower,
		Followee: event.Followee,
	}
}
