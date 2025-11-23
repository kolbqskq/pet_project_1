package enterprise

import "context"

type EnterpriseManagerDeps struct {
	EndGameManager EnterpriseManager
	Cancel         context.CancelFunc
}

func EndGame(deps EnterpriseManagerDeps) error {
	if err := deps.EndGameManager.ValidateEndGame(); err != nil {
		return err
	}
	deps.Cancel()
	return nil
}
