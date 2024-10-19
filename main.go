package main

import (
	"database/sql"

	"malstat/scrapper/cmd"

	// imports as package "cli"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Build is the last GIT commit
var Build string

func db(connectionString string) (*gorm.DB, error) {
	sqlDB, err := sql.Open("pgx", connectionString)
	if err != nil {
		return nil, err
	}
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return gormDB, nil
}

func main() {
	cmd.Release.Build = Build
	cmd.Run()

	// err := godotenv.Load()
	// if err != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	// connStr := os.Getenv("DATABASE_URL")
	// db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	// log.Print("Connected to database")
	// conn, err := pgx.Connect(context.Background(), connStr)
	// if err != nil {

	// 	panic(err)
	// }
	// defer conn.Close(context.Background())
	// _, err = conn.Exec(context.Background(), "CREATE TABLE IF NOT EXISTS playing_with_neon(id SERIAL PRIMARY KEY, name TEXT NOT NULL, value REAL);")
	// if err != nil {
	// 	panic(err)
	// }
	// _, err = conn.Exec(context.Background(), "INSERT INTO playing_with_neon(name, value) SELECT LEFT(md5(i::TEXT), 10), random() FROM generate_series(1, 10) s(i);")
	// if err != nil {
	// 	panic(err)
	// }
	// rows, err := conn.Query(context.Background(), "SELECT * FROM playing_with_neon")
	// if err != nil {
	// 	panic(err)
	// }
	// defer rows.Close()
	// for rows.Next() {
	// 	var id int32
	// 	var name string
	// 	var value float32
	// 	if err := rows.Scan(&id, &name, &value); err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Printf("%d | %s | %f\n", id, name, value)
	// }
}
