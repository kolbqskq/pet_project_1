package enterprise

type EnterpriseManager interface {
	ValidateEndGame() error
}

type BallanceManager interface {
	GetBallance() int
}
