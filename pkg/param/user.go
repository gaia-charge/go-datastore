package param

import "github.com/gaia-charge/go-datastore/pkg/db"

func NewUpdateUserParams(user db.User) db.UpdateUserParams {
	return db.UpdateUserParams{
		ID:                user.ID,
		CommissionPercent: user.CommissionPercent,
		DeviceToken:       user.DeviceToken,
		LinkingPubkey:     user.LinkingPubkey,
		NodeID:            user.NodeID,
		Pubkey:            user.Pubkey,
		IsRestricted:      user.IsRestricted,
		ReferralCode:      user.ReferralCode,
		CircuitUserID:     user.CircuitUserID,
		Name:              user.Name,
		Address:           user.Address,
		PostalCode:        user.PostalCode,
		City:              user.City,
		BatteryCapacity:   user.BatteryCapacity,
		BatteryPowerAc:    user.BatteryPowerAc,
		BatteryPowerDc:    user.BatteryPowerDc,
	}
}
