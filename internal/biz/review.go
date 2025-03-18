package biz

import (
	"context"
	v1 "review-service/api/review/v1"
	"review-service/internal/data/model"
	"review-service/pkg/snowflake"

	"github.com/go-kratos/kratos/v2/log"
)

// ReviewRepo is a Greater repo.
type ReviewRepo interface {
	SaveReview(context.Context, *model.ReviewInfo) (*model.ReviewInfo, error)
	GetReviewByOrderID(context.Context, int64) ([]*model.ReviewInfo, error)
}

// ReviewUsecase is a review usecase.
type ReviewUsecase struct {
	repo ReviewRepo
	log  *log.Helper
}

// NewReviewUsecase new a review usecase.
func NewReviewUsecase(repo ReviewRepo, logger log.Logger) *ReviewUsecase {
	return &ReviewUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateReview creates a Greeter, and returns the new Greeter.
// func (uc *ReviewUsecase) CreateReview(ctx context.Context, r *Review) (*Greeter, error) {
// 	uc.log.WithContext(ctx).Infof("CreateGreeter: %v")
// 	return uc.repo.Save(ctx, g)
// }

func (uc ReviewUsecase) CreateReview(ctx context.Context, review *model.ReviewInfo) (*model.ReviewInfo, error) {
	uc.log.WithContext(ctx).Debugf("[biz] CreateReview, req:%v\n", review)
	// 1、数据校验
	reviewList, err := uc.repo.GetReviewByOrderID(ctx, review.OrderID)
	if err != nil {
		return nil, v1.ErrorDbFailed("query failed")
	}
	if len(reviewList) > 0 {
		//fmt.Println("已评价")
		return nil, v1.ErrorOrderReviewd("order id:%d already exist review", review.OrderID)
	}
	// 2、生成reviewID
	review.ReviewID = snowflake.GenID()
	// 3、查询订单和商品快照消息
	// rpc调用订单服务和商家服务

	// 4、瓶装数据入库
	return uc.repo.SaveReview(ctx, review)
}
