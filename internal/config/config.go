package config

type (
	Config struct {
		App     App     `yaml:"app"`
		Mysql   Mysql   `yaml:"mysql"`
		Account Account `yaml:"account"`
	}

	App struct {
		Name    string `yaml:"name" envconfig:"CLEANSERVICE_APP_NAME"`
		Address string `yaml:"port" envconfig:"CLEANSERVICE_APP_ADDRESS"`
	}

	Mysql struct {
		Username string `yaml:"username" envconfig:"CLEANSERVICE_MYSQL_USERNAME"`
		Password string `yaml:"password" envconfig:"CLEANSERVICE_MYSQL_PASSWORD"`
		DBName   string `yaml:"db_name" envconfig:"CLEANSERVICE_MYSQL_DBNAME"`
		Host     string `yaml:"host" envconfig:"CLEANSERVICE_MYSQL_HOST"`
		Port     string `yaml:"port" envconfig:"CLEANSERVICE_MYSQL_PORT"`
	}

	Account struct {
		MinUsernameLength int `yaml:"min_username_length"`
	}
)
