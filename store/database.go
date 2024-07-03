package store

type Configure struct {
	Host     string
	Port     int
	Database string
	User     string
	Password string
	SSLMode  bool
}
