package conf

type Config struct {
	DB   *DBConfig
	Vars *VarConfig
}

type DBConfig struct {
	Username string
	Password string
	Dbname   string
	Dbhost   string
	Dbport   string
	Charset  string
}

type VarConfig struct {
	Port string
}

//Get config
func GetConfig() *Config {
	return &Config{
		DB: &DBConfig{
			Username: "admin",
			Password: "123456",
			Dbhost:   "localhost",
			Dbport:   "3306",
			Dbname:   "service_grpc",
			Charset:  "utf8",
		},
		Vars: &VarConfig{
			Port: ":8091",
		},
	}
}
