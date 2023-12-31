// Code generated by goctl. DO NOT EDIT.
// Source: product.proto

package product

import (
	"context"

	"zero-shop/product/rpc/pb"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Carousel                 = pb.Carousel
	CarouselReq              = pb.CarouselReq
	CarouselResp             = pb.CarouselResp
	Category                 = pb.Category
	CategoryListReq          = pb.CategoryListReq
	CategoryListResp         = pb.CategoryListResp
	CategoryProductListReq   = pb.CategoryProductListReq
	CategoryProductListResp  = pb.CategoryProductListResp
	CheckProductExistsReq    = pb.CheckProductExistsReq
	CheckProductExistsResp   = pb.CheckProductExistsResp
	CheckSeckillExistsReq    = pb.CheckSeckillExistsReq
	CheckSeckillExistsResp   = pb.CheckSeckillExistsResp
	CollectProductListReq    = pb.CollectProductListReq
	CollectProductListResp   = pb.CollectProductListResp
	Comment                  = pb.Comment
	CreateCollectProductReq  = pb.CreateCollectProductReq
	CreateCollectProductResp = pb.CreateCollectProductResp
	CreateProductReq         = pb.CreateProductReq
	CreateProductResp        = pb.CreateProductResp
	CreateSeckillReq         = pb.CreateSeckillReq
	CreateSeckillResp        = pb.CreateSeckillResp
	DeleteCollectProductReq  = pb.DeleteCollectProductReq
	DeleteCollectProductResp = pb.DeleteCollectProductResp
	DeleteProductReq         = pb.DeleteProductReq
	DeleteProductResp        = pb.DeleteProductResp
	DeleteSeckillReq         = pb.DeleteSeckillReq
	DeleteSeckillResp        = pb.DeleteSeckillResp
	GetProductListByIDReq    = pb.GetProductListByIDReq
	GetProductListByIDResp   = pb.GetProductListByIDResp
	Product                  = pb.Product
	ProductCommentListReq    = pb.ProductCommentListReq
	ProductCommentListResp   = pb.ProductCommentListResp
	ProductDetailReq         = pb.ProductDetailReq
	ProductDetailResp        = pb.ProductDetailResp
	RecommendReq             = pb.RecommendReq
	RecommendResp            = pb.RecommendResp
	SearchProductReq         = pb.SearchProductReq
	SearchProductResp        = pb.SearchProductResp
	SeckillDetailReq         = pb.SeckillDetailReq
	SeckillDetailResp        = pb.SeckillDetailResp
	SeckillListReq           = pb.SeckillListReq
	SeckillListResp          = pb.SeckillListResp
	SeckillProduct           = pb.SeckillProduct
	ShelfProductReq          = pb.ShelfProductReq
	ShelfProductResp         = pb.ShelfProductResp
	SmallProduct             = pb.SmallProduct
	SmallSeckill             = pb.SmallSeckill
	SoldoutProductReq        = pb.SoldoutProductReq
	SoldoutProductResp       = pb.SoldoutProductResp
	UpdateProductReq         = pb.UpdateProductReq
	UpdateProductResp        = pb.UpdateProductResp

	ProductZrpcClient interface {
		// commonProduct
		Carousel(ctx context.Context, in *CarouselReq, opts ...grpc.CallOption) (*CarouselResp, error)
		CategoryList(ctx context.Context, in *CategoryListReq, opts ...grpc.CallOption) (*CategoryListResp, error)
		Recommend(ctx context.Context, in *RecommendReq, opts ...grpc.CallOption) (*RecommendResp, error)
		SearchProduct(ctx context.Context, in *SearchProductReq, opts ...grpc.CallOption) (*SearchProductResp, error)
		CategoryProductList(ctx context.Context, in *CategoryProductListReq, opts ...grpc.CallOption) (*CategoryProductListResp, error)
		ProductDetail(ctx context.Context, in *ProductDetailReq, opts ...grpc.CallOption) (*ProductDetailResp, error)
		ProductCommentList(ctx context.Context, in *ProductCommentListReq, opts ...grpc.CallOption) (*ProductCommentListResp, error)
		// storeProduct
		CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductResp, error)
		ShelfProduct(ctx context.Context, in *ShelfProductReq, opts ...grpc.CallOption) (*ShelfProductResp, error)
		UpdateProduct(ctx context.Context, in *UpdateProductReq, opts ...grpc.CallOption) (*UpdateProductResp, error)
		SoldoutProduct(ctx context.Context, in *SoldoutProductReq, opts ...grpc.CallOption) (*SoldoutProductResp, error)
		DeleteProduct(ctx context.Context, in *DeleteProductReq, opts ...grpc.CallOption) (*DeleteProductResp, error)
		CreateSeckill(ctx context.Context, in *CreateSeckillReq, opts ...grpc.CallOption) (*CreateSeckillResp, error)
		DeleteSeckill(ctx context.Context, in *DeleteSeckillReq, opts ...grpc.CallOption) (*DeleteSeckillResp, error)
		// userProduct
		CreateCollectProduct(ctx context.Context, in *CreateCollectProductReq, opts ...grpc.CallOption) (*CreateCollectProductResp, error)
		CollectProductList(ctx context.Context, in *CollectProductListReq, opts ...grpc.CallOption) (*CollectProductListResp, error)
		DeleteCollectProduct(ctx context.Context, in *DeleteCollectProductReq, opts ...grpc.CallOption) (*DeleteCollectProductResp, error)
		// seckillProduct
		SeckillList(ctx context.Context, in *SeckillListReq, opts ...grpc.CallOption) (*SeckillListResp, error)
		SeckillDetail(ctx context.Context, in *SeckillDetailReq, opts ...grpc.CallOption) (*SeckillDetailResp, error)
		// others
		CheckProductExists(ctx context.Context, in *CheckProductExistsReq, opts ...grpc.CallOption) (*CheckProductExistsResp, error)
		CheckSeckillExists(ctx context.Context, in *CheckSeckillExistsReq, opts ...grpc.CallOption) (*CheckSeckillExistsResp, error)
		GetProductListByID(ctx context.Context, in *GetProductListByIDReq, opts ...grpc.CallOption) (*GetProductListByIDResp, error)
	}

	defaultProductZrpcClient struct {
		cli zrpc.Client
	}
)

