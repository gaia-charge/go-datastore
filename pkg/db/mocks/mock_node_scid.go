package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type NodeScidMockData struct {
	NodeScid db.NodeScid
	Error    error
}

type NodeScidsMockData struct {
	NodeScids []db.NodeScid
	Error     error
}

func (r *MockRepositoryService) CountNodeScids(ctx context.Context, nodeID int64) (int64, error) {
	if len(r.countNodeScidsMockData) == 0 {
		return 0, nil
	}

	response := r.countNodeScidsMockData[0]
	r.countNodeScidsMockData = r.countNodeScidsMockData[1:]
	return response.Count, response.Error
}

func (r *MockRepositoryService) CreateNodeScid(ctx context.Context, arg db.CreateNodeScidParams) (db.NodeScid, error) {
	r.createNodeScidMockData = append(r.createNodeScidMockData, arg)
	return db.NodeScid{}, nil
}

func (r *MockRepositoryService) DeleteNodeScid(ctx context.Context, id int64) error {
	r.deleteNodeScidMockData = append(r.deleteNodeScidMockData, id)
	return nil
}

func (r *MockRepositoryService) GetNodeScid(ctx context.Context, id int64) (db.NodeScid, error) {
	if len(r.getNodeScidMockData) == 0 {
		return db.NodeScid{}, ErrorNotFound()
	}

	response := r.getNodeScidMockData[0]
	r.getNodeScidMockData = r.getNodeScidMockData[1:]
	return response.NodeScid, response.Error
}

func (r *MockRepositoryService) SetCountNodeScidsMockData(response CountMockData) {
	r.countNodeScidsMockData = append(r.countNodeScidsMockData, response)
}

func (r *MockRepositoryService) GetCreateNodeScidMockData() (db.CreateNodeScidParams, error) {
	if len(r.createNodeScidMockData) == 0 {
		return db.CreateNodeScidParams{}, ErrorNotFound()
	}

	response := r.createNodeScidMockData[0]
	r.createNodeScidMockData = r.createNodeScidMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) GetDeleteNodeScidMockData() (int64, error) {
	if len(r.deleteNodeScidMockData) == 0 {
		return 0, ErrorNotFound()
	}

	response := r.deleteNodeScidMockData[0]
	r.deleteNodeScidMockData = r.deleteNodeScidMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetNodeScidMockData(response NodeScidMockData) {
	r.getNodeScidMockData = append(r.getNodeScidMockData, response)
}
