package model

// ******** Request ********

type LoginParams struct {
	Email      string `json:"email" binding:"required"`
	VerifyCode string `json:"verifyCode" binding:"required"`
}
