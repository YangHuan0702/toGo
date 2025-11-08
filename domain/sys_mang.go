package domain

import "gorm.io/datatypes"

type Menu struct {
	Id        int64          `gorm:"primary_key;auto_increment" json:"id"`
	Name      string         `gorm:"name" json:name`
	Path      string         `gorm:"path" json:path`
	Component string         `gorm:"component" json:component`
	ParentId  int64          `gorm:"parent_id" json:"parent_id"`
	Meta      datatypes.JSON `gorm:"type:jsonb;column:meta" json:meta`
}

func (Menu) TableName() string {
	return "tobe_menu"
}
