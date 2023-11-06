// +build wireinject

package operation

import (
	"github.com/khulnasoft-lab/vul/pkg/db"
	"github.com/google/wire"
)

func initializeDBClient(cacheDir string, quiet bool) db.Client {
	wire.Build(db.SuperSet)
	return db.Client{}
}
