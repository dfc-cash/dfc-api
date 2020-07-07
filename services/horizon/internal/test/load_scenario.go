package test

import (
	"github.com/dfc/go/services/horizon/internal/test/scenarios"
)

func loadScenario(scenarioName string, includeHorizon bool) {
	dfcCorePath := scenarioName + "-core.sql"
	horizonPath := scenarioName + "-horizon.sql"

	if !includeHorizon {
		horizonPath = "blank-horizon.sql"
	}

	scenarios.Load(StellarCoreDatabaseURL(), dfcCorePath)
	scenarios.Load(DatabaseURL(), horizonPath)
}
