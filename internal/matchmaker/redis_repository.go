package matchmaker

import (
	"fmt"
	"log"
	"context"
	"github.com/fasmat/trueskill"
)

var ctx = context.Background()

func keyForPlayer(playerId int) string {
	return fmt.Sprintf("player:%d", playerId)
}

func (eng *Engine) SetPlayer(player trueskill.Player) {
	hashFields := []string{
		"id", fmt.Sprint(player.GetID()),
		"mu", fmt.Sprintf("%f", player.GetMu()),
		"sigma", fmt.Sprintf("%f", player.GetSigma()),
	}

	res, err := eng.redis.Client().HSet(ctx, keyForPlayer(player.GetID()), hashFields).Result()

	if err != nil {
		panic(err)
	}
	log.Printf("[REDIS] SETTING PLAYER %s %d\n", fmt.Sprint(player.GetID()), res)
}

func (eng *Engine) GetPlayerOrDefault(playerId int) trueskill.Player {
	exists, err := eng.redis.Client().HExists(ctx, keyForPlayer(playerId), "id").Result()

	if exists && err != nil {
		res, err := eng.redis.Client().HGetAll(ctx, keyForPlayer(playerId)).Result()

		if err != nil {
			panic(err)
		}
		log.Println(res)
	}

	return trueskill.NewDefaultPlayer(playerId)
}
