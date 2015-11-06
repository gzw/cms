package main

import (
	"encoding/json"
)

const (
	CMS_MEDICINE_NO_ERROR             = 0
	CMS_MEDICINE_NOT_FOUND            = 1
	CMS_MEDICINE_PARAM_ERROR          = 2
	CMS_MEDICINE_QUERY_ERROR          = 3
	CMS_MEDICINE_REQUEST_METHOD_ERROR = 4
	CMS_MEDICINE_ADD_BODY_PARSE_ERROR = 5
	CMS_MEDICINE_ADD_PARAM_ERROR      = 6
	CMS_MEDICINE_ADD_INSERT_DB_ERROR  = 7
)

type CmsError struct {
	ErrCode int    `json:errCode`
	ErrDesc string `json:"errDesc"`
}

func NewCmsError(code int, desc string) *CmsError {
	return &CmsError{
		ErrCode: code,
		ErrDesc: desc,
	}
}

func CmsErrorToJsonData(code int, desc string) []byte {
	e := NewCmsError(code, desc)
	b, _ := json.Marshal(e)
	return b
}

func CmsErrorNoErrToJsonData(desc string) []byte {
	return CmsErrorToJsonData(CMS_MEDICINE_NO_ERROR, desc)
}

func (e *CmsError) Error() string {
	b, _ := json.Marshal(e)
	return string(b)
}

func (e *CmsError) CmsErrorToJsonStr() string {
	d, err := json.Marshal(e)
	if err != nil {
		return ""
	}
	return string(d)
}
