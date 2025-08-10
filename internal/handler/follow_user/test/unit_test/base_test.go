package follow_user_test

import (
	"bytes"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog/log"
)

var ctrl *gomock.Controller
var loggerOutput bytes.Buffer

func setUp(t *testing.T) {
	ctrl = gomock.NewController(t)
	log.Logger = log.Output(&loggerOutput)
}
