package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type EmailSubscriptionMockData struct {
	EmailSubscription db.EmailSubscription
	Error             error
}

func (r *MockRepositoryService) CreateEmailSubscription(ctx context.Context, arg db.CreateEmailSubscriptionParams) (db.EmailSubscription, error) {
	r.createEmailSubscriptionMockData = append(r.createEmailSubscriptionMockData, arg)
	return db.EmailSubscription{}, nil
}

func (r *MockRepositoryService) GetEmailSubscriptionByEmail(ctx context.Context, email string) (db.EmailSubscription, error) {
	if len(r.getEmailSubscriptionByEmailMockData) == 0 {
		return db.EmailSubscription{}, ErrorNotFound()
	}

	response := r.getEmailSubscriptionByEmailMockData[0]
	r.getEmailSubscriptionByEmailMockData = r.getEmailSubscriptionByEmailMockData[1:]
	return response.EmailSubscription, response.Error
}

func (r *MockRepositoryService) UpdateEmailSubscription(ctx context.Context, arg db.UpdateEmailSubscriptionParams) (db.EmailSubscription, error) {
	r.updateEmailSubscriptionMockData = append(r.updateEmailSubscriptionMockData, arg)
	return db.EmailSubscription{}, nil
}

func (r *MockRepositoryService) GetCreateEmailSubscriptionMockData() (db.CreateEmailSubscriptionParams, error) {
	if len(r.createEmailSubscriptionMockData) == 0 {
		return db.CreateEmailSubscriptionParams{}, ErrorNotFound()
	}

	response := r.createEmailSubscriptionMockData[0]
	r.createEmailSubscriptionMockData = r.createEmailSubscriptionMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetEmailSubscriptionByEmailMockData(response EmailSubscriptionMockData) {
	r.getEmailSubscriptionByEmailMockData = append(r.getEmailSubscriptionByEmailMockData, response)
}

func (r *MockRepositoryService) GetUpdateEmailSubscriptionMockData() (db.UpdateEmailSubscriptionParams, error) {
	if len(r.updateEmailSubscriptionMockData) == 0 {
		return db.UpdateEmailSubscriptionParams{}, ErrorNotFound()
	}

	response := r.updateEmailSubscriptionMockData[0]
	r.updateEmailSubscriptionMockData = r.updateEmailSubscriptionMockData[1:]
	return response, nil
}
