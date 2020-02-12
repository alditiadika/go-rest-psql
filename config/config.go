package config

//Config struct
type Config struct {
	Host     string
	Port     uint32
	DBName   string
	User     string
	Password string
}

//GetConf function
func GetConf() Config {
	return Config{
		Host:     "localhost",
		Port:     5432,
		DBName:   "postgres",
		User:     "postgres",
		Password: "P@ssw0rd",
	}
}
