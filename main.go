package main

import (
	"flag"
	"log"
	"path/filepath"
	"youras/application/factory"
	"youras/infra/config"
	"youras/infra/database"
	"youras/infra/web"
	"youras/interfaces"
	"youras/pkg/ylog"
)

func main() {
	cfgPath := flag.String("c", ".", "config file path")
	flag.Parse()

	cfgFile := filepath.Join(*cfgPath, "config.yaml")
	cfg, err := config.ReadConfigFromYamlFile(cfgFile)
	if err != nil {
		log.Fatalf("read config file failed, err:%v", err)
	}
	ylog.InitProductionLogger(cfg.App.ServiceName, &cfg.Log)
	gormLogger := ylog.NewGormLogger(ylog.Log().Desugar(), &cfg.Log)
	db, err := database.OpenPg(&cfg.Db, gormLogger)
	if err != nil {
		ylog.Log().Fatalf("open database failed, err:%v", err)
	}
	if err = database.Migrate(db, cfg.Db.MigrationTable); err != nil {
		ylog.Log().Fatalf("migrate failed, err:%v", err)
	}
	f := factory.NewFactory(db)
	handler := interfaces.NewHttpHandler(f, cfg.App)
	s := web.NewServer(cfg.App.ServeAddr(), handler)
	ylog.Log().Info("start server,address:", cfg.App.ServeAddr())
	err = s.ListenAndServe()
	if err != nil {
		ylog.Log().Fatalf("start server failed, err:%v", err)
	}
}
