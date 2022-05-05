package model

// response APIレスポンス
type Response struct {
	Results Results `json:"results"`
}

type Results struct {
	Shop []Shop `json:"shop"`
}

type Shop struct {
	Name    string `json:"name"`
	Address string `json"address"`
}
