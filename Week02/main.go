package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
	"log"
)
type Dao struct {
}


func (dao *Dao) UpdateDb() (int, error) {
	//dao层返回错误
	return 0 ,errors.Wrap(sql.ErrNoRows,"biz error")
}
type Service struct {
	Dao
}

func ServiceNew()  *Service {
	return &Service{}
}

func (svr *Service) Update() (int,error) {
	fmt.Println("service UpdateDb")
	_, err := svr.UpdateDb()
	if err != nil {
		// service层处理
		if errors.Is(errors.Cause(err), sql.ErrNoRows) {
			log.Printf("发生了sql.ErrNoRows错误\n")
			log.Printf("原始错误发生信息：%T %v\n", errors.Cause(err), errors.Cause(err))
			log.Printf("堆栈信息：\n%+v\n", err)
			return 0,err
		}
	}
	return 10,nil; //fake data
}
func main() {
	svr := ServiceNew()
	num, err := svr.Update()
	if err != nil {
		log.Println("HTTP 500")
	}
	log.Println(num)

}
