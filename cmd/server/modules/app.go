package modules

import "github.com/joseMarciano/crypto-manager/internal/infra"

func NewApp() *infra.Application {
	app := infra.New()
	userModule(app)
	exchangesModule(app)
	reportModule(app)
	return app
}
