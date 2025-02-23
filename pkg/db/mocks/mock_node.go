package mocks

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type NodeMockData struct {
	Node  db.Node
	Error error
}

type NodesMockData struct {
	Nodes []db.Node
	Error error
}

func (r *MockRepositoryService) CreateNode(ctx context.Context, arg db.CreateNodeParams) (db.Node, error) {
	r.createNodeMockData = append(r.createNodeMockData, arg)
	return db.Node{}, nil
}

func (r *MockRepositoryService) GetNode(ctx context.Context, id int64) (db.Node, error) {
	if len(r.getNodeMockData) == 0 {
		return db.Node{}, ErrorNotFound()
	}

	response := r.getNodeMockData[0]
	r.getNodeMockData = r.getNodeMockData[1:]
	return response.Node, response.Error
}

func (r *MockRepositoryService) GetNodeByPubkey(ctx context.Context, pubkey string) (db.Node, error) {
	if len(r.getNodeByPubkeyMockData) == 0 {
		return db.Node{}, ErrorNotFound()
	}

	response := r.getNodeByPubkeyMockData[0]
	r.getNodeByPubkeyMockData = r.getNodeByPubkeyMockData[1:]
	return response.Node, response.Error
}

func (r *MockRepositoryService) GetNodeByUserID(ctx context.Context, id int64) (db.Node, error) {
	if len(r.getNodeByUserIDMockData) == 0 {
		return db.Node{}, ErrorNotFound()
	}

	response := r.getNodeByUserIDMockData[0]
	r.getNodeByUserIDMockData = r.getNodeByUserIDMockData[1:]
	return response.Node, response.Error
}

func (r *MockRepositoryService) ListNodes(ctx context.Context) ([]db.Node, error) {
	if len(r.listNodesMockData) == 0 {
		return []db.Node{}, nil
	}

	response := r.listNodesMockData[0]
	r.listNodesMockData = r.listNodesMockData[1:]
	return response.Nodes, response.Error
}

func (r *MockRepositoryService) ListActiveNodes(ctx context.Context, isLsp bool) ([]db.Node, error) {
	if len(r.listActiveNodesMockData) == 0 {
		return []db.Node{}, nil
	}

	response := r.listActiveNodesMockData[0]
	r.listActiveNodesMockData = r.listActiveNodesMockData[1:]
	return response.Nodes, response.Error
}

func (r *MockRepositoryService) UpdateNode(ctx context.Context, arg db.UpdateNodeParams) (db.Node, error) {
	r.updateNodeMockData = append(r.updateNodeMockData, arg)
	return db.Node{}, nil
}

func (r *MockRepositoryService) GetCreateNodeMockData() (db.CreateNodeParams, error) {
	if len(r.createNodeMockData) == 0 {
		return db.CreateNodeParams{}, ErrorNotFound()
	}

	response := r.createNodeMockData[0]
	r.createNodeMockData = r.createNodeMockData[1:]
	return response, nil
}

func (r *MockRepositoryService) SetGetNodeMockData(response NodeMockData) {
	r.getNodeMockData = append(r.getNodeMockData, response)
}

func (r *MockRepositoryService) SetGetNodeByPubkeyMockData(response NodeMockData) {
	r.getNodeMockData = append(r.getNodeByPubkeyMockData, response)
}

func (r *MockRepositoryService) SetGetNodeByUserIDMockData(response NodeMockData) {
	r.getNodeMockData = append(r.getNodeByUserIDMockData, response)
}

func (r *MockRepositoryService) SetListNodesMockData(response NodesMockData) {
	r.listNodesMockData = append(r.listNodesMockData, response)
}

func (r *MockRepositoryService) SetListActiveNodesMockData(response NodesMockData) {
	r.listActiveNodesMockData = append(r.listActiveNodesMockData, response)
}

func (r *MockRepositoryService) GetUpdateNodeMockData() (db.UpdateNodeParams, error) {
	if len(r.updateNodeMockData) == 0 {
		return db.UpdateNodeParams{}, ErrorNotFound()
	}

	response := r.updateNodeMockData[0]
	r.updateNodeMockData = r.updateNodeMockData[1:]
	return response, nil
}
