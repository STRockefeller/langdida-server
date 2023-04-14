package time

import (
	"time"

	"github.com/golang/protobuf/ptypes/timestamp"
)

type UnixTime int64

// NewFromTime converts a time.Time value to UnixTime.
func NewFromTime(t time.Time) UnixTime {
	return UnixTime(t.Unix())
}

// NewFromTimeStamp converts a timestamp.Timestamp value to UnixTime.
func NewFromTimeStamp(ts *timestamp.Timestamp) UnixTime {
	return UnixTime(ts.GetSeconds())
}

// ToTime converts UnixTime to a time.Time value.
func (u UnixTime) ToTime() time.Time {
	return time.Unix(int64(u), 0)
}

// ToTimeStamp converts UnixTime to a timestamp.Timestamp value.
func (u UnixTime) ToTimeStamp() *timestamp.Timestamp {
	return &timestamp.Timestamp{
		Seconds: int64(u),
	}
}
