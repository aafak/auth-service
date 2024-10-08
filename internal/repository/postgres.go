package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/aafak/auth-service/internal/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgresDB(dbUrl string) (*gorm.DB, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in NewPostgresDB", r)
		}
	}()
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Build the PostgreSQL connection string
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=UTC",
		dbHost, dbUser, dbPassword, dbName, dbPort)
	fmt.Println(dbUrl)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// log.Fatal will exit, and `defer func(){...}(...)` will not run
		//log.Fatal("Failed to connect to PostgreSQL database:", err)
		return nil, err
	}

	err = db.AutoMigrate(&model.User{})
	if err != nil {
		log.Fatal("Failed to migrate database schema:", err)
	}
	return db, err
}

/*
aafak@aafak-virtual-machine:~$ sudo -u postgres psql
[sudo] password for aafak:
could not change directory to "/home/aafak": Permission denied
psql (14.13 (Ubuntu 14.13-0ubuntu0.22.04.1))
Type "help" for help.

postgres=# \d
Did not find any relations.
postgres=# \l
                             List of databases
   Name    |  Owner   | Encoding | Collate | Ctype |   Access privileges
-----------+----------+----------+---------+-------+-----------------------
 postgres  | postgres | UTF8     | en_IN   | en_IN |
 template0 | postgres | UTF8     | en_IN   | en_IN | =c/postgres          +
           |          |          |         |       | postgres=CTc/postgres
 template1 | postgres | UTF8     | en_IN   | en_IN | =c/postgres          +
           |          |          |         |       | postgres=CTc/postgres
 testdb    | postgres | UTF8     | en_IN   | en_IN | =Tc/postgres         +
           |          |          |         |       | postgres=CTc/postgres+
           |          |          |         |       | aafak2=CTc/postgres
(4 rows)

postgres=#
aafak@aafak-virtual-machine:~$ su aafak2
Password:

aafak2@aafak-virtual-machine:/home/aafak$ psql -U aafak2 -d testdb
could not change directory to "/home/aafak": Permission denied
psql (14.13 (Ubuntu 14.13-0ubuntu0.22.04.1))
Type "help" for help.

testdb=#
testdb=# \d
             List of relations
 Schema |     Name     |   Type   | Owner
--------+--------------+----------+--------
 public | table1       | table    | aafak2
 public | users        | table    | aafak2
 public | users_id_seq | sequence | aafak2
(3 rows)


testdb=# select * from users;
 id | username | password
----+----------+----------
  1 | aafak    | test
(1 row)

testdb=#

*/
