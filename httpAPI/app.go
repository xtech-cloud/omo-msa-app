package httpAPI

import (
	"github.com/xtech-cloud/omo-msa-app/dao"
	"github.com/xtech-cloud/omo-msa-app/processor"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type CreateAppRequest struct {
	AppName string `json:"appname" binding:"required"`
}

func HandleCreateApp(_context *gin.Context) {
	var req CreateAppRequest
	err := _context.ShouldBindJSON(&req)
	if nil != err {
		renderBadError(_context, err)
		return
	}
	appkey, appsecret, publicKey, privateKey, err := processor.CreateApp(req.AppName)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	err = dao.CreateApp(req.AppName, appkey, appsecret, publicKey, privateKey)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	renderOK(_context, gin.H{})
}

type QueryAppRequest struct {
	AppName string `json:"appname" binding:"required"`
}

func HandleQueryApp(_context *gin.Context) {
	var req QueryAppRequest
	err := _context.ShouldBindJSON(&req)
	if nil != err {
		renderBadError(_context, err)
		return
	}

	application, err := dao.QueryApp(req.AppName)
	if nil != err {
		renderModuleError(_context, err)
		return
	}
	renderOK(_context, gin.H{
		"AppName":    application.AppName,
		"AppKey":     application.AppKey,
		"AppSecret":  application.AppSecret,
		"PublicKey":  application.PublicKey,
		"PrivateKey": application.PrivateKey,
		"Profile":    application.Profile,
		"CreatedAt":  application.GModel.CreatedAt.String(),
	})
}

func HandleListApp(_context *gin.Context) {
	applications, err := dao.ListApp()
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	names := make([]string, len(applications))
	for i := 0; i < len(applications); i++ {
		names[i] = applications[i].AppName
	}
	renderOK(_context, gin.H{
		"Applications": names,
	})
}

type ResetSecretRequest struct {
	AppName string `json:"appname" binding:"required"`
}

func HandleResetSecret(_context *gin.Context) {
	var req ResetSecretRequest
	err := _context.ShouldBindJSON(&req)
	if nil != err {
		renderBadError(_context, err)
		return
	}

	secret := processor.ResetSecret(req.AppName)

	err = dao.UpdateAppSecret(req.AppName, secret)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	renderOK(_context, gin.H{})
}

type ResetKeyRequest struct {
	AppName string `json:"appname" binding:"required"`
}

func HandleResetKey(_context *gin.Context) {
	var req ResetKeyRequest
	err := _context.ShouldBindJSON(&req)
	if nil != err {
		renderBadError(_context, err)
		return
	}

	publicKey, privateKey, err := processor.ResetKey(req.AppName)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	err = dao.UpdateAppKey(req.AppName, publicKey, privateKey)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	renderOK(_context, gin.H{})
}

type ModifyAppProfileRequest struct {
	AppName string `json:"appname" binding:"required"`
	Profile string `json:"profile"`
}

func HandleModifyAppProfile(_context *gin.Context) {
	var req ModifyAppProfileRequest
	err := _context.ShouldBindJSON(&req)
	if nil != err {
		renderBadError(_context, err)
		return
	}

	err = dao.UpdateAppProfile(req.AppName, req.Profile)
	if nil != err {
		renderModuleError(_context, err)
		return
	}

	renderOK(_context, gin.H{})
}
