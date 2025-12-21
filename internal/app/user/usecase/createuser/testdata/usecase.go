package testdata

import (
	"time"

	"github.com/joseMarciano/crypto-manager/internal/app/user/domain"
	"github.com/joseMarciano/crypto-manager/internal/app/user/usecase/createuser"
)

var MockedDate = time.Now()
var MockedUUID = "b48f6da2-adda-46be-9e2f-17f2be1b0a28"

func ExpectedCreatedUser() domain.User {
	i := DefaultInput()
	return domain.User{ID: MockedUUID, Name: i.Name, Birthday: i.Birthday, DocumentNumber: i.DocumentNumber}
}

func DefaultInput() createuser.Input {
	return createuser.Input{Name: "John", Birthday: MockedDate, DocumentNumber: "1234"}
}

func ExpectedOutput() createuser.Output {
	i := DefaultInput()
	return createuser.Output{ID: MockedUUID, Name: i.Name, Birthday: i.Birthday, DocumentNumber: i.DocumentNumber}
}
