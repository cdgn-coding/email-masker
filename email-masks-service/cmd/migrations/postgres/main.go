package main

import "email-masks-service/src/infrastructure/drivers/postgresql"

func main() {
	migrations := postgresql.NewMigrations()
	migrations.Apply()
}
