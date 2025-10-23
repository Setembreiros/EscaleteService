package e2e_test_arrange

import (
	"escalateservice/cmd/startup"
	"escalateservice/test/e2e_test_common"
	"testing"
)

func PublishEvent(t *testing.T, eventName string, data any) {
	provider := startup.NewProvider(e2e_test_common.Env, e2e_test_common.ConnStr)

	eventBus, err := provider.ProvideEventBus()
	if err != nil {
		t.Fatalf("failed to marshal event: %v", err)
	}
	err = eventBus.Publish(eventName, data)
	if err != nil {
		t.Fatalf("failed to marshal event: %v", err)
	}
}
