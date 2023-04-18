package gormmodels

import (
	"time"

	"github.com/STRockefeller/langdida-server/models/protomodels"
)

type Log struct {
	Date          string `gorm:"primaryKey"`
	ReviewCards   int32
	NewCards      int32
	Streak        int32
	StreakUpdated bool
}

func NewDefaultLog(newStreak int32) Log {
	date := time.Now().Format("2006-01-02")
	return Log{
		Date:        date,
		ReviewCards: 0,
		NewCards:    0,
		Streak:      newStreak,
	}
}

func (l Log) WithReviewedCard() Log {
	l.ReviewCards++
	return l
}

func (l Log) WithNewCard() Log {
	l.NewCards++
	return l
}

func NewLog(l protomodels.Log) Log {
	return Log{
		Date:        l.Date,
		ReviewCards: l.ReviewCards,
		NewCards:    l.NewCards,
		Streak:      l.Streak,
	}
}

func (l Log) ToProtoModel() protomodels.Log {
	return protomodels.Log{
		Date:        l.Date,
		ReviewCards: l.ReviewCards,
		NewCards:    l.NewCards,
		Streak:      l.Streak,
	}
}
