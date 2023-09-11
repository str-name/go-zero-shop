package storeProduct

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"io"
	"net/http"
	"strings"
	"zero-shop/common/ctxData"
	"zero-shop/common/qiniu"
	"zero-shop/common/tool"
	"zero-shop/common/xerr"
	"zero-shop/product/rpc/product"
	"zero-shop/user/rpc/user"

	"zero-shop/product/api/internal/svc"
	"zero-shop/product/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	r      *http.Request
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext, r *http.Request) *UpdateProductLogic {
	return &UpdateProductLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		r:      r,
	}
}

func (l *UpdateProductLogic) UpdateProduct(req *types.UpdateProductReq) error {
	// todo: add your logic here and delete this line

	// 获取bossID
	bossID := ctxData.GetUserIDFromCtx(l.ctx)
	// 判断用户是否存在
	existResp, err := l.svcCtx.UserRpc.CheckUserExists(l.ctx, &user.CheckUserExistsReq{UserID: bossID})
	if err != nil {
		return err
	}
	if !existResp.IsExists {
		return errors.Wrapf(xerr.NewErrCode(xerr.USER_NOT_EXISTS_ERROR), "USER NOT EXISTS, UserID: %v", bossID)
	}

	var banner string
	form := l.r.MultipartForm
	images := form.File["image"]
	if images == nil {
		banner = ""
	} else {
		for i, img := range images {
			imgName := img.Filename
			// 判断该文件是否为图片
			nameList := strings.Split(imgName, ".")
			suffix := strings.ToLower(nameList[len(nameList)-1])
			if !tool.InList(suffix, tool.WhiteImageList) {
				return xerr.NewErrMsg("请上传图片类型")
			}
			// 读取图片，生成图片hash
			imgObj, err := img.Open()
			if err != nil {
				return xerr.NewErrMsg("打开图片失败")
			}
			byteData, err := io.ReadAll(imgObj)
			if err != nil {
				return xerr.NewErrMsg("读取图片失败")
			}
			// 上传图片到七牛云
			imgPath, err := qiniu.UploadImage(l.ctx, byteData, imgName, req.Title, l.svcCtx.Config.Qiniu.Prefix,
				l.svcCtx.Config.Qiniu.AccessKey, l.svcCtx.Config.Qiniu.SecretKey, l.svcCtx.Config.Qiniu.Bucket,
				l.svcCtx.Config.Qiniu.Zone, l.svcCtx.Config.Qiniu.CDN)
			if err != nil {
				// 删除已经上传的图片
				go qiniu.DeleteImageList(banner, l.svcCtx.Config.Qiniu.AccessKey,
					l.svcCtx.Config.Qiniu.SecretKey, l.svcCtx.Config.Qiniu.Bucket, l.svcCtx.Config.Qiniu.Zone)
				return xerr.NewErrMsg("七牛云上传图片失败")
			}
			if i == 0 {
				banner = imgPath
				continue
			}
			banner = fmt.Sprintf("%s,%s", banner, imgPath)
		}
	}

	_, err = l.svcCtx.ProductRpc.UpdateProduct(l.ctx, &product.UpdateProductReq{
		ProductID:     req.ProductID,
		CategoryID:    req.CategoryID,
		Title:         req.Title,
		SubTitle:      req.SubTitle,
		Banner:        banner,
		Introduction:  req.Introduction,
		Price:         tool.YuanToFen(req.Price),
		DiscountPrice: tool.YuanToFen(req.DiscountPrice),
		OnSale:        req.OnSale,
		Stock:         req.Stock,
		BossID:        bossID,
	})

	return nil
}
