package timelog

import (
	"time"

	"github.com/google/uuid"
)

type timelogMockFactory struct {
	now       time.Time
	uuids     []uuid.UUID
	uuidIndex int
}

func (mock *timelogMockFactory) NewTime() time.Time {
	return mock.now
}

func (mock *timelogMockFactory) NewUUID() uuid.UUID {
	uid := uuid.New()
	if mock.uuidIndex < len(mock.uuids) {
		uid = mock.uuids[mock.uuidIndex]
	}
	mock.uuidIndex++
	if mock.uuidIndex >= len(mock.uuids) {
		mock.uuidIndex = 0
	}

	return uid
}
