package repository

import (
	"abc/model"
	"database/sql"
	"log"
)

type IRepositoryPort interface {
	// Port ให้คนอื่นมาเรียกใช้ return ค่าเป็น slice
	GetRep() ([]model.GetRequest, error)
	GetRepById(id string) (model.GetRequest, error)
}

type repositoryAdapter struct {
	// รับข้อมูล จาก NewRepositoryAdapter เพื่อนำไปใช้งานต่อ
	db *sql.DB
}

func NewRepositoryAdapter(db *sql.DB) *repositoryAdapter{
	// รับค่าจาก main มาใช้งาน
	return &repositoryAdapter{db: db}
}

func (r repositoryAdapter) GetRep() ([]model.GetRequest, error) {
	// รับค่าจาก repositoryAdapter มาใช้งานโดยตั้งชื่อเป็น r ดึงข้อมูลมาจาก sql
	query := "SELECT id, name, type, detail, url FROM beer"
	rows, err := r.db.Query(query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	// ถ้าทำงานสำเร็จ ให้ปิดการทำงาน
	defer rows.Close()

	// สร้างตัวแปร beers เพื่อเก็บข้อมูลที่ได้จาก sql
	var beers []model.GetRequest

	// วนลูปเพื่อดึงข้อมูลที่ได้จาก sql
	for rows.Next() {
		// สร้างตัวแปร beer เพื่อเก็บข้อมูลที่ได้จาก sql
		var beer model.GetRequest
		err := rows.Scan(&beer.ID, &beer.Name, &beer.Type, &beer.Detail, &beer.URL)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		// นำข้อมูลที่ได้จาก sql มาเก็บไว้ใน beers
		beers = append(beers, beer)
	}

	return beers, nil
}

func (r repositoryAdapter) GetRepById(id string) (model.GetRequest, error) {
	// รับค่าจาก repositoryAdapter มาใช้งานโดยตั้งชื่อเป็น r ดึงข้อมูลมาจาก sql
	query := "SELECT id, name, type, detail, url FROM beer WHERE id = ?"
	row := r.db.QueryRow(query, id)

	// สร้างตัวแปร beer เพื่อเก็บข้อมูลที่ได้จาก sql
	var beer model.GetRequest
	err := row.Scan(&beer.ID, &beer.Name, &beer.Type, &beer.Detail, &beer.URL)
	if err != nil {
		log.Println(err)
		return model.GetRequest{}, err
	}

	return beer, nil
}