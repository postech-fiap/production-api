package domain

type Status string

const (
	PENDING    Status = "PENDENTE"
	RECEIVED   Status = "RECEBIDO"
	IN_PREPARE Status = "EM_PREPARO"
	DONE       Status = "FINALIZADO"
	FINISHED   Status = "ENTREGUE"
)
