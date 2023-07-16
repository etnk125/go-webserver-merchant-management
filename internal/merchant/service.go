package merchant

import (
	"github.com/etnk125/go-webserver-merchant-management/model"
	"github.com/google/uuid"
)

type MerchantService struct {
	repo merchantRepo
}

type merchantRepo interface {
	GetDefaultCredential() *model.Credential
	CreateMerchant(merchant *model.Merchant) (*model.Merchant, error)
	GetMerchantInfo(id string) (*model.Merchant, error)
}

func NewMerchantService(repo merchantRepo) *MerchantService {
	return &MerchantService{
		repo: repo,
	}
}

func (s *MerchantService) RegisterMerchant(req *model.RegisterMerchantRequest) (*model.Merchant, error) {
	uuid := uuid.New().String()

	defaultCredential := s.repo.GetDefaultCredential()

	merchant := &model.Merchant{
		ID:          uuid,
		Name:        req.Name,
		BankAccount: req.BankAccount,
		Username:    defaultCredential.Username,
		Password:    defaultCredential.Password,
	}

	return s.repo.CreateMerchant(merchant)
}

func (s *MerchantService) GetMerchantInfo(id string) (*model.Merchant, error) {
	return s.repo.GetMerchantInfo(id)
}

func (s *MerchantService) UpdateMerchantInfo(id string, req *model.UpdateMerchantRequest) (*model.Merchant, error) {
	merchant, err := s.repo.GetMerchantInfo(id)
	if err != nil {
		return nil, err
	}

	merchant, err = s.repo.UpdateMerchantInfo(merchant, req)
	if err != nil {
		return nil, err
	}

}
