package biz

import (
	"context"
	v1 "review-service/api/review/v1"
	"review-service/internal/data/model"
	"review-service/pkg/snowflake"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

// ReviewRepo is a Greater repo.
type ReviewRepo interface {
	SaveReview(context.Context, *model.ReviewInfo) (*model.ReviewInfo, error)
	GetReviewByOrderID(context.Context, int64) ([]*model.ReviewInfo, error)
	GetReviewByID(context.Context, int64) (*model.ReviewInfo, error)
	GetReplyInfoByReviewID(context.Context, int64) (*model.ReviewReplyInfo, error)
	SaveReply(context.Context, *model.ReviewReplyInfo) (*model.ReviewReplyInfo, error)
}

// ReviewUsecase is a review usecase.
type ReviewUsecase struct {
	repo ReviewRepo
	log  *log.Helper
}

func (uc *ReviewUsecase) ReplyReview(ctx context.Context, replyInfo *model.ReviewReplyInfo) (*model.ReviewReplyInfo, error) {
	uc.log.WithContext(ctx).Debugf("[biz] replyReview, req:%v\n", replyInfo)
	// 1. 参数校验
	// 1.1 未回复过
	info, err := uc.repo.GetReviewByID(ctx, replyInfo.ReviewID)
	if err != nil {
		return nil, v1.ErrorDbFailed("query failed")
	}
	if info.HasReply == 1 {
		return nil, v1.ErrorReviewReplyAlreadyExist("reply to review already exist")
	}

	// 1.2 水平越权
	if info.StoreID != replyInfo.StoreID {
		return nil, v1.ErrorStoreIDNotMatch("水平越级")
	}
	replyInfo.ReplyID = snowflake.GenID()
	return uc.repo.SaveReply(ctx, replyInfo)
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

func (uc *ReviewUsecase) GetReview(ctx context.Context, reviewID int64) (*model.ReviewInfo, error) {
	uc.log.WithContext(ctx).Debugf("[biz] GetReview, reviewID:%v\n", reviewID)
	review, err := uc.repo.GetReviewByID(ctx, reviewID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, v1.ErrorResultNotFound("no review found")
		}
		// todo record not found
		return nil, v1.ErrorDbFailed("query failed")
	}
	return review, nil
}

// NewReviewUsecase new a review usecase.
func NewReviewUsecase(repo ReviewRepo, logger log.Logger) *ReviewUsecase {
	return &ReviewUsecase{repo: repo, log: log.NewHelper(logger)}
}
