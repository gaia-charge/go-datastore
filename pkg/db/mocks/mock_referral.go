package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type ReferralMockData struct {
	Referral db.Referral
	Error    error
}

func (r *MockRepositoryService) CreateReferral(ctx context.Context, arg db.CreateReferralParams) (db.Referral, error) {
	r.createReferralMockData = append(r.createReferralMockData, arg)
	return db.Referral{}, nil
}

func (r *MockRepositoryService) GetReferralByIpAddress(ctx context.Context, ipAddress string) (db.Referral, error) {
	if len(r.getReferralByIpAddressMockData) == 0 {
		return db.Referral{}, ErrorNotFound()
	}

	response := r.getReferralByIpAddressMockData[0]
	r.getReferralByIpAddressMockData = r.getReferralByIpAddressMockData[1:]
	return response.Referral, response.Error
}

func (r *MockRepositoryService) GetCreateReferralMockData() (db.CreateReferralParams, error) {
	if len(r.createReferralMockData) == 0 {
		return db.CreateReferralParams{}, ErrorNotFound()
	}

	response := r.createReferralMockData[0]
	r.createReferralMockData = r.createReferralMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetReferralByIpAddressMockData(response ReferralMockData) {
	r.getReferralByIpAddressMockData = append(r.getReferralByIpAddressMockData, response)
}
