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

func (p productRepo) GetProduct(ctx context.Context, id uint32) (*productmodel.Product, error) {
	var product productmodel.Product

	if err := p.DB.Where("id = ? ", id).First(&product).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, sdkcm.ErrDataNotFound
		}
		return nil, sdkcm.ErrDB(err)
	}

	return &product, nil
}

func (p productRepo) ListProduct(ctx context.Context, filter *productmodel.ListFilter, paging *sdkcm.Paging, moreKeys ...string) ([]productmodel.Product, error) {
	var data []productmodel.Product

	db := p.DB.Table(productmodel.Product{}.TableName())

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

func NewProductRepo(DB *gorm.DB) *productRepo {
	return &productRepo{DB: DB}
}
