package utils

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

var ErrCoinNotEnough = errors.New("coin not enough")

func (u *User) GetCoin() uint {
	return u.DailyCoin + u.HonorCoin
}

func (u *User) CostCoin(db *gorm.DB, amount uint) error {

	if amount > u.DailyCoin+u.HonorCoin {
		return ErrCoinNotEnough
	}

	if u.DailyCoin < amount {
		u.HonorCoin -= amount - u.DailyCoin
		u.DailyCoin = 0
	} else {
		u.DailyCoin -= amount
	}

	if err := db.Where(u).Select("daily_coin", "honor_coin").Updates(u).Error; err != nil {
		return fmt.Errorf("failed to update coins: %w", err)
	}

	return nil
}

func (u *User) HasAnyRole(roles ...RoleID) bool {
	if u == nil {
		return false
	}

	roleSet := make(map[RoleID]struct{}, len(roles))
	for _, role := range roles {
		roleSet[role] = struct{}{}
	}

	for _, userRole := range u.UserRoles {
		if _, ok := roleSet[userRole.Role.ID]; ok {
			return true
		}
	}
	return false
}
