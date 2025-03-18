package data

import (
	"errors"
	"fmt"
	"review-service/internal/conf"
	"review-service/internal/data/query"
	"strings"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewDB, NewData, NewReviewRepo)

// Data .
type Data struct {
	// TODO wrapped database client
	query *query.Query
	log   *log.Helper
}

// NewData .
func NewDB(cfg *conf.Data) (*gorm.DB, error) {
	if cfg == nil {
		panic("[data] NewDB read cfg is nil")
	}
	switch strings.ToLower(cfg.Database.GetDriver()) {
	case "mysql":
		db, err := gorm.Open(mysql.Open(cfg.Database.GetSource()))
		if err != nil {
			panic(fmt.Errorf("[data] connect db fail: %w", err))
		}
		return db, nil
	default:
		return nil, errors.New("[data] connect to db fail, unsupported driver")
	}
}
func NewData(db *gorm.DB, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	query.SetDefault(db)
	return &Data{query: query.Q, log: log.NewHelper(logger)}, cleanup, nil
}
