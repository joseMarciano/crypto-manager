package modules

import (
	"github.com/joseMarciano/crypto-manager/internal/app/user/handler/createuser"
	"github.com/joseMarciano/crypto-manager/internal/app/user/repository"
	createuseruc "github.com/joseMarciano/crypto-manager/internal/app/user/usecase/createuser"
	"github.com/joseMarciano/crypto-manager/internal/infra"
	createuserpb "github.com/joseMarciano/crypto-manager/pkg/proto/user/create"
)

func userModule(app *infra.Application) {
	repo := repository.New(app.DB)
	createUserUC := createuseruc.New(repo, repo)

	createuserpb.RegisterCreateUserHandlerServer(app.Server.GRPC, createuser.New(createUserUC))
}
