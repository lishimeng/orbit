package main

import (
	"context"
	"fmt"
	"github.com/lishimeng/app-starter"
	etc2 "github.com/lishimeng/app-starter/etc"
	"github.com/lishimeng/go-log"
	"github.com/lishimeng/orbit/cmd/orbit/ddd"
	"github.com/lishimeng/orbit/internal/etc"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	err := _main()
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(time.Millisecond * 100)
}

func _main() (err error) {
	configName := "config"

	application := app.New()

	err = application.Start(func(ctx context.Context, builder *app.ApplicationBuilder) error {

		var err error
		err = builder.LoadConfig(&etc.Config, func(loader etc2.Loader) {
			loader.SetFileSearcher(configName, ".").SetEnvPrefix("").SetEnvSearcher()
		})
		if err != nil {
			return err
		}
		if len(etc.Config.Web.Listen) == 0 {
			etc.Config.Web.Listen = ":80"
		}
		//配置过期时间
		if etc.Config.Token.Ttl <= 0 {
			etc.Config.Token.Ttl = 24 // 默认24小时
		}
		etc.TokenTTL = time.Hour * 24 * time.Duration(etc.Config.Token.Ttl)

		builder.SetWebLogLevel("debug").
			EnableWeb(etc.Config.Web.Listen, ddd.Route).
			PrintVersion()
		//ComponentBefore(setup.JobClearExpireTask).
		//ComponentBefore(setup.BeforeStarted).
		//ComponentAfter(setup.AfterStarted)

		return err
	}, func(s string) {
		log.Info(s)
	})
	return
}
