package helper

type APIResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Meta    any    `json:"meta,omitempty"`
}

type Meta struct {
	TotalItems   int `json:"total_items"`
	CurrentPage  int `json:"current_page"`
	ItemsPerPage int `json:"items_per_page"`
	TotalPages   int `json:"total_pages"`
}

const DefaultItemsPerPage int = 20

func ResponseFormat(code int, message string, data any, meta any) APIResponse {
	return APIResponse{
		Code:    code,
		Message: message,
		Data:    data,
		Meta:    meta,
	}
}

func MetaResponse(currentPage int, totalItems int) Meta {
	return Meta{
		TotalItems:   totalItems,
		CurrentPage:  currentPage,
		ItemsPerPage: DefaultItemsPerPage,
		TotalPages:   (totalItems / DefaultItemsPerPage) + 1,
	}
}
