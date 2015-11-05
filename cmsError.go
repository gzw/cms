package main

type CmsError struct {
	ErrDesc string `json:"errDesc"`
	ErrCode int	`json:errCode`
}