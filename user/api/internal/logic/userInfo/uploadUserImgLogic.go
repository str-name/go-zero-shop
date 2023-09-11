package userInfo

import (
	"context"
	"gorm.io/gorm"
	"io"
	"net/http"
	"strings"
	"zero-shop/common/globalKey"
	"zero-shop/common/qiniu"
	"zero-shop/common/tool"
	"zero-shop/common/xerr"
	"zero-shop/user/db/model"

	"zero-shop/user/api/internal/svc"
	"zero-shop/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadUserImgLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUploadUserImgLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UploadUserImgLogic {
	return &UploadUserImgLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UploadUserImgLogic) UploadUserImg(req *types.UploadUserImgReq) error {
	// todo: add your logic here and delete this line

	// 获取图片信息
	form := l.r.MultipartForm
	image := form.File["image"][0]
	// 获取用户
	userID := req.UserID
	var u model.User
	err := l.svcCtx.UserDB.Where("id = ? and del_state = ?", userID, globalKey.DelStateNo).Take(&u).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return xerr.NewErrCode(xerr.USER_NOT_EXISTS_ERROR)
		}
		return xerr.NewErrCode(xerr.DB_ERROR)
	}

	// 获取图片后缀，判断是否为图片类型
	imgName := image.Filename
	nameList := strings.Split(imgName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	if !tool.InList(suffix, tool.WhiteImageList) {
		return xerr.NewErrMsg("请上传图片类型")
	}

	// 读取图片，生成图片hash
	imgObj, err := image.Open()
	if err != nil {
		return xerr.NewErrMsg("打开图片失败")
	}
	byteData, err := io.ReadAll(imgObj)
	if err != nil {
		return xerr.NewErrMsg("读取图片失败")
	}
	imgHash := tool.Md5(byteData)
	// 根据hash判断图片是否存在
	var i model.UserHeaderImage
	err = l.svcCtx.UserDB.Where("hash = ? and del_state = ?", imgHash, globalKey.DelStateNo).Take(&i).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return xerr.NewErrCode(xerr.DB_ERROR)
	}
	if err == nil {
		// 图片存在，更新用户信息
		u.HeaderImageID = i.ID
		err = l.svcCtx.UserDB.Updates(&u).Error
		if err != nil {
			return xerr.NewErrCode(xerr.DB_UPDATE_ZERO_ERROR)
		}
	}

	// 图片保存到七牛云和数据库，同时更新用户头像ID
	err = l.svcCtx.UserDB.Transaction(func(tx *gorm.DB) error {
		// 上传图片
		imgPath, err := qiniu.UploadImage(l.ctx, byteData, imgName, imgHash, l.svcCtx.Config.Qiniu.Prefix,
			l.svcCtx.Config.Qiniu.AccessKey, l.svcCtx.Config.Qiniu.SecretKey, l.svcCtx.Config.Qiniu.Bucket,
			l.svcCtx.Config.Qiniu.Zone, l.svcCtx.Config.Qiniu.CDN)
		if err != nil {
			return xerr.NewErrMsg("七牛云上传图片失败")
		}
		// 保存图片
		i.Path = imgPath
		i.Type = 2
		i.Hash = &imgHash
		err = tx.Create(&i).Error
		if err != nil {
			return xerr.NewErrCode(xerr.DB_ERROR)
		}
		// 更新用户信息
		u.HeaderImageID = i.ID
		err = tx.Updates(&u).Error
		if err != nil {
			return xerr.NewErrCode(xerr.DB_UPDATE_ZERO_ERROR)
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}
