package types

//Money представляет собой денежную сумму в минимальной денежных еденицах (центы, копейки, дирамы и т.д) .
type Money int64

//Category представляет собой категорию, в уоторой был совершён плтёж (авто, аптеки, рестораны и т.д.) .
type Category string

//Status представляет собой платёж
type Status string

//Предопределённые статусы платежа
const (
	StatusOk Status = "OK"
	StatusFail Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

//Payment представляет информацию о платеже
type Payment struct {
	ID int
	Amount Money 
	Category Category
	Status Status
}