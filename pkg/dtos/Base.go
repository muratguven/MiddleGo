package dtos

import (
	"time"

	"github.com/google/uuid"
)

type BaseDto struct {
	Id           uuid.UUID `json:"id"`
	IsDeleted    bool      `json:"isDeleted"`
	CreatedDate  time.Time `json:"createdDate"`
	ModifiedDate time.Time `json:"modifiedDate"`
}
