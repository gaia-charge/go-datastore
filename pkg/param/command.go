package param

import "github.com/gaia-charge/go-datastore/pkg/db"

func NewUpdateCommandReservationParams(command db.CommandReservation) db.UpdateCommandReservationParams {
	return db.UpdateCommandReservationParams{
		ID:          command.ID,
		Status:      command.Status,
		ExpiryDate:  command.ExpiryDate,
		EvseUid:     command.EvseUid,
		LastUpdated: command.LastUpdated,
	}
}

func NewUpdateCommandStartParams(command db.CommandStart) db.UpdateCommandStartParams {
	return db.UpdateCommandStartParams{
		ID:          command.ID,
		Status:      command.Status,
		LastUpdated: command.LastUpdated,
	}
}

func NewUpdateCommandStopParams(command db.CommandStop) db.UpdateCommandStopParams {
	return db.UpdateCommandStopParams{
		ID:          command.ID,
		Status:      command.Status,
		LastUpdated: command.LastUpdated,
	}
}

func NewUpdateCommandUnlockParams(command db.CommandUnlock) db.UpdateCommandUnlockParams {
	return db.UpdateCommandUnlockParams{
		ID:          command.ID,
		Status:      command.Status,
		LastUpdated: command.LastUpdated,
	}
}
