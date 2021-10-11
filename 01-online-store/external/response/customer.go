package response

type CustomerResp struct {
	CustomerCode int64 `json:"customer_code"`
}

type SupplierResp struct {
	SupplierCode int64 `json:"supplier_code"`
}

type Paginate struct {
	Count     int64   `json:"count"`
	TotalPage float64 `json:"totalpage"`
	Page      int64   `json:"page"`
}
