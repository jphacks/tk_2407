package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:           "./svc/pkg/query", // 出力パス
		ModelPkgPath:      "./svc/pkg/domain",
		FieldNullable:     true,
		FieldWithIndexTag: true,
		FieldWithTypeTag:  true,
		Mode:              gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=user password=password dbname=yosegaki port=5432 sslmode=disable TimeZone=Asia/Tokyo", // data source name, refer https://github.com/jackc/pgx
		PreferSimpleProtocol: true,                                                                                                       // disables implicit prepared statement usage. By default pgx automatically uses the extended protocol
	}), &gorm.Config{})
	if err != nil {
		panic(err)
		return
	}

	g.UseDB(db)
	all := g.GenerateAllTable() // database to table model.
	g.ApplyBasic(all...)

	// Generate the code
	g.Execute()
}
