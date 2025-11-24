package enterprise

type EnterpriseWalletDeps struct {
	BalanceManager BallanceManager
}

func GetBalance(deps EnterpriseWalletDeps) int {
	return deps.BalanceManager.GetBallance()
}
