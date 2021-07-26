package configs

import (
	"github.com/danmaina/models"
)

const (
	Path           = "./messaging.yaml"
	DefaultConfigs = `generalConfigs:
  port: "9001"
  logLevel: 3
mysql:
  host: "localhost"
  port: "3306"
  username: "root"
  password: "root"
  database: "profiles"
  totalConnections: 20
  maxIdleConnections: 5
  maxLifetime: 180
redis:
  host: "localhost"
  port: "6379"
  connections: 100
  connectionType: "TCP"
profileStatusIds:
  active: 1
  dormant: 2
  deleted: 3
  locked: 4
  unverified: 5
  archived: 6
`
)

type Config struct {
	ApplicationConfigs   GeneralConfigs     `yaml:"generalConfigs"`
	Email                models.Email      `yaml:"mysql"`
	Redis                storage.Redis      `yaml:"redis"`
	ProfileStatusCodeIds ProfileStatusCodes `yaml:"profileStatusIds"`
}

type GeneralConfigs struct {
	Port     string `yaml:"port"`
	LogLevel int    `yaml:"logLevel"`
}

type ProfileStatusCodes struct {
	Active     uint64 `yaml:"active"`
	Dormant    uint64 `yaml:"dormant"`
	Deleted    uint64 `yaml:"deleted"`
	Locked     uint64 `yaml:"locked"`
	Unverified uint64 `yaml:"unverified"`
	Archived   uint64 `yaml:"archived"`
}
