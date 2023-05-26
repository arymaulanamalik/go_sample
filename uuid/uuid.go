package uuid

import (
	uid "github.com/google/uuid"
)

type UUID struct {
	UUID uid.UUID
}

func NewUUID() UUID {
	uid := uid.New()

	return UUID{UUID: uid}
}

func NewUUIDString() string {
	uid := uid.New()

	return uid.String()
}

func (u *UUID) String() string {
	return u.UUID.String()
}

func FromStr(r string) (uid.UUID, error) {
	uid, err := uid.FromBytes([]byte(r))
	if err != nil {
		return uid, err
	}
	return uid, nil
}
