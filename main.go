package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"tugas/database"
	"tugas/handlers"
	"tugas/repositories"
	"tugas/services"

	"github.com/spf13/viper"
)

type Config struct {
	Port	string	`mapstructure:"PORT"`
	DB		string	`mapstructure:"DB"`
}

func main() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".","_"))

	if _, err := os.Stat(".env"); err == nil {
		viper.SetConfigFile(".env")
		_ = viper.ReadInConfig();
	}

	config := Config{
		Port: viper.GetString("PORT"),
		DB: viper.GetString("DB"),
	}
	
	db, err := database.InitDB(config.DB)
	if err != nil {
		log.Fatal("Gagal dalam mengisialisasi database:", err)
	}
	defer db.Close()
	
	CategoriesRepo := repositories.NewCategoriesRepository(db)
	CategoriesService := services.NewCategoriesService(CategoriesRepo)
	CategoriesHandler := handlers.NewCategoriesHandler(CategoriesService)

	http.HandleFunc("/api/categories", CategoriesHandler.HandleCategories)
	http.HandleFunc("/api/categories/", CategoriesHandler.HandleCategoriesByID)

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"stats" : "OK",
			"message": "API Jalan",
		})
	})

	fmt.Println("Server berjalan di localhost:8000")
	err = http.ListenAndServe(":"+config.Port, nil)
	if err != nil {
		fmt.Println("Gagal jalankan server")
	}
}
