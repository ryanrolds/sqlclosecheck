package stmt

import (
	"database/sql"
	"math/rand"
	"sync"
)

var (
	globalOnce sync.Once
	globalStmt *sql.Stmt
	globalErr  error
)

func globalVarTest() {
	defer func() {
		if globalStmt != nil {
			_ = globalStmt.Close()
		}
	}()

	for i := 0; i < 100; i++ {
		if rand.Float64() > 0.8 {
			err := optionalDbOp()
			if err != nil {
				return
			}
		}
	}
}

func optionalDbOp() error {
	globalOnce.Do(func() {
		globalStmt, globalErr = db.Prepare("INSERT INTO `test` (`id`) VALUES(1)")
	})
	if globalErr != nil {
		return globalErr
	}

	_, err := globalStmt.Exec()
	return err
}
