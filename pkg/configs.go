package pkg


import (
	"io/ioutil"
	"gopkg.in/yaml.v2"
)

// Mysql
type Mysql struct {
	Host     string `yaml:"host"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Log      bool   `yaml:"log"`
}

// MiddlewareElastic
type MiddlewareElastic struct {
	Enabled bool `yaml:"enabled"`
}

// Elastic
type Elastic struct {
	Log  bool     `yaml:"log"`
	Host []string `yaml:"host"`
}

// Credential
type Credential struct {
	AuthMechanism string `yaml:"authMechanism"`
	AuthSource    string `yaml:"authSource"`
	Username      string `yaml:"username"`
	Password      string `yaml:"password"`
	PasswordSet   bool   `yaml:"passwordSet"`
}

// Storage
type Storage struct {
	Option     string     `yaml:"option"`
	UrlExpire  int        `yaml:"urlExpire"`
	PartExpire int        `yaml:"partExpire"`
	Launch     bool       `yaml:"launch"`
	Storages   []Storages `yaml:"storages"`
}

// Kafka
type Kafka struct {
	Broker []string `yaml:"broker"`
}

// MiddlewareKafka
type MiddlewareKafka struct {
	Enabled bool `yaml:"enabled"`
}

// MiddlewareRedis
type MiddlewareRedis struct {
	Enabled  bool   `yaml:"enabled"`
	Password string `yaml:"password"`
}

// Middleware
type Configs struct {
	Image            Image          `yaml:"image"`
	ImagePullSecrets interface{}    `yaml:"imagePullSecrets"`
	Config           Config         `yaml:"config"`
	Elastic          MiddlewareElastic `yaml:"elastic"`
	Mongo            MiddlewareMongo   `yaml:"mongo"`
	Kafka            MiddlewareKafka   `yaml:"kafka"`
	Minio            Minio          `yaml:"minio"`
	Mysql            MiddlewareMysql   `yaml:"mysql"`
	Redis            MiddlewareRedis   `yaml:"redis"`
}

// Config
type Config struct {
	Redis   Redis   `yaml:"redis"`
	Elastic Elastic `yaml:"elastic"`
	Kafka   Kafka   `yaml:"kafka"`
	Mongo   Mongo   `yaml:"mongo"`
	Email   Email   `yaml:"email"`
	Storage Storage `yaml:"storage"`
	Mysql   Mysql   `yaml:"mysql"`
}

// Redis
type Redis struct {
	Addrs    []string    `yaml:"addrs"`
	Username interface{} `yaml:"username"`
	Password string      `yaml:"password"`
}

// Minio
type Minio struct {
	Enabled   bool   `yaml:"enabled"`
	AccessKey string `yaml:"accessKey"`
	SecretKey string `yaml:"secretKey"`
}

// Email
type Email struct {
	Emails []Emails `yaml:"emails"`
}

// Emails
type Emails struct {
	Password  string `yaml:"password"`
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Emailfrom string `yaml:"emailfrom"`
	Username  string `yaml:"username"`
	Aliasname string `yaml:"aliasname"`
}

// Storages
type Storages struct {
	Name       string `yaml:"name"`
	Protocol   string `yaml:"protocol"`
	Domain     string `yaml:"domain"`
	AccessKey  string `yaml:"accessKey"`
	SecretKey  string `yaml:"secretKey"`
	Location   string `yaml:"location"`
	BucketName string `yaml:"bucketName"`
}

// MiddlewareMysql
type MiddlewareMysql struct {
	Enabled      bool   `yaml:"enabled"`
	RootPassword string `yaml:"rootPassword"`
}

// Image
type Image struct {
	Repo string `yaml:"repo"`
	Tag  string `yaml:"tag"`
}

// Mongo
type Mongo struct {
	Hosts      []string   `yaml:"hosts"`
	Direct     bool       `yaml:"direct"`
	Credential Credential `yaml:"credential"`
}

// MiddlewareMongo
type MiddlewareMongo struct {
	Password string `yaml:"rootPassword"`
	Enabled  bool   `yaml:"enabled"`
}

func ParaseConfig(filepath string) (*Configs,error) {
	c := new(Configs)
	buf,err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil,err
	}
	err = yaml.Unmarshal(buf,c)
	if err != nil {
		return nil,err
	}
	return c,nil
}
