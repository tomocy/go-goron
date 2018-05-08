package storages

import "errors"

type Storage interface {
	InitSession()
}

func Get(storage string) (Storage, error) {
	switch storage {
	case "memory":
		return &memory{}, nil
	default:
		return nil, errors.New("Not found")
	}
}
