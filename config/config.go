package config

type Config struct {
	DbHost     string
	DbName     string
	DbUser     string
	DbPassword string
}

func GetConfig() Config {
	return Config{
		DbHost:     "localhost",
		DbName:     "todo-ar-db",
		DbUser:     "postgres",
		DbPassword: "postgres",
	}
}
