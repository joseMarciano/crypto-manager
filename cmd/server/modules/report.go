package modules

import (
	exchangerepository "github.com/joseMarciano/crypto-manager/internal/app/exchange/repository"
	balancehandler "github.com/joseMarciano/crypto-manager/internal/app/reports/balance/handler"
	balanceuc "github.com/joseMarciano/crypto-manager/internal/app/reports/balance/usecase"
	exchangetransactionhandler "github.com/joseMarciano/crypto-manager/internal/app/reports/exchangetransaction/handler"
	exchangetransactionrepository "github.com/joseMarciano/crypto-manager/internal/app/reports/exchangetransaction/repository"
	exchangetransactionuc "github.com/joseMarciano/crypto-manager/internal/app/reports/exchangetransaction/usecase"
	userrepository "github.com/joseMarciano/crypto-manager/internal/app/user/repository"
	"github.com/joseMarciano/crypto-manager/internal/infra"
	balancepb "github.com/joseMarciano/crypto-manager/pkg/proto/report/balance"
	exchangetransactionpb "github.com/joseMarciano/crypto-manager/pkg/proto/report/exchangetransaction"
)

func reportModule(app *infra.Application) {
	userRepo := userrepository.New(app.DB)
	exchangeRepo := exchangerepository.NewAccountRepository(app.DB)
	exchangeTransactionRepo := exchangetransactionrepository.NewTransactionRepository(app.DB)

	balanceUC := balanceuc.New(userRepo, exchangeRepo)
	exchangeTransactionUC := exchangetransactionuc.New(exchangeTransactionRepo)

	balancepb.RegisterBalanceHandlerServer(app.Server.GRPC, balancehandler.New(balanceUC))
	exchangetransactionpb.RegisterExchangeTransactionHandlerServer(app.Server.GRPC, exchangetransactionhandler.New(exchangeTransactionUC))
}
