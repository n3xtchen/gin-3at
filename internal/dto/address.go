package dto

// CreateAddressRequest represents the request to create a new address.
type CreateAddressReq struct {
	Name    string `json:"name" binding:"required,max=20"`    // Name of the person
	Phone   string `json:"phone" binding:"required,len=11"`   // Phone number
	Address string `json:"address" binding:"required,max=50"` // Full address
}
