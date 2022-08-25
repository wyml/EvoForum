package seed

import (
	"forum/pkg/console"
	"forum/pkg/database"
	"gorm.io/gorm"
)

// 存放所有 Seeder
var seeders []Seeder

// 按顺序执行的 Seeder 数组
var orderedSeederNames []string

type SeederFunc func(*gorm.DB)

type Seeder struct {
	Func SeederFunc
	Name string
}

func Add(name string, fn SeederFunc) {
	seeders = append(seeders, Seeder{
		Name: name,
		Func: fn,
	})
}

func SetRunOrder(names []string) {
	orderedSeederNames = names
}

// GetSeeder 通过名称来获取 Seeder 对象
func GetSeeder(name string) Seeder {
	for _, sdr := range seeders {
		if name == sdr.Name {
			return sdr
		}
	}
	return Seeder{}
}

func RunAll() {

	executed := make(map[string]string)
	for _, name := range orderedSeederNames {
		sdr := GetSeeder(name)
		if len(sdr.Name) > 0 {
			console.Warning("Running Odered Seeder: " + sdr.Name)
			sdr.Func(database.DB)
			executed[name] = name
		}
	}

	// 再运行剩下的
	for _, sdr := range seeders {
		// 过滤已运行
		if _, ok := executed[sdr.Name]; !ok {
			console.Warning("Running Seeder: " + sdr.Name)
			sdr.Func(database.DB)
		}
	}
}

// RunSeeder 运行单个 Seeder
func RunSeeder(name string) {
	for _, sdr := range seeders {
		if name == sdr.Name {
			sdr.Func(database.DB)
			break
		}
	}
}
