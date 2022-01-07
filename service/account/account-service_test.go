package account

// import (
// 	"testing"

// 	"github.com/busraarsln/fintech-challenge/models"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/mock"
// )

// type MockRepository struct {
// 	mock.Mock
// }

// func (mock *MockRepository) AddAccount(account *models.Account) (*models.Account, error) {
// 	args := mock.Called()
// 	result := args.Get(0)
// 	return result.(*models.Account), args.Error(1)
// }
// func (mock *MockRepository) GetAccounts() ([]models.Account, error) {
// 	args := mock.Called()
// 	result := args.Get(0)
// 	return result.([]models.Account), args.Error(1)
// }

// func TestValidateEmptyAccount(t *testing.T) {
// 	testService := NewAccountService(nil, nil, nil)

// 	err := testService.Validate(nil)
// 	assert.NotNil(t, err)
// 	assert.Equal(t, "The account is empty", err.Error())

// }

// func TestGetAccounts(t *testing.T) {
// 	mockRepo := new(MockRepository)

// 	account := models.Account{ID: 1, AccountNo: "1212", Iban: "TR454"}
// 	//Setup expectations
// 	mockRepo.On("GetAccounts").Return([]models.Account{account}, nil)

// 	// testService := NewAccountService(mockRepo)

// 	// result, _ := testService.GetAccounts()

// 	//Mock Assertion: Behavioral
// 	mockRepo.AssertExpectations(t)

// 	//Data Assertion
// 	// assert.Equal(t, "1212", result[0].AccountNo)
// 	// assert.Equal(t, "TR454", result[0].Iban)

// }

// func TestAddAccount(t *testing.T) {
// 	mockRepo := new(MockRepository)

// 	account := models.Account{ID: 1, AccountNo: "1212", Iban: "TR454"}

// 	//Setup expectation
// 	mockRepo.On("AddAccount").Return(&account, nil)

// 	// testService := NewAccountService(mockRepo)

// 	// result, err := testService.AddAccount(&account)

// 	//Mock Assertion: Behavioral
// 	mockRepo.AssertExpectations(t)

// 	//Data Assertion

// 	// assert.NotNil(t, result.ID)
// 	// assert.Equal(t, "1212", result.AccountNo)
// 	// assert.Nil(t, err)

// }
