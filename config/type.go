package config

type Config struct {
	DB              DBConfig              `json:"db"`
	Server          ServerConfig          `json:"server"`
	Redis           RedisConfig           `json:"redis"`
	SuperAdmin      SuperAdmin            `json:"superAdmin"`
	RoutemapService RoutemapServiceConfig `json:"routemapService"`
	AdminService AdminServiceConfig `json:"adminService"`
	HotelService HotelServiceConfig `json:"hotelService"`
	BankService  BankServiceConfig `json:"bankService"`
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
type HotelServiceConfig struct {
	Port uint `json:"port"`
}
type BankServiceConfig struct {
	Port uint `json:"port"`
}

type RoutemapServiceConfig struct {
	Port uint   `json:"port"`
	Host string `json:"host"`
}
