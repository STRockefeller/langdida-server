package review

import "time"

func NextReviewDate(familiarity int32) time.Time {
	now := time.Now()
	switch familiarity {
	case 0:
		return now
	case 1:
		return now.AddDate(0, 0, 1)
	case 2:
		return now.AddDate(0, 0, 3)
	case 3:
		return now.AddDate(0, 0, 7)
	case 4:
		return now.AddDate(0, 0, 15)
	case 5:
		return now.AddDate(0, 0, 30)
	case 6:
		return now.AddDate(0, 0, 60)
	case 7:
		return now.AddDate(0, 0, 90)
	default:
		return now.AddDate(0, 0, 90)
	}
}
