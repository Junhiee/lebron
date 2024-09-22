package logic

import (
	"context"

	"github.com/Junhiee/lebron/apps/app/internal/svc"
	"github.com/Junhiee/lebron/apps/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CategoryListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 分类商品列表
func NewCategoryListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CategoryListLogic {
	return &CategoryListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CategoryListLogic) CategoryList(req *types.CategoryListRequest) (resp *types.CategoryListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
