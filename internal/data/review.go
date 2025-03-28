package data

import (
	"review-service/internal/biz"

	"context"
	"review-service/internal/data/model"
	"review-service/internal/data/query"

	"github.com/go-kratos/kratos/v2/log"
)

type reviewRepo struct {
	data *Data
	log  *log.Helper
}

// NewReviewRepo .
func (r *reviewRepo) SaveReview(ctx context.Context, review *model.ReviewInfo) (*model.ReviewInfo, error) {
	err := r.data.query.ReviewInfo.WithContext(ctx).Save(review)
	return review, err
}

func (r *reviewRepo) GetReviewByOrderID(ctx context.Context, id int64) ([]*model.ReviewInfo, error) {
	return r.data.query.ReviewInfo.WithContext(ctx).Where(r.data.query.ReviewInfo.OrderID.Eq(id)).Find()
}

func (r *reviewRepo) GetReviewByID(ctx context.Context, id int64) (*model.ReviewInfo, error) {
	return r.data.query.ReviewInfo.WithContext(ctx).Where(r.data.query.ReviewInfo.ReviewID.Eq(id)).First()
}

func (r *reviewRepo) GetReplyInfoByReviewID(ctx context.Context, reviewID int64) (*model.ReviewReplyInfo, error) {
	return r.data.query.ReviewReplyInfo.WithContext(ctx).Where(r.data.query.ReviewReplyInfo.ReviewID.Eq(reviewID)).First()
}
func (r *reviewRepo) SaveReply(ctx context.Context, replyInfo *model.ReviewReplyInfo) (*model.ReviewReplyInfo, error) {
	err := r.data.query.Transaction(func(tx *query.Query) error {
		err := tx.ReviewReplyInfo.WithContext(ctx).Save(replyInfo)
		if err != nil {
			r.log.WithContext(ctx).Errorf("SaveReply create reply fail, err:%v", err)
			return err
		}
		_, err = tx.ReviewInfo.WithContext(ctx).
			Where(tx.ReviewInfo.ReviewID.Eq(replyInfo.ReviewID)).
			Update(r.data.query.ReviewInfo.HasReply, 1)
		if err != nil {
			r.log.WithContext(ctx).Errorf("SaveReply update reply fail, err:%v", err)
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return replyInfo, nil
}
func NewReviewRepo(data *Data, logger log.Logger) biz.ReviewRepo {
	return &reviewRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
