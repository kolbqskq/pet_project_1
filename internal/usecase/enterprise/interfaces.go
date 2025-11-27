package enterprise

type EnterpriseManager interface {
	ValidateEndGame() error
}

type BalanceManager interface {
	GetBalance() int
}
