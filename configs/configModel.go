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

// Config Is a wrapper that wraps both GeneralConfigs and EmailConfigs
type Config struct {
	ApplicationConfigs GeneralConfigs `yaml:"generalConfigs"`
	Email              EmailConfigs   `yaml:"email"`
}

// GeneralConfigs are a set of application specific configurations
type GeneralConfigs struct {
	Port     string `yaml:"port"`
	LogLevel int    `yaml:"logLevel"`
}

// EmailConfigs is a set of configurable parameters from your email service provider
type EmailConfigs struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     string `yaml:"port"`
}
