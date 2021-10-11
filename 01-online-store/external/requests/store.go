package requests

type CreateCustomerRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Kota        string `json:"kota"`
}

type CreateSupplierRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Kota        string `json:"kota"`
}

type CreateProductRequest struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	Address     string `json:"address"`
	Kota        string `json:"kota"`
}

type CreateCategoryRequest struct {
	Name string `json:"name"`
}
