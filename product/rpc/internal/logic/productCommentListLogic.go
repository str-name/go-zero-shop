package logic

import (
	"context"
	"github.com/pkg/errors"
	"zero-shop/common/globalKey"
	"zero-shop/common/xerr"
	"zero-shop/product/db/model"

	"zero-shop/product/rpc/internal/svc"
	"zero-shop/product/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type ProductCommentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewProductCommentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductCommentListLogic {
	return &ProductCommentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ProductCommentListLogic) ProductCommentList(in *pb.ProductCommentListReq) (*pb.ProductCommentListResp, error) {
	// todo: add your logic here and delete this line

	/*
		商品评论所有的评论都为根评论，不存在
	*/

	offset, limit := int((in.Page-1)*in.Size), int(in.Size)
	// 获取该商品的所有根评论
	var bootComm []model.ProductComment
	err := l.svcCtx.ProductDB.Where("product_id = ? and del_state = ?", in.ProductID, globalKey.DelStateNo).
		Offset(offset).Limit(limit).Find(&bootComm).Error
	if err != nil {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DB_ERROR), "MYSQL FIND product`comments ERROR: %+v", err)
	}

	var list []*pb.Comment
	for _, bc := range bootComm {
		list = append(list, &pb.Comment{
			ID:         bc.ID,
			UserID:     bc.UserID,
			ProductID:  bc.ProductID,
			IsGood:     bc.IsGood,
			Content:    bc.Content,
			AddContent: bc.AddContent,
		})
	}

	return &pb.ProductCommentListResp{Comments: list}, nil
}
