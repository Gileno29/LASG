package models

type Especificacoes struct {
	CampoDeConhecimento string            `json:"campo_de_conhecimento"`
	AreasAplicaveis     map[string]string `json:"areas_aplicaveis"`
}

type Contexto struct {
	FaseID             string         `json:"fase_id"`
	TempoGastoSegundos int            `json:"tempo_gasto_segundos"`
	Pontuacao          int            `json:"pontuacao"`
	Especificacoes     Especificacoes `json:"especificacoes"`
}

type LogEvento struct {
	UserID    string   `json:"user_id"`
	EventType string   `json:"event_type"`
	Timestamp string   `json:"timestamp"`
	Context   Contexto `json:"context"`
}
