package storage

import (
	"github.com/STRockefeller/go-linq"
	glinq "github.com/STRockefeller/gorm-linq"
	itime "github.com/STRockefeller/langdida-server/internal/time"
	"github.com/STRockefeller/langdida-server/models/protomodels"
)

type ListCardsRequest struct {
	conditions []glinq.QueryString
}

func NewListCardRequest() ListCardsRequest {
	return *new(ListCardsRequest)
}

func (req ListCardsRequest) where(condition glinq.QueryString) ListCardsRequest {
	req.conditions = append(req.conditions, condition)
	return req
}

func (req ListCardsRequest) WhereCardIndexIn(indices []protomodels.CardIndex) ListCardsRequest {
	return req.where(glinq.NewQueryString(
		"(name, language) IN ?",
		linq.Select(indices, func(index protomodels.CardIndex) []any {
			return []any{index.Name, index.Language}
		}).ToSlice(),
	))
}

func (req ListCardsRequest) WhereLanguage(lang protomodels.Language) ListCardsRequest {
	return req.where(glinq.NewQueryString("language = ?", lang))
}

func (req ListCardsRequest) WhereLabelContains(label string) ListCardsRequest {
	return req.where(glinq.NewQueryString("labels LIKE ?", "%"+label+"%"))
}

func (req ListCardsRequest) WhereNeedReview(today itime.UnixTime) ListCardsRequest {
	return req.where(glinq.NewQueryString("review_date < ?", today))
}

func (req ListCardsRequest) Where() []glinq.QueryString {
	return req.conditions
}
