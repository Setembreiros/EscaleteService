package review_created

import (
	common_data "escalateservice/internal/common/data"
	model "escalateservice/internal/model/domain"
	event_model "escalateservice/internal/model/event"

	"github.com/rs/zerolog/log"
)

//go:generate mockgen -source=handler.go -destination=test/mock/handler.go

type Service interface {
	AddReview(data *model.Review)
}

type ReviewWasCreatedEventHandler struct {
	service Service
}

func NewReviewWasCreatedEventHandler(service Service) *ReviewWasCreatedEventHandler {
	return &ReviewWasCreatedEventHandler{
		service: service,
	}
}

func (handler *ReviewWasCreatedEventHandler) Handle(event []byte) {
	var reviewWasCreatedEvent event_model.ReviewWasCreatedEvent
	log.Info().Msg("Handling ReviewWasRegisteredEvent")

	err := common_data.DeserializeData(event, &reviewWasCreatedEvent)
	if err != nil {
		log.Error().Stack().Err(err).Msg("Invalid event data")
		return
	}

	data := mapData(reviewWasCreatedEvent)
	handler.service.AddReview(data)
}

func mapData(event event_model.ReviewWasCreatedEvent) *model.Review {
	return &model.Review{
		ReviewId: event.ReviewId,
		PostId:   event.PostId,
		Reviewer: event.Username,
		Rating:   event.Rating,
	}
}
