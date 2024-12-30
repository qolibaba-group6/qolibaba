package config
import (
	"log"
	"os"
	"github.com/spf13/viper"
)
type Config struct {
	Server struct {
		Port int `mapstructure:"port"`
	} `mapstructure:"server"`
	Database struct {
		Host     string `mapstructure:"host"`
		Port     int    `mapstructure:"port"`
		User     string `mapstructure:"user"`
		Password string `mapstructure:"password"`
		Name     string `mapstructure:"name"`
	} `mapstructure:"database"`
	JWT struct {
		Secret string `mapstructure:"secret"`
		Expiry string `mapstructure:"expiry"`
	} `mapstructure:"jwt"`
}
var AppConfig Config
func LoadConfig() {
	configPath := os.Getenv("CONFIG_PATH")
	if configPath == "" {
		configPath = "/Users/ehsansobhani/Desktop/travel-booking-app/config/config.yaml" // مقدار پیش‌فرض
	}
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file at %s, %s", configPath, err)
	}
	err = viper.Unmarshal(&AppConfig)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}
	log.Printf("Config loaded from %s", configPath)
}
