package config

type DatabaseType struct {
	Connection string
	Host       string
	Port       int
	Username   string
	Password   string
}

var (
	Database = DatabaseType{
		Connection: "sqlite",
		Host:       "database.db",
	}
)

func (d *DatabaseType) Set(db *DatabaseType) {
	Database = *db
}
