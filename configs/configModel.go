package configs

const (
	Path           = "./messaging.yaml"
	DefaultConfigs = `generalConfigs:
  port: "9002"
  logLevel: 3
email:
  host: "smtp.gmail.com"
  port: "587"
  username: "user"
  pass: "root"
  from: 20
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
	Pass     string `yaml:"pass"`
	From     string `yaml:"from"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}
