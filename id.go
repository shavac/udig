package graph

import(
	uuid "github.com/streadway/simpleuuid"
	"time"
)

func NewID() (uuid.UUID, error) {
	if uid, err := uuid.NewTime(time.Now()); err != nil {
		return nil, err
	} else {
		return uid, nil
	}
}
