package uuid

import (
	"github.com/rs/xid"
	uuid "github.com/satori/go.uuid"
)

func Uuid() string {
	return uuid.NewV4().String()
}

func ShortUuid() string {
	return xid.New().String()
}
