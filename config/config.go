package config

import (
	"log"

	"github.com/spf13/viper"
)

// Fungsi LoadConfig sekarang mengembalikan nilai bertipe Config
func LoadConfig() Config {
	var cfg Config

	// Menetapkan nama file konfigurasi (tanpa ekstensi)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./files/config") // Direktori tempat mencari file konfigurasi

	// Membaca file konfigurasi
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file, %v", err)
	}

	// Memasukkan data konfigurasi ke dalam struct cfg
	err = viper.Unmarshal(&cfg)
	if err != nil {
		log.Fatalf("Unable to decode into struct, %v", err)
	}

	// Mengembalikan objek cfg bertipe Config
	return cfg
}
