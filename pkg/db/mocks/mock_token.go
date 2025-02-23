package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type CountMockData struct {
	Count int64
	Error error
}

type TokenMockData struct {
	Token db.Token
	Error error
}

type TokensMockData struct {
	Tokens []db.Token
	Error  error
}

func (r *MockRepositoryService) CountTokens(ctx context.Context, arg db.CountTokensParams) (int64, error) {
	if len(r.countTokensMockData) == 0 {
		return 0, nil
	}

	response := r.countTokensMockData[0]
	r.countTokensMockData = r.countTokensMockData[1:]
	return response.Count, response.Error
}

func (r *MockRepositoryService) CreateToken(ctx context.Context, arg db.CreateTokenParams) (db.Token, error) {
	r.createTokenMockData = append(r.createTokenMockData, arg)
	return db.Token{}, nil
}

func (r *MockRepositoryService) DeleteTokenByUid(ctx context.Context, uid string) error {
	r.deleteTokenByUidMockData = append(r.deleteTokenByUidMockData, uid)
	return nil
}

func (r *MockRepositoryService) GetToken(ctx context.Context, id int64) (db.Token, error) {
	if len(r.getTokenMockData) == 0 {
		return db.Token{}, ErrorNotFound()
	}

	response := r.getTokenMockData[0]
	r.getTokenMockData = r.getTokenMockData[1:]
	return response.Token, response.Error
}

func (r *MockRepositoryService) GetTokenByAuthID(ctx context.Context, authID string) (db.Token, error) {
	if len(r.getTokenByAuthIDMockData) == 0 {
		return db.Token{}, ErrorNotFound()
	}

	response := r.getTokenByAuthIDMockData[0]
	r.getTokenByAuthIDMockData = r.getTokenByAuthIDMockData[1:]
	return response.Token, response.Error
}

func (r *MockRepositoryService) GetTokenByUid(ctx context.Context, uid string) (db.Token, error) {
	if len(r.getTokenByUidMockData) == 0 {
		return db.Token{}, ErrorNotFound()
	}

	response := r.getTokenByUidMockData[0]
	r.getTokenByUidMockData = r.getTokenByUidMockData[1:]
	return response.Token, response.Error
}

func (r *MockRepositoryService) GetTokenByUserID(ctx context.Context, arg db.GetTokenByUserIDParams) (db.Token, error) {
	if len(r.getTokenByUserIDMockData) == 0 {
		return db.Token{}, ErrorNotFound()
	}

	response := r.getTokenByUserIDMockData[0]
	r.getTokenByUserIDMockData = r.getTokenByUserIDMockData[1:]
	return response.Token, response.Error
}

func (r *MockRepositoryService) ListTokens(ctx context.Context, arg db.ListTokensParams) ([]db.Token, error) {
	if len(r.listTokensMockData) == 0 {
		return []db.Token{}, nil
	}

	response := r.listTokensMockData[0]
	r.listTokensMockData = r.listTokensMockData[1:]
	return response.Tokens, response.Error
}

func (r *MockRepositoryService) ListRfidTokensByUserID(ctx context.Context, userID int64) ([]db.Token, error) {
	if len(r.listRfidTokensByUserIDMockData) == 0 {
		return []db.Token{}, nil
	}

	response := r.listRfidTokensByUserIDMockData[0]
	r.listRfidTokensByUserIDMockData = r.listRfidTokensByUserIDMockData[1:]
	return response.Tokens, response.Error
}

func (r *MockRepositoryService) ListTokensByUserID(ctx context.Context, userID int64) ([]db.Token, error) {
	if len(r.listTokensByUserIDMockData) == 0 {
		return []db.Token{}, nil
	}

	response := r.listTokensByUserIDMockData[0]
	r.listTokensByUserIDMockData = r.listTokensByUserIDMockData[1:]
	return response.Tokens, response.Error
}

func (r *MockRepositoryService) UpdateTokenByUid(ctx context.Context, arg db.UpdateTokenByUidParams) (db.Token, error) {
	r.updateTokenByUidMockData = append(r.updateTokenByUidMockData, arg)
	return db.Token{}, nil
}

func (r *MockRepositoryService) SetCountTokensMockData(response CountMockData) {
	r.countTokensMockData = append(r.countTokensMockData, response)
}

func (r *MockRepositoryService) GetCreateTokenMockData() (db.CreateTokenParams, error) {
	if len(r.createTokenMockData) == 0 {
		return db.CreateTokenParams{}, ErrorNotFound()
	}

	response := r.createTokenMockData[0]
	r.createTokenMockData = r.createTokenMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteTokenByUidMockData() (string, error) {
	if len(r.deleteTokenByUidMockData) == 0 {
		return "", ErrorNotFound()
	}

	response := r.deleteTokenByUidMockData[0]
	r.deleteTokenByUidMockData = r.deleteTokenByUidMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetUpdateTokenByUidMockData() (db.UpdateTokenByUidParams, error) {
	if len(r.updateTokenByUidMockData) == 0 {
		return db.UpdateTokenByUidParams{}, ErrorNotFound()
	}

	response := r.updateTokenByUidMockData[0]
	r.updateTokenByUidMockData = r.updateTokenByUidMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetTokenMockData(response TokenMockData) {
	r.getTokenMockData = append(r.getTokenMockData, response)
}

func (r *MockRepositoryService) SetGetTokenByAuthIDMockData(response TokenMockData) {
	r.getTokenByAuthIDMockData = append(r.getTokenByAuthIDMockData, response)
}

func (r *MockRepositoryService) SetGetTokenByUidMockData(response TokenMockData) {
	r.getTokenByUidMockData = append(r.getTokenByUidMockData, response)
}

func (r *MockRepositoryService) SetGetTokenByUserIDMockData(response TokenMockData) {
	r.getTokenByUidMockData = append(r.getTokenByUserIDMockData, response)
}

func (r *MockRepositoryService) SetListTokensMockData(response TokensMockData) {
	r.listTokensMockData = append(r.listTokensMockData, response)
}

func (r *MockRepositoryService) SetListRfidTokensByUserIDMockData(response TokensMockData) {
	r.listRfidTokensByUserIDMockData = append(r.listRfidTokensByUserIDMockData, response)
}

func (r *MockRepositoryService) SetListTokensByUserIDMockData(response TokensMockData) {
	r.listTokensByUserIDMockData = append(r.listTokensByUserIDMockData, response)
}
