package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var Db *sql.DB

var err error

type Person struct {
	Name string
	Age int
}

func main()  {
	Db, _ = sql.Open("postgres","user=user dbname=go_template sslmode=disable")
	if err != nil {
		log.Panicln(err)
	}
	defer Db.Close()

	//C
	// cmd := "INSERT INTO persons (name,age) VALUES ($1,$2)"
	// _, err := Db.Exec(cmd, "Nancy",20)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	//R
	cmd := "SELECT * FROM persons where age = $1"
	//QueryEow 1レコード取得
	row := Db.QueryRow(cmd, 20)
	var p Person
	err = row.Scan(&p.Name, &p.Age)
	if err != nil {
		//データがなかったら
		if err == sql.ErrNoRows {
			log.Println("No row")
			//それ以外のエラー
		} else {
			log.Println(err)
		}
	}
	fmt.Println(p.Name, p.Age)

	cmd = "SELECT * FROM persons"
	//Queryは条件に合うものを全て取得
	rows, _ := Db.Query(cmd)
	defer rows.Close()
	//structを作成
	var pp []Person
	//取得したデータをループでスライスに追加 for rows.Next()
	for rows.Next(){
		var p Person
		//scanデータ追加
		err := rows.Scan(&p.Name, &p.Age)
		//一つずつエラーハンドリングver
		if err != nil {
			log.Println(err)
		}
		pp = append(pp,p)
	}
	//まとめてエラーハンドリングver
	err = rows.Err()
	if err != nil {
		log.Fatalln(err)
	}
	//表示
	for _,p := range pp {
		fmt.Println(p.Name, p.Age)
	}
}