package db

func NewDBStore(engine string) Store {
	switch {
	default:
		return newMemoryStore()
	}
}
