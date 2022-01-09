package productmodel

import (
	"gshop/sdk/sdkcm"
)

type ListFilter struct {
	Category *string `json:"category" form:"category""`
}

type ListParam struct {
	sdkcm.Paging `json:",inline"`
	*ListFilter  `json:",inline"`
}
