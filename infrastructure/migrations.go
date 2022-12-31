package infrastructure

import "github.com/cledupe/jwt-auth/models"

func MigrateTables() {
	err := DB.Db.AutoMigrate(models.User{})
	if err != nil {
		panic(err)
	}
}
