package main

import "strconv"

// 可选参数
type params struct {
	clientID    int
	countryCode string
}

// 构造选项
type Option func(params *params)

func GetOptions(opts ...Option) string {
	var r params
	for _, opt := range opts {
		opt(&r)
	}
	return structToString(r)
}

func SetClientID(ID int) Option {
	return func(params *params) {
		params.clientID = ID
	}
}

func SetCountryCode(code string) Option {
	return func(params *params) {
		params.countryCode = code
	}
}
func structToString(params params) string {
	return "clientID:" + strconv.Itoa(params.clientID) + ",countryCode:" + params.countryCode
}
