package model

type MemoryDB struct {
	Merchants         map[string]*Merchant
	Products          map[string]*Product
	Transactions      map[string]*map[string]*[]Transaction
	DefaultCredential Credential
}

type Merchant struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	BankAccount string `json:"bank_account"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}

type Product struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Price      int    `json:"price"`
	MerchantID string `json:"merchant_id"`
}

type Transaction struct {
	ID         string `json:"id"`
	ProductID  string `json:"product_id"`
	Quantity   int    `json:"quantity"`
	MerchantID string `json:"merchant_id"`
}

type RegisterMerchantRequest struct {
	Name        string `json:"name"`
	BankAccount string `json:"bank_account"`
}

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdateMerchantRequest struct {
	Name        string `json:"name"`
	BankAccount string `json:"bank_account"`
	Username    string `json:"username"`
	Password    string `json:"password"`
}
