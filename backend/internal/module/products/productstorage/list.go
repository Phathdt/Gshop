package productstorage

import (
	"context"

	"gshop/internal/module/products/productmodel"
	"gshop/pkg/sdkcm"
)

func (s *productSQLStorage) ListProduct(ctx context.Context, filter *productmodel.ListFilter, paging *sdkcm.Paging, moreKeys ...string) ([]productmodel.Product, error) {
	var data []productmodel.Product

	db := s.db.Table(productmodel.Product{}.TableName())

	if f := filter; f != nil {
		if v := f.CategoryId; v != nil {
			db = db.Where("category_id = ?", v)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, sdkcm.ErrDB(err)
	}

	if len(moreKeys) > 0 {
		for _, k := range moreKeys {
			db = db.Preload(k)
		}
	}

	db = db.Limit(paging.Limit).Order("id desc")

	if paging.Cursor == 0 {
		db = db.Offset((paging.Page - 1) * paging.Limit)
	} else {
		db = db.Where("id > ?", paging.Cursor)
	}

	if err := db.Find(&data).Error; err != nil {
		return nil, sdkcm.ErrDB(err)
	}

	paging.HasNext = len(data) >= paging.Limit

	return data, nil
}
