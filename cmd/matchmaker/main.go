package main

import (
	"log"
	"github.com/AndreMicheletti/matchmaker-server-go/internal/config"
	"github.com/AndreMicheletti/matchmaker-server-go/internal/api"
	"github.com/AndreMicheletti/matchmaker-server-go/internal/matchmaker"
	"github.com/AndreMicheletti/matchmaker-server-go/internal/database"
)

func main() {
	// 1. Carregar Config
	cfg := config.Load()

	// 2. Iniciar Infra (Redis)
	redisClient := database.NewRedis(cfg)

	// 3. Iniciar Lógica (Matchmaker)
	engine := matchmaker.NewEngine(redisClient)

	// 4. Iniciar Transporte (API/Server)
	// Aqui você passa o Engine para o Servidor
	srv := api.NewServer(cfg.PORT, engine)
	defer srv.Close()

	// 5. Start
	log.Fatal(srv.Run())
}
