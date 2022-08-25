package seeders

import (
	"forum/pkg/seed"
)

func Initialize() {
	seed.SetRunOrder([]string{
		"SeedUsersTable",
	})
}
