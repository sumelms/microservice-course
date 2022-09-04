package course

import (
	"github.com/ThreeDotsLabs/watermill"
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"

	"github.com/go-kit/log"

	"github.com/ThreeDotsLabs/watermill-amqp/v2/pkg/amqp"

	"github.com/sumelms/microservice-course/internal/course/database"
	"github.com/sumelms/microservice-course/internal/course/domain"
	"github.com/sumelms/microservice-course/internal/course/transport"
)

func NewService(db *sqlx.DB, logger log.Logger) (*domain.Service, error) {
	course, err := database.NewCourseRepository(db)
	if err != nil {
		return nil, err
	}
	subscription, err := database.NewSubscriptionRepository(db)
	if err != nil {
		return nil, err
	}

	// @TODO TEST ONLY - move to adapter
	amqpConfig := amqp.NewDurableQueueConfig("amqp://sumelms:mypassword@nuc:5672/")

	subscriber, err := amqp.NewSubscriber(
		amqpConfig,
		watermill.NewStdLogger(false, false),
	)
	if err != nil {
		panic(err)
	}

	publisher, err := amqp.NewPublisher(amqpConfig, watermill.NewStdLogger(false, false))
	if err != nil {
		panic(err)
	}

	service, err := domain.NewService(
		domain.WithLogger(logger),
		domain.WithCourseRepository(course),
		domain.WithSubscriptionRepository(subscription),
		domain.WithPublisher(publisher),
		domain.WithSubscriber(subscriber))
	if err != nil {
		return nil, err
	}
	return service, nil
}

func NewHTTPService(router *mux.Router, service domain.ServiceInterface, logger log.Logger) error {
	transport.NewHTTPHandler(router, service, logger)
	return nil
}
