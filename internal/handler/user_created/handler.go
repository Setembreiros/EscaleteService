package user_created

import (
	common_data "escalateservice/internal/common/data"
	model "escalateservice/internal/model/domain"
	event_model "escalateservice/internal/model/event"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=handler.go -destination=test/mock/handler.go

type Service interface {
	AddUser(data *model.User)
}

type UserWasRegisteredEventHandler struct {
	service Service
}

func NewUserWasRegisteredEventHandler(service Service) *UserWasRegisteredEventHandler {
	return &UserWasRegisteredEventHandler{
		service: service,
	}
}

func (handler *UserWasRegisteredEventHandler) Handle(event []byte) {
	var userWasRegisteredEvent event_model.UserWasRegisteredEvent
	log.Info().Msg("Handling UserWasRegisteredEvent")

	err := common_data.DeserializeData(event, &userWasRegisteredEvent)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Invalid event data")
		return
	}

	data := mapData(userWasRegisteredEvent)
	handler.service.AddUser(data)
}

func mapData(event event_model.UserWasRegisteredEvent) *model.User {
	return &model.User{
		Username: event.Username,
	}
}
