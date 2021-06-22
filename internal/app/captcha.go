package app

import (
	"go-gin/internal/app/config"
	"go-gin/pkg/logger"

	"github.com/LyricTian/captcha"
	"github.com/LyricTian/captcha/store"
	"github.com/go-redis/redis"
)

// InitCaptcha 初始化图形验证码
func InitCaptcha() {
	cfg := config.GetGlobalConfig().Captcha
	if cfg.Store == "redis" {
		rc := config.GetGlobalConfig().Redis
		captcha.SetCustomStore(store.NewRedisStore(&redis.Options{
			Addr:     rc.Addr,
			Password: rc.Password,
			DB:       cfg.RedisDB,
		}, captcha.Expiration, logger.StandardLogger(), cfg.RedisPrefix))
	}
}
