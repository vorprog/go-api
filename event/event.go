package event

import (
	"time"

	"github.com/google/uuid"
)

type Event struct {
	Id        uuid.UUID `json:"id"`
	Type      string    `json:"type"`
	SourceId  uuid.UUID `json:"sourceId"`
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}
