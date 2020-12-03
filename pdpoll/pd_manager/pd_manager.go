package pg_manager

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

type MyDB struct {
	pool chan *sql.DB
}

var ErrPoolClosed = errors.New("连接池已经关闭！")

const (
	host     = "192.168.4.31"
	port     = 5432
	user     = "postgres"
	password = "123456"
	dbname   = "airdb"
)

func (md *MyDB) New(size int) {
	md.pool = make(chan *sql.DB, size)
}

func (md *MyDB) Close() {

	close(md.pool)
	for db := range md.pool {
		db.Close()
	}
	log.Println("Close:", "资源回收成功！")

}

func (md *MyDB) createDBConn() (*sql.DB) {
	pgSqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", pgSqlInfo)

	if err != nil {
		log.Println("DB Open:", err)
		return nil
	}
	err = db.Ping()
	if err != nil {
		log.Println("DB PING: ", err)
		return nil
	}
	return db
}

func (md *MyDB) getDBConn() (*sql.DB, error) {
	select {
	case r, ok := <-md.pool:
		log.Println("getDBConn:", "从连接池中获取资源...")
		if !ok {
			return nil, ErrPoolClosed
		}
		return r, nil
	default:
		log.Println("getDBConn:", "创建新的资源...")
		return md.createDBConn(), nil
	}
}

func (md *MyDB) putDBConn(db *sql.DB) {
	select {
	case md.pool <- db:
		log.Println("putDBConn", "连接放入池中...")
	default:
		log.Println("putDBConn", "队列已满,关闭当前连接...")
		db.Close()
	}
}

func (md *MyDB) ReadMany(readSql string) (r []string) {
	db, err := md.getDBConn()
	if err != nil {
		log.Println(err)
	}
	md.putDBConn(db)
	var result []string
	rows, err := db.Query(readSql)
	if err != nil {
		log.Println(err)
	}

	for rows.Next() {
		var image string
		err := rows.Scan(&image)
		if err != nil {
			log.Println("Read Many:", err)
		}
		result = append(result, image)
	}
	return result
}

func (md *MyDB) ReadOne(readSql string) (r string) {

	db, err := md.getDBConn()
	if err != nil {
		log.Println(err)
	}
	defer md.putDBConn(db)
	row := db.QueryRow(readSql, 1)

	var image string
	err = row.Scan(&image)
	if err != nil {
		log.Println("Read Many:", err)
	}
	return image
}

