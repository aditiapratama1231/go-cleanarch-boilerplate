package request

type (
	CreateProductRequest struct {
		Data struct {
			ProductName        string `json:"product_name"`
			ProductDescription string `json:"product_description"`
			Quantity           int    `json:"qty"`
		}
	}
)
