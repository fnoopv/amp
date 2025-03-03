package migrate

import (
	"fmt"
	"io/fs"
	"net/url"

	"github.com/amacneil/dbmate/v2/pkg/dbmate"
	_ "github.com/amacneil/dbmate/v2/pkg/driver/postgres"
	"goyave.dev/goyave/v5/util/errors"
)

// Migrate 运行数据库迁移, 格式: postgres://username:password@127.0.0.1:5432/database_name?sslmode=disable
func Migrate(dsn string, fs fs.FS) error {
	u, err := url.Parse(dsn)
	if err != nil {
		return errors.New(err)
	}

	db := dbmate.New(u)
	db.FS = fs
	db.MigrationsDir = []string{"database/migrations"}
	migrations, err := db.FindMigrations()
	if err != nil {
		return errors.New(err)
	}

	for _, m := range migrations {
		fmt.Println(m.Version, m.FilePath)
	}

	err = db.CreateAndMigrate()

	return errors.New(err)
}
