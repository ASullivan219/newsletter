package db

type SqliteDriver struct {
	I_database
}

func new_sqlite() SqliteDriver {
	return SqliteDriver{}
}

func (db *SqliteDriver) GetSubscriber(email string) (subscriberModel, error) {
	return subscriberModel{}, nil
}
func (db *SqliteDriver) PutSubscriber(email string, name string) error {
	return nil
}
