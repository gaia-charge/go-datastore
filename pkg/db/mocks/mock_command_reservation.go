package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type CommandReservationMockData struct {
	CommandReservation db.CommandReservation
	Error              error
}

type CommandReservationsMockData struct {
	CommandReservations []db.CommandReservation
	Error               error
}

func (r *MockRepositoryService) CreateCommandReservation(ctx context.Context, arg db.CreateCommandReservationParams) (db.CommandReservation, error) {
	r.createCommandReservationMockData = append(r.createCommandReservationMockData, arg)
	return db.CommandReservation{}, nil
}

func (r *MockRepositoryService) GetCommandReservation(ctx context.Context, id int64) (db.CommandReservation, error) {
	if len(r.getCommandReservationMockData) == 0 {
		return db.CommandReservation{}, ErrorNotFound()
	}

	response := r.getCommandReservationMockData[0]
	r.getCommandReservationMockData = r.getCommandReservationMockData[1:]
	return response.CommandReservation, response.Error
}

func (r *MockRepositoryService) UpdateCommandReservation(ctx context.Context, arg db.UpdateCommandReservationParams) (db.CommandReservation, error) {
	r.updateCommandReservationMockData = append(r.updateCommandReservationMockData, arg)
	return db.CommandReservation{}, nil
}

func (r *MockRepositoryService) GetCreateCommandReservationMockData() (db.CreateCommandReservationParams, error) {
	if len(r.createCommandReservationMockData) == 0 {
		return db.CreateCommandReservationParams{}, ErrorNotFound()
	}

	response := r.createCommandReservationMockData[0]
	r.createCommandReservationMockData = r.createCommandReservationMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateCommandReservationMockData() (db.UpdateCommandReservationParams, error) {
	if len(r.updateCommandReservationMockData) == 0 {
		return db.UpdateCommandReservationParams{}, ErrorNotFound()
	}

	response := r.updateCommandReservationMockData[0]
	r.updateCommandReservationMockData = r.updateCommandReservationMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetCommandReservationMockData(response CommandReservationMockData) {
	r.getCommandReservationMockData = append(r.getCommandReservationMockData, response)
}
