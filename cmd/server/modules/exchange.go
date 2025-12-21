package modules

import (
	"github.com/joseMarciano/crypto-manager/internal/app/exchange/handler/createaccount"
	"github.com/joseMarciano/crypto-manager/internal/app/exchange/handler/createexchange"
	"github.com/joseMarciano/crypto-manager/internal/app/exchange/handler/deposit"
	"github.com/joseMarciano/crypto-manager/internal/app/exchange/handler/withdraw"
	"github.com/joseMarciano/crypto-manager/internal/app/exchange/repository"
	createaccountuc "github.com/joseMarciano/crypto-manager/internal/app/exchange/usecase/createaccount"
	createexchangeuc "github.com/joseMarciano/crypto-manager/internal/app/exchange/usecase/createexchange"
	deposituc "github.com/joseMarciano/crypto-manager/internal/app/exchange/usecase/deposit"
	withdrawuc "github.com/joseMarciano/crypto-manager/internal/app/exchange/usecase/withdraw"
	userrepository "github.com/joseMarciano/crypto-manager/internal/app/user/repository"
	"github.com/joseMarciano/crypto-manager/internal/infra"
	createexchangepb "github.com/joseMarciano/crypto-manager/pkg/proto/exchange/create"
	createaccountpb "github.com/joseMarciano/crypto-manager/pkg/proto/exchange/createaccount"
	depositpb "github.com/joseMarciano/crypto-manager/pkg/proto/exchange/deposit"
	withdrawpb "github.com/joseMarciano/crypto-manager/pkg/proto/exchange/withdraw"
)

func exchangesModule(app *infra.Application) {
	exchangeRepo := repository.NewExchangeRepository(app.DB)
	accountRepo := repository.NewAccountRepository(app.DB)
	transactionRepo := repository.NewTransactionRepository(app.DB)
	userRepo := userrepository.New(app.DB)

	createExchangeUC := createexchangeuc.New(exchangeRepo)
	createAccountUC := createaccountuc.New(accountRepo, accountRepo, exchangeRepo, userRepo)
	depositUC := deposituc.New(transactionRepo, accountRepo, accountRepo, exchangeRepo)
	withDrawUC := withdrawuc.New(transactionRepo, accountRepo, accountRepo, exchangeRepo)

	createexchangepb.RegisterCreateExchangeHandlerServer(app.Server.GRPC, createexchange.New(createExchangeUC))
	createaccountpb.RegisterCreateAccountHandlerServer(app.Server.GRPC, createaccount.New(createAccountUC))
	depositpb.RegisterDepositHandlerServer(app.Server.GRPC, deposit.New(depositUC))
	withdrawpb.RegisterWithdrawHandlerServer(app.Server.GRPC, withdraw.New(withDrawUC))
}
