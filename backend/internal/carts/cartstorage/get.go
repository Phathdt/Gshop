package cartstorage

import (
	"context"

	"gorm.io/gorm"
	"gshop/internal/carts/cartmodel"
	"gshop/pkg/sdkcm"
)

func (s *cartSQLStorage) GetCartByCondition(ctx context.Context, cond map[string]interface{}, moreKeys ...string) (*cartmodel.Cart, error) {
	var cart cartmodel.Cart

	db := s.db.Table(cartmodel.Cart{}.TableName())

	if len(moreKeys) > 0 {
		for _, k := range moreKeys {
			db = db.Preload(k)
		}
	}

	if err := db.Where(cond).First(&cart).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, sdkcm.ErrDB(err)
		}
	}

	return &cart, nil
}
