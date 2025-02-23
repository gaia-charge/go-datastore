package channelrequest

import (
	"context"

	"github.com/gaia-charge/go-datastore/pkg/db"
)

type ChannelRequestRepository interface {
	CreateChannelRequest(ctx context.Context, arg db.CreateChannelRequestParams) (db.ChannelRequest, error)
	CreateChannelRequestHtlc(ctx context.Context, arg db.CreateChannelRequestHtlcParams) (db.ChannelRequestHtlc, error)
	GetChannelRequest(ctx context.Context, id int64) (db.ChannelRequest, error)
	GetChannelRequestByChannelPoint(ctx context.Context, arg db.GetChannelRequestByChannelPointParams) (db.ChannelRequest, error)
	GetChannelRequestByPaymentHash(ctx context.Context, paymentHash []byte) (db.ChannelRequest, error)
	GetChannelRequestByPendingChanId(ctx context.Context, pendingChanId []byte) (db.ChannelRequest, error)
	GetChannelRequestHtlc(ctx context.Context, channelRequestID int64) (db.ChannelRequestHtlc, error)
	GetChannelRequestHtlcByCircuitKey(ctx context.Context, arg db.GetChannelRequestHtlcByCircuitKeyParams) (db.ChannelRequestHtlc, error)
	ListChannelRequestHtlcs(ctx context.Context, channelRequestID int64) ([]db.ChannelRequestHtlc, error)
	UpdateChannelRequest(ctx context.Context, arg db.UpdateChannelRequestParams) (db.ChannelRequest, error)
	UpdateChannelRequestHtlcByCircuitKey(ctx context.Context, arg db.UpdateChannelRequestHtlcByCircuitKeyParams) (db.ChannelRequestHtlc, error)
	UpdatePendingChannelRequestByPubkey(ctx context.Context, arg db.UpdatePendingChannelRequestByPubkeyParams) (db.ChannelRequest, error)
}

func NewRepository(repositoryService *db.RepositoryService) ChannelRequestRepository {
	return ChannelRequestRepository(repositoryService)
}
