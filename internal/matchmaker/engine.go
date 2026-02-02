package matchmaker

import (
	"log"

	"github.com/AndreMicheletti/matchmaker-server-go/internal/database"
)

type Engine struct {
	redis *database.Redis
	cmdCh chan []byte
}

func NewEngine(redis database.Redis) Engine {
	cmdCh := make(chan []byte)
	engine := Engine{redis: &redis, cmdCh: cmdCh}
	go engine.processCmd()
	return engine
}

func (eng *Engine) RedisCli() *database.Redis {
	return eng.redis
}

func (eng *Engine) CommandChannel() chan []byte {
	return eng.cmdCh
}

func (eng *Engine) Close() {
	close(eng.cmdCh)
	eng.redis.Close()
}

func (eng *Engine) processCmd() {
	for cmd_bytes := range eng.cmdCh {
		cmd := string(cmd_bytes)
		switch cmd {
		case "!findmatch":
			eng.findMatch()
		}
	}
}

func (eng *Engine) findMatch() {
	log.Println("FIND MATCH COMMAND RECEIVED")
}