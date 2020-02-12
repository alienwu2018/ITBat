package downlink

import (
	"ITBat/model"
	"ITBat/pkg/logging"
)

type DownLink struct {
	model.Model

	BId  int    `gorm:"column:bid" json:"b_id"`
	PUrl string `gorm:"column:purl" json:"p_url"`
	PWD  string `gorm:"column:pwd" json:"pwd"`
}

func (DownLink) TableName() string {
	return "downlink"
}

func QueryByBId(BId int) (dl DownLink, err error) {
	if err = model.DB.Debug().Model(DownLink{}).Where("bid = ?", BId).First(&dl).Error; err != nil {
		logging.DebugLog(err)
		return
	}
	return
}
