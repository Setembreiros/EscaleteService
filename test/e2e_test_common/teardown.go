package e2e_test_common

import (
	e2e_test_arrange "escalateservice/test/e2e_test_common/arrange"

	"github.com/IBM/sarama"
	"github.com/rs/zerolog/log"
)

func TearDownE2E() {
	db := e2e_test_arrange.CreateTestDatabase()
	db.Client.Clean()
	deleteAllKafkaTopics()
	app.Shutdown()
}

func deleteAllKafkaTopics() error {
	brokers := []string{"localhost:9093"}
	admin, err := sarama.NewClusterAdmin(brokers, sarama.NewConfig())
	if err != nil {
		log.Error().Err(err).Msgf("failed to create cluster admin: %w", err)
		return err
	}
	defer admin.Close()

	topics, err := admin.ListTopics()
	if err != nil {
		log.Error().Err(err).Msgf("failed to list topics: %w", err)
		return err
	}

	for topic := range topics {
		// Podes engadir un filtro para NON borrar __consumer_offsets, etc.
		if topic == "__consumer_offsets" {
			continue
		}
		log.Info().Msgf("Deleting topic: %s", topic)
		if err := admin.DeleteTopic(topic); err != nil {
			log.Warn().Err(err).Msgf("Failed to delete topic: %s", topic)
		}
	}

	return nil
}
