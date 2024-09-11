package main

// Load mớ này ở đâu huuhuu!
type ServiceConfig struct {
	BuildEnv string `mapstructure:"BUILD_ENV"`

	// service info
	ServiceName    string `mapstructure:"SERVICE_NAME"`
	ServiceHost    string `mapstructure:"SERVICE_HOST"`
	ServicePort    int    `mapstructure:"SERVICE_PORT"`
	ServiceTimeout int    `mapstructure:"SERVICE_TIMEOUT"`

	// database info
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBName     string `mapstructure:"DB_NAME"`
	DBUserName string `mapstructure:"DB_USERNAME"`
	DBPassword string `mapstructure:"DB_PASSWORD"`

	JwtSecretKey string `mapstructure:"JWT_SECRET_KEY"`

	MailFrom   string `mapstructure:"MAIL_FROM"`
	MailServer string `mapstructure:"MAIL_SERVER"`
	MailPort   string `mapstructure:"MAIL_PORT"`
	MailPass   string `mapstructure:"MAIL_PASS"`
}

// func mustConfig(configPath string) ServiceConfig {
// 	var result ServiceConfig

// }
