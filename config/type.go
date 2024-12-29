package config

type Config struct {
	DB              DBConfig              `json:"db"`
	Server          ServerConfig          `json:"server"`
	Redis           RedisConfig           `json:"redis"`
	SuperAdmin      SuperAdmin            `json:"superAdmin"`
	AdminService    AdminServiceConfig    `json:"adminService"`
	RoutemapService RoutemapServiceConfig `json:"routemapService"`
}

type DBConfig struct {
	Host     string `json:"host"`
	Port     uint   `json:"port"`
	Database string `json:"database"`
	Schema   string `json:"schema"`
	User     string `json:"user"`
	Password string `json:"password"`
}

type ServerConfig struct {
	HttpPort          uint   `json:"httpPort"`
	Secret            string `json:"secret"`
	AuthExpMinute     uint   `json:"authExpMin"`
	AuthRefreshMinute uint   `json:"authExpRefreshMin"`
}

type RedisConfig struct {
	Host string `json:"host"`
	Port uint   `json:"port"`
}

type SuperAdmin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AdminServiceConfig struct {
	Port uint `json:"port"`
}

type RoutemapServiceConfig struct {
	Port uint   `json:"port"`
	Host string `json:"host"`
}
