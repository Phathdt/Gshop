package productstorage

import (
	"context"

	"gorm.io/gorm"
	"gshop/module/products/productmodel"
	"gshop/pkg/sdkcm"
)

func (s *productSQLStorage) GetProductByCondition(ctx context.Context, cond map[string]interface{}) (*productmodel.Product, error) {
	var product productmodel.Product

	if err := s.db.Where(cond).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, sdkcm.ErrDataNotFound
		}
		return nil, sdkcm.ErrDB(err)
	}

	return &product, nil
}
