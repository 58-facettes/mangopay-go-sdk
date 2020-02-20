package data

type Manager interface {
	SaveIdempotencyKey(key, url string) error
}

type MokeDB struct {
	data map[string]*Idempotency
}

func NewMokeDB() Manager {
	return &MokeDB{
		data: make(map[string]*Idempotency),
	}
}

type Idempotency struct {
	Key string
	URL string
}

func (db *MokeDB) SaveIdempotencyKey(key, url string) error {
	db.data[key] = &Idempotency{
		Key: key,
		URL: url,
	}
	return nil
}
