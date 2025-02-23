package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type CommandUnlockMockData struct {
	CommandUnlock db.CommandUnlock
	Error         error
}

type CommandUnlocksMockData struct {
	CommandUnlocks []db.CommandUnlock
	Error          error
}

func (r *MockRepositoryService) CreateCommandUnlock(ctx context.Context, arg db.CreateCommandUnlockParams) (db.CommandUnlock, error) {
	r.createCommandUnlockMockData = append(r.createCommandUnlockMockData, arg)
	return db.CommandUnlock{}, nil
}

func (r *MockRepositoryService) GetCommandUnlock(ctx context.Context, id int64) (db.CommandUnlock, error) {
	if len(r.getCommandUnlockMockData) == 0 {
		return db.CommandUnlock{}, ErrorNotFound()
	}

	response := r.getCommandUnlockMockData[0]
	r.getCommandUnlockMockData = r.getCommandUnlockMockData[1:]
	return response.CommandUnlock, response.Error
}

func (r *MockRepositoryService) UpdateCommandUnlock(ctx context.Context, arg db.UpdateCommandUnlockParams) (db.CommandUnlock, error) {
	r.updateCommandUnlockMockData = append(r.updateCommandUnlockMockData, arg)
	return db.CommandUnlock{}, nil
}

func (r *MockRepositoryService) GetCreateCommandUnlockMockData() (db.CreateCommandUnlockParams, error) {
	if len(r.createCommandUnlockMockData) == 0 {
		return db.CreateCommandUnlockParams{}, ErrorNotFound()
	}

	response := r.createCommandUnlockMockData[0]
	r.createCommandUnlockMockData = r.createCommandUnlockMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateCommandUnlockMockData() (db.UpdateCommandUnlockParams, error) {
	if len(r.updateCommandUnlockMockData) == 0 {
		return db.UpdateCommandUnlockParams{}, ErrorNotFound()
	}

	response := r.updateCommandUnlockMockData[0]
	r.updateCommandUnlockMockData = r.updateCommandUnlockMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetCommandUnlockMockData(response CommandUnlockMockData) {
	r.getCommandUnlockMockData = append(r.getCommandUnlockMockData, response)
}
