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
