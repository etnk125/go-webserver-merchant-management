package merchant

import "github.com/etnk125/go-webserver-merchant-management/model"

type MerchantRepository struct {
	model.MemoryDB
}

func NewMerchantRepository() *MerchantRepository {
	return &MerchantRepository{
		MemoryDB: model.MemoryDB{
			Merchants:         make(map[string]*model.Merchant),
			Products:          make(map[string]*model.Product),
			Transactions:      make(map[string]*map[string]*[]model.Transaction),
			DefaultCredential: *DefaultCredential(),
		},
	}
}
func (r *MerchantRepository) GetDefaultCredential() *model.Credential {
	return &r.DefaultCredential
}
func (r *MerchantRepository) CreateMerchant(merchant *model.Merchant) (*model.Merchant, error) {
	r.Merchants[merchant.ID] = merchant
	return merchant, nil
}
func (r *MerchantRepository) GetMerchantInfo(id string) (*model.Merchant, error) {
	merchant, ok := r.Merchants[id]
	if !ok {
		return nil, ErrMerchantNotFound()
	}
	return merchant, nil
}
func (r *MerchantRepository) UpdateMerchantInfo(id string, req *model.UpdateMerchantRequest) (*model.Merchant, error) {
	merchant, err := r.GetMerchantInfo(id)
	if err != nil {
		return nil, err
	}

	if req.Name != "" {
		merchant.Name = req.Name
	}

	if req.BankAccount != "" {
		merchant.BankAccount = req.BankAccount
	}

	if req.Username != "" {
		merchant.Username = req.Username
	}

	if req.Password != "" {
		merchant.Password = req.Password
	}

	return merchant, nil
}
func (r *MerchantRepository) AddProduct(merchantID string, product *model.Product) (*model.Product, error) {
	merchant, err := r.GetMerchantInfo(merchantID)
	if err != nil {
		return nil, err
	}

	product.MerchantID = merchant.ID

	r.Products[product.ID] = product

	return product, nil
}
