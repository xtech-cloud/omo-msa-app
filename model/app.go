package model

import "github.com/jinzhu/gorm"

type Application struct {
	GModel     gorm.Model `gorm:"embedded"`
	AppName    string     `gorm:"column:app_name;type:varchar(32);unique_index;not null"`
	AppKey     string     `gorm:"column:app_key;type:char(32);unique_index;not null"`
	AppSecret  string     `gorm:"column:app_secret;type:char(32);unique_index;not null"`
	PublicKey  string     `gorm:"column:key_public;type:text;unique_index;not null"`
	PrivateKey string     `gorm:"column:key_private;type:text;unique_index;not null"`
	Profile    string     `gorm:"column:profile;type:text"`
}

func (Application) TableName() string {
	return "app_application"
}
