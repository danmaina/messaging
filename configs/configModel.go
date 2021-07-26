package configs

const (
	Path           = "./messaging.yaml"
	DefaultConfigs = `generalConfigs:
  port: "9002"
  logLevel: 3
email:
  host: "smtp.mailtrap.io"
  port: "587"
  username: "6b5e91a7479084"
  password: "5f96e81989ef58"
`
)

type Config struct {
	ApplicationConfigs GeneralConfigs `yaml:"generalConfigs"`
	Email              Email          `yaml:"email"`
}

type GeneralConfigs struct {
	Port     string `yaml:"port"`
	LogLevel int    `yaml:"logLevel"`
}

type Email struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}
