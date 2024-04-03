package initialize

import (
	"github.com/ArtisanCloud/PowerWeChat/v3/src/kernel"
	"github.com/ArtisanCloud/PowerWeChat/v3/src/miniProgram"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
)

func MiniProgram() {
	mpCfg := global.GVA_CONFIG.MiniProgram
	redisCfg := global.GVA_CONFIG.Redis
	MiniProgramApp, err := miniProgram.NewMiniProgram(&miniProgram.UserConfig{
		AppID:     mpCfg.Appid,  // 小程序 appid
		Secret:    mpCfg.Secret, // 小程序 secret
		HttpDebug: mpCfg.HttpDebug,
		Log: miniProgram.Log{
			Level: mpCfg.LogLevel,
			File:  mpCfg.LogInfoFile,
			Error: mpCfg.LogErrorFile,
		},
		Cache: kernel.NewRedisClient(&kernel.UniversalOptions{
			Addrs:    []string{redisCfg.Addr},
			Password: redisCfg.Password,
			DB:       redisCfg.DB,
		}),
	})
	if err != nil {
		global.GVA_LOG.Error("MiniProgram init error, err:", zap.Error(err))
		panic(err)
	}

	global.GVA_LOG.Info("MiniProgram init success")
	global.MiniProgram = MiniProgramApp
}