func NewProductZrpcClient(cli zrpc.Client) ProductZrpcClient {
	return &defaultProductZrpcClient{
		cli: cli,
	}
}

// commonProduct
func (m *defaultProductZrpcClient) Carousel(ctx context.Context, in *CarouselReq, opts ...grpc.CallOption) (*CarouselResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.Carousel(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) CategoryList(ctx context.Context, in *CategoryListReq, opts ...grpc.CallOption) (*CategoryListResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryList(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) Recommend(ctx context.Context, in *RecommendReq, opts ...grpc.CallOption) (*RecommendResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.Recommend(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) SearchProduct(ctx context.Context, in *SearchProductReq, opts ...grpc.CallOption) (*SearchProductResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.SearchProduct(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) CategoryProductList(ctx context.Context, in *CategoryProductListReq, opts ...grpc.CallOption) (*CategoryProductListResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CategoryProductList(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) ProductDetail(ctx context.Context, in *ProductDetailReq, opts ...grpc.CallOption) (*ProductDetailResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.ProductDetail(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) ProductCommentList(ctx context.Context, in *ProductCommentListReq, opts ...grpc.CallOption) (*ProductCommentListResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.ProductCommentList(ctx, in, opts...)
}

// storeProduct
func (m *defaultProductZrpcClient) CreateProduct(ctx context.Context, in *CreateProductReq, opts ...grpc.CallOption) (*CreateProductResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CreateProduct(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) ShelfProduct(ctx context.Context, in *ShelfProductReq, opts ...grpc.CallOption) (*ShelfProductResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.ShelfProduct(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) UpdateProduct(ctx context.Context, in *UpdateProductReq, opts ...grpc.CallOption) (*UpdateProductResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.UpdateProduct(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) SoldoutProduct(ctx context.Context, in *SoldoutProductReq, opts ...grpc.CallOption) (*SoldoutProductResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.SoldoutProduct(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) DeleteProduct(ctx context.Context, in *DeleteProductReq, opts ...grpc.CallOption) (*DeleteProductResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.DeleteProduct(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) CreateSeckill(ctx context.Context, in *CreateSeckillReq, opts ...grpc.CallOption) (*CreateSeckillResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CreateSeckill(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) DeleteSeckill(ctx context.Context, in *DeleteSeckillReq, opts ...grpc.CallOption) (*DeleteSeckillResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.DeleteSeckill(ctx, in, opts...)
}

// userProduct
func (m *defaultProductZrpcClient) CreateCollectProduct(ctx context.Context, in *CreateCollectProductReq, opts ...grpc.CallOption) (*CreateCollectProductResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CreateCollectProduct(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) CollectProductList(ctx context.Context, in *CollectProductListReq, opts ...grpc.CallOption) (*CollectProductListResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CollectProductList(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) DeleteCollectProduct(ctx context.Context, in *DeleteCollectProductReq, opts ...grpc.CallOption) (*DeleteCollectProductResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.DeleteCollectProduct(ctx, in, opts...)
}

// seckillProduct
func (m *defaultProductZrpcClient) SeckillList(ctx context.Context, in *SeckillListReq, opts ...grpc.CallOption) (*SeckillListResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.SeckillList(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) SeckillDetail(ctx context.Context, in *SeckillDetailReq, opts ...grpc.CallOption) (*SeckillDetailResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.SeckillDetail(ctx, in, opts...)
}

// others
func (m *defaultProductZrpcClient) CheckProductExists(ctx context.Context, in *CheckProductExistsReq, opts ...grpc.CallOption) (*CheckProductExistsResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CheckProductExists(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) CheckSeckillExists(ctx context.Context, in *CheckSeckillExistsReq, opts ...grpc.CallOption) (*CheckSeckillExistsResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.CheckSeckillExists(ctx, in, opts...)
}

func (m *defaultProductZrpcClient) GetProductListByID(ctx context.Context, in *GetProductListByIDReq, opts ...grpc.CallOption) (*GetProductListByIDResp, error) {
	client := pb.NewProductClient(m.cli.Conn())
	return client.GetProductListByID(ctx, in, opts...)
}
