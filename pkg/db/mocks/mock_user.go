package mocks

import (
	"context"
	"database/sql"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type UserMockData struct {
	User  db.User
	Error error
}

func (r *MockRepositoryService) CreateUser(ctx context.Context, arg db.CreateUserParams) (db.User, error) {
	r.createUserMockData = append(r.createUserMockData, arg)
	return db.User{}, nil
}

func (r *MockRepositoryService) GetUser(ctx context.Context, id int64) (db.User, error) {
	if len(r.getUserMockData) == 0 {
		return db.User{}, ErrorNotFound()
	}

	response := r.getUserMockData[0]
	r.getUserMockData = r.getUserMockData[1:]
	return response.User, response.Error
}

func (r *MockRepositoryService) GetUserByDeviceToken(ctx context.Context, deviceToken sql.NullString) (db.User, error) {
	if len(r.getUserByDeviceTokenMockData) == 0 {
		return db.User{}, ErrorNotFound()
	}

	response := r.getUserByDeviceTokenMockData[0]
	r.getUserByDeviceTokenMockData = r.getUserByDeviceTokenMockData[1:]
	return response.User, response.Error
}

func (r *MockRepositoryService) GetUserByLinkingPubkey(ctx context.Context, linkingKey string) (db.User, error) {
	if len(r.getUserByLinkingPubkeyMockData) == 0 {
		return db.User{}, ErrorNotFound()
	}

	response := r.getUserByLinkingPubkeyMockData[0]
	r.getUserByLinkingPubkeyMockData = r.getUserByLinkingPubkeyMockData[1:]
	return response.User, response.Error
}

func (r *MockRepositoryService) GetUserByPubkey(ctx context.Context, pubkey string) (db.User, error) {
	if len(r.getUserByPubkeyMockData) == 0 {
		return db.User{}, ErrorNotFound()
	}

	response := r.getUserByPubkeyMockData[0]
	r.getUserByPubkeyMockData = r.getUserByPubkeyMockData[1:]
	return response.User, response.Error
}

func (r *MockRepositoryService) GetUserByReferralCode(ctx context.Context, referralCode sql.NullString) (db.User, error) {
	if len(r.getUserByReferralCodeMockData) == 0 {
		return db.User{}, ErrorNotFound()
	}

	response := r.getUserByReferralCodeMockData[0]
	r.getUserByReferralCodeMockData = r.getUserByReferralCodeMockData[1:]
	return response.User, response.Error
}

func (r *MockRepositoryService) GetUserBySessionID(ctx context.Context, id int64) (db.User, error) {
	if len(r.getUserBySessionIDMockData) == 0 {
		return db.User{}, ErrorNotFound()
	}

	response := r.getUserBySessionIDMockData[0]
	r.getUserBySessionIDMockData = r.getUserBySessionIDMockData[1:]
	return response.User, response.Error
}

func (r *MockRepositoryService) GetUserByTokenID(ctx context.Context, id int64) (db.User, error) {
	if len(r.getUserByTokenIDMockData) == 0 {
		return db.User{}, ErrorNotFound()
	}

	response := r.getUserByTokenIDMockData[0]
	r.getUserByTokenIDMockData = r.getUserByTokenIDMockData[1:]
	return response.User, response.Error
}

func (r *MockRepositoryService) UpdateUser(ctx context.Context, arg db.UpdateUserParams) (db.User, error) {
	r.updateUserMockData = append(r.updateUserMockData, arg)
	return db.User{}, nil
}

func (r *MockRepositoryService) UpdateUserByPubkey(ctx context.Context, arg db.UpdateUserByPubkeyParams) (db.User, error) {
	r.updateUserByPubkeyMockData = append(r.updateUserByPubkeyMockData, arg)
	return db.User{}, nil
}

func (r *MockRepositoryService) GetCreateUserMockData() (db.CreateUserParams, error) {
	if len(r.createUserMockData) == 0 {
		return db.CreateUserParams{}, ErrorNotFound()
	}

	response := r.createUserMockData[0]
	r.createUserMockData = r.createUserMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetUserMockData(response UserMockData) {
	r.getUserMockData = append(r.getUserMockData, response)
}

func (r *MockRepositoryService) SetGetUserByDeviceTokenMockData(response UserMockData) {
	r.getUserByDeviceTokenMockData = append(r.getUserByDeviceTokenMockData, response)
}

func (r *MockRepositoryService) SetGetUserByLinkingPubkeyMockData(response UserMockData) {
	r.getUserByLinkingPubkeyMockData = append(r.getUserByLinkingPubkeyMockData, response)
}

func (r *MockRepositoryService) SetGetUserByPubkeyMockData(response UserMockData) {
	r.getUserByPubkeyMockData = append(r.getUserByPubkeyMockData, response)
}

func (r *MockRepositoryService) SetGetUserByReferralCodeMockData(response UserMockData) {
	r.getUserByReferralCodeMockData = append(r.getUserByReferralCodeMockData, response)
}

func (r *MockRepositoryService) SetGetUserBySessionIDMockData(response UserMockData) {
	r.getUserBySessionIDMockData = append(r.getUserBySessionIDMockData, response)
}

func (r *MockRepositoryService) SetGetUserByTokenIDMockData(response UserMockData) {
	r.getUserByTokenIDMockData = append(r.getUserByTokenIDMockData, response)
}

func (r *MockRepositoryService) GetUpdateUserMockData() (db.UpdateUserParams, error) {
	if len(r.updateUserMockData) == 0 {
		return db.UpdateUserParams{}, ErrorNotFound()
	}

	response := r.updateUserMockData[0]
	r.updateUserMockData = r.updateUserMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateUserByPubkeyMockData() (db.UpdateUserByPubkeyParams, error) {
	if len(r.updateUserByPubkeyMockData) == 0 {
		return db.UpdateUserByPubkeyParams{}, ErrorNotFound()
	}

	response := r.updateUserByPubkeyMockData[0]
	r.updateUserByPubkeyMockData = r.updateUserByPubkeyMockData[1:]
	return response, nil
}
