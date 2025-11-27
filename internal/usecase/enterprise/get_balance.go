package enterprise

type EnterpriseWalletDeps struct {
	BalanceManager BalanceManager
}

func GetBalance(deps EnterpriseWalletDeps) int {
	return deps.BalanceManager.GetBalance()
}
