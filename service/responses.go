package service

import "github.com/STRockefeller/langdida-server/models/protomodels"

type LogStatus struct {
	CardsShouldBeReviewed    []protomodels.CardIndex
	FamiliarityDistribution  [10]int
	Streak                   int
	NewCardCountToday        int
	ReviewedCardCountToday   int
	NewCardCountForWeek      [7]int
	ReviewedCardCountForWeek [7]int
}
