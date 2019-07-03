package dao

import (
	"errors"

	"github.com/xtech-cloud/omo-msa-app/model"
)

func CreateApp(_appname string, _appKey string, _appSecret string, _publicKey string, _privateKey string) error {
	db, err := model.OpenDB()
	if nil != err {
		return err
	}
	defer model.CloseDB(db)

	application := model.Application{
		AppName:    _appname,
		AppKey:     _appKey,
		AppSecret:  _appSecret,
		PublicKey:  _publicKey,
		PrivateKey: _privateKey,
	}

	isBlank := db.NewRecord(application)
	if !isBlank {
		return errors.New("application is exists!")
	}
	return db.Create(&application).Error
}

func QueryApp(_appname string) (*model.Application, error) {
	db, err := model.OpenDB()
	if nil != err {
		return nil, err
	}
	defer model.CloseDB(db)

	application := model.Application{}

	err = db.Where("app_name = ?", _appname).First(&application).Error
	if nil != err {
		return nil, err
	}

	return &application, nil
}

func FindApp(_appkey string, _appsecret string) (*model.Application, error) {
	db, err := model.OpenDB()
	if nil != err {
		return nil, err
	}
	defer model.CloseDB(db)

	application := model.Application{}

	err = db.Where("app_key= ? AND app_secret = ?", _appkey, _appsecret).First(&application).Error
	if nil != err {
		return nil, err
	}

	return &application, nil
}

func ListApp() ([]*model.Application, error) {
	db, err := model.OpenDB()
	if nil != err {
		return nil, err
	}
	defer model.CloseDB(db)

	applications := make([]*model.Application, 0)

	err = db.Find(&applications).Error
	if nil != err {
		return nil, err
	}

	return applications, nil
}

func UpdateAppSecret(_appname string, _secret string) error {
	db, err := model.OpenDB()
	if nil != err {
		return err
	}
	defer model.CloseDB(db)

	err = db.Model(&model.Application{}).Where("app_name= ?", _appname).Update("app_secret", _secret).Error
	return err
}

func UpdateAppKey(_appname string, _publicKey string, _privateKey string) error {
	db, err := model.OpenDB()
	if nil != err {
		return err
	}
	defer model.CloseDB(db)

	fields := map[string]interface{}{
		"key_public":  _publicKey,
		"key_private": _privateKey,
	}

	err = db.Model(&model.Application{}).Where("app_name= ?", _appname).Update(fields).Error
	return err
}

func UpdateAppProfile(_appname string, _profile string) error {
	db, err := model.OpenDB()
	if nil != err {
		return err
	}
	defer model.CloseDB(db)

	err = db.Model(&model.Application{}).Where("app_name= ?", _appname).Update("profile", _profile).Error
	return err
}
