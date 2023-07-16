package merchant

import (
	"testing"

	"github.com/etnk125/go-webserver-merchant-management/model"
	"github.com/stretchr/testify/assert"
)

type MockMerchantRepository struct {
	model.MemoryDB
}

func NewMockMerchantRepository() *MockMerchantRepository {
	return &MockMerchantRepository{
		MemoryDB: model.MemoryDB{
			Merchants:         make(map[string]*model.Merchant),
			Products:          make(map[string]*model.Product),
			Transactions:      make(map[string]*map[string]*[]model.Transaction),
			DefaultCredential: model.Credential{Username: "test_username", Password: "test_password"},
		},
	}
}

func (r *MockMerchantRepository) GetDefaultCredential() *model.Credential {
	return &r.DefaultCredential
}

func (r *MockMerchantRepository) CreateMerchant(merchant *model.Merchant) (*model.Merchant, error) {
	r.Merchants[merchant.ID] = merchant
	return merchant, nil
}

func (r *MockMerchantRepository) GetMerchantInfo(id string) (*model.Merchant, error) {
	merchant, ok := r.Merchants[id]
	if !ok {
		return nil, ErrMerchantNotFound()
	}
	return merchant, nil
}

// -------------------------------
func TestService_RegisterMerchant(t *testing.T) {

	type testCase struct {
		req      *model.RegisterMerchantRequest
		expected *model.Merchant
		err      error
	}
	tcs := map[string]testCase{
		"should return correct merchant info": {
			req: &model.RegisterMerchantRequest{
				Name:        "test_name",
				BankAccount: "test_bank_account",
			},
			expected: &model.Merchant{
				Name:        "test_name",
				BankAccount: "test_bank_account",
				Username:    "test_username",
				Password:    "test_password",
			},
			err: nil,
		},
	}
	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			repo := NewMockMerchantRepository()
			svc := NewMerchantService(repo)
			actual, err := svc.RegisterMerchant(tc.req)

			// err
			assert.Equal(t, tc.err, err)
			// merchant
			assert.NotNil(t, actual)
			assert.NotNil(t, actual.ID)
			assert.Equal(t, tc.expected.Name, actual.Name)
			assert.Equal(t, tc.expected.BankAccount, actual.BankAccount)
			assert.Equal(t, tc.expected.Username, actual.Username)
			assert.Equal(t, tc.expected.Password, actual.Password)
		})
	}
}

func TestService_GetMerchantInfo(t *testing.T) {
	type testCase struct {
		id       string
		expected *model.Merchant
		err      error
	}
	tcs := map[string]testCase{
		"should return correct merchant info": {
			id: "test_id",
			expected: &model.Merchant{
				ID:          "test_id",
				Name:        "test_name",
				BankAccount: "test_bank_account",
				Username:    "test_username",
				Password:    "test_password",
			},
			err: nil,
		},
		"should return error when merchant not found": {
			id:       "not_found_id",
			expected: nil,
			err:      ErrMerchantNotFound(),
		},
	}
	for name, tc := range tcs {
		t.Run(name, func(t *testing.T) {
			repo := NewMockMerchantRepository()
			repo.Merchants["test_id"] = &model.Merchant{
				ID:          "test_id",
				Name:        "test_name",
				BankAccount: "test_bank_account",
				Username:    "test_username",
				Password:    "test_password",
			}
			svc := NewMerchantService(repo)
			actual, err := svc.GetMerchantInfo(tc.id)

			// err
			assert.Equal(t, tc.err, err)
			// merchant
			assert.Equal(t, tc.expected, actual)
		})
	}
}
