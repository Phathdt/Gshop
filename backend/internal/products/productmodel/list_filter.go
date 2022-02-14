package productmodel

import (
	"gshop/pkg/sdkcm"
)

type ListFilter struct {
	CategoryId *uint `query:"category_id" json:"category_id"`
}

type ListParam struct {
	sdkcm.Paging `json:",inline"`
	*ListFilter  `json:",inline"`
}
