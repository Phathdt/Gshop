package productrepo

import (
	"context"

	"gorm.io/gorm"
	"gshop/module/products/productmodel"
	"gshop/sdk/sdkcm"
)

type productRepo struct {
	DB *gorm.DB
}

func (p productRepo) ListProduct(ctx context.Context, filter *productmodel.ListFilter, paging *sdkcm.Paging) ([]productmodel.Product, error) {
	var data []productmodel.Product

	db := p.DB.Table(productmodel.Product{}.TableName())

	if f := filter; f != nil {
		if v := f.Category; v != nil {
			db = db.Where("category = ?", v)
		}
	}

	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, sdkcm.ErrDB(err)
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

func NewProductRepo(DB *gorm.DB) *productRepo {
	return &productRepo{DB: DB}
}
