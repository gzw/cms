package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

var (
	MysqlDb *sql.DB
)

func main() {
	InitDatabase()
	defer MysqlDb.Close()
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/addmedicine", AddMedicine)
	router.HandleFunc("/querymedicine", Querymedicine)
	router.HandleFunc("/", Index)
	router.HandleFunc("/login", Login)
	log.Fatal(http.ListenAndServe(":8002", router))
}

func AddMedicine(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		w.Write(CmsErrorToJsonData(CMS_MEDICINE_REQUEST_METHOD_ERROR, "添加中药需要POST方法！"))
		return
	}

	decoder := json.NewDecoder(req.Body)
	var m Medicine
	err := decoder.Decode(&m)
	if err != nil {
		w.Write(CmsErrorToJsonData(CMS_MEDICINE_ADD_BODY_PARSE_ERROR, err.Error()))
		return
	}

	if len(m.Id) == 0 || len(m.Name) == 0 {
		cmserr := NewCmsError(CMS_MEDICINE_PARAM_ERROR, "中药id或者名称为空")
		w.Write([]byte(cmserr.CmsErrorToJsonStr()))
		return
	}

	err = InsertMedicine(&m)
	if err != nil {
		w.Write(CmsErrorToJsonData(CMS_MEDICINE_ADD_INSERT_DB_ERROR, err.Error()))
		return
	}

	w.Write(CmsErrorNoErrToJsonData("插入成功！"))
}

func Querymedicine(w http.ResponseWriter, req *http.Request) {
	name := req.FormValue("name")
	if len(name) == 0 {
		err := &CmsError{
			ErrCode: CMS_MEDICINE_PARAM_ERROR,
			ErrDesc: "name参数不能为空！",
		}
		w.Write([]byte(err.Error()))
		return
	}
	log.Println(name)
	ms, err := QuerymedicineByName(name)
	if err != nil {
		e := &CmsError{
			ErrCode: CMS_MEDICINE_QUERY_ERROR,
			ErrDesc: err.Error(),
		}
		w.Write([]byte(e.Error()))
		return
	}

	data, _ := json.Marshal(ms)
	str := string(data)
	log.Println(str)
	w.Write(data)
	return
}

func Index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("中药字典查询接口"))
}

func Login(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("登录接口"))
}

func InitDatabase() {
	var err error
	MysqlDb, err = sql.Open("mysql", "root:123456@/db_cms")

	if err != nil {
		log.Fatal(err)
	} else {
		MysqlDb.SetMaxIdleConns(10)
		MysqlDb.SetMaxOpenConns(100)
		log.Println("连接mysql数据库成功！")
	}
}

func QuerymedicineByName(name string) ([]*Medicine, error) {
	log.Println(name)
	rows, err := MysqlDb.Query("SELECT * FROM `tb_medication` WHERE `tb_name` like ?", fmt.Sprintf("%%%s%%", name))
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ms := make([]*Medicine, 0, 120)
	for rows.Next() {
		m := NewMedicine()
		cols, err := rows.Columns()
		if err == nil {
			log.Println(cols)
		}
		err = rows.Scan(&m.Id, &m.Name, &m.Alias, &m.Where, &m.Efficacy,
			&m.PlantMorphology, &m.OriginDistribution, &m.HarvestingAndProcessing,
			&m.MedicinalProperties,
			&m.Taste, &m.Application, &m.Basis, &m.Taboo, &m.Processing, &m.Other, &m.ImageUrl)
		log.Println(m)
		ms = append(ms, m)
	}
	return ms, err
}

func InsertMedicine(m *Medicine) error {
	_, err := MysqlDb.Exec("INSERT INTO `tb_medication`(`id`, "+
		"`tb_name`, "+
		"`tb_alias`, "+
		"`tb_where`, "+
		"`tb_efficacy`, "+
		"`tb_plantMorphology`, "+
		"`tb_originDistribution`, "+
		"`tb_harvestingAndProcessing`, "+
		"`tb_medicinalProperties`, "+
		"`tb_taste`, "+
		"`tb_application`, "+
		"`tb_basis`, "+
		"`tb_taboo`, "+
		"`tb_processing`,"+
		"`tb_other`, "+
		"`tb_imageurl`) VALUES(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		m.Id,
		m.Name,
		m.Alias,
		m.Where,
		m.Efficacy,
		m.PlantMorphology,
		m.OriginDistribution,
		m.HarvestingAndProcessing,
		m.MedicinalProperties,
		m.Taste,
		m.Application,
		m.Basis,
		m.Taboo,
		m.Processing,
		m.Other,
		m.ImageUrl)
	if err != nil {
		return err
	}
	return nil
}
