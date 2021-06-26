package models

// Senha representa o formato da reuisição de alteração de senha
type Senha struct {
	Nova  string `json:"nova"`
	Atual string `json:"atual"`
}
