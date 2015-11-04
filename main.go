package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	MysqlDb *sql.DB
)

func main() {
	InitDatabase();
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/addmedicine", AddMedicine)
	router.HandleFunc("/querymedicine", Querymedicine)
	router.HandleFunc("/", Index)
	log.Fatal(http.ListenAndServe(":8002", router))
}

func AddMedicine(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte(fmt.Sprintf("host: %s ID：%s name: %s", req.Host, req.FormValue("id"), req.FormValue("name"))))
}


func Querymedicine(w http.ResponseWriter, req *http.Request) {
	m := &Medicine{Id: "123123123",
		Name:                    "黄芪",
		Alias:                   "黄耆、木耆、绵黄芪",
		Where:                   "豆科植物蒙古黄芪Astragalus membranaceus Bge. var.mongholicus(Bge.) Hsiao和膜荚黄芪A.membranaceus (Fisch.) Bge.的根。",
		OriginDistribution:      "蒙古黄芪生于向阳草地及山坡;膜荚黄芪生于林缘、灌丛、林间草地及疏林下。分布于黑龙江、吉林、辽宁、河北、内蒙古等地。",
		HarvestingAndProcessing: "野生黄芪春秋两季均可采挖，除净泥土及须根，切去根头，晒至七八成干，按粗细、长短不同分级。栽培黄芪应3年以后采收",

		PlantMorphology:     "蒙古黄芪：多年生草本。主根长而粗壮，条较顺直。奇数羽状复叶，小叶12~18对，小叶片下面被柔毛。总状花序腋生，花冠黄色至淡黄色，雄蕊10枚，二体。荚果膨胀，无毛。膜荚黄芪：小叶6~13对，荚果有毛。",
		MedicinalProperties: "蒙古黄芪：表面灰黄色，栓皮不易脱落。质硬而韧，断面纤维性并显粉性。皮部黄白色，木部淡黄色。气微，味微甜，有豆腥味。膜荚黄芪：表面灰黄色、黄棕色，质硬，较难折断。",
		Taste:               "性微温，味甘。归脾经、肺经。",
		Efficacy:            "补气固表、利尿、托毒排脓、生肌。属补虚药下属分类的补气药。",
		Application:         "用量9~30克，治疗气短心悸、乏力、虚脱、自汗、盗汗、体虚浮肿、慢性肾炎、久泻、脱肛、子宫脱垂、痈疽难溃、疮口久不愈合、小儿支气管哮喘、慢性乙型肝炎、慢性肾炎和病毒性心肌炎。补气宜炙用，止汗、利尿、托毒排脓生肌宜生用。",
		Basis:               "主要含三萜皂苷、黄酮及多糖等成分。水煎剂具调节免疫、抗衰老和抗应激作用。毒性：煎剂LDso为(40±5)克/千克。",
		Taboo:               "表实邪盛、湿阻气滞、肠胃积滞、阴虚阳亢、痈疽初起或溃后热毒尚盛者，均禁服。",
		Processing:          "将炼蜜加适量开水稀释后，加入净黄芪，拌匀，闷透，置锅内，用文火炒至不粘手，取出，放凉(每100千克黄芪，用炼蜜25千克)。",
	}
	b, err := json.Marshal(m)
	if err != nil {
		w.Write([]byte("药剂查询不正确。"))
	} else {
		w.Write(b)
	}

}

func Index(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("中药字典查询接口"))
}

func InitDatabase() {
	db, err := sql.Open("mysql", "root:123456@/db_cms")
	MysqlDb = db
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("连接mysql数据库成功！")
		stmt, err := MysqlDb.Prepare("INSERT INTO `tb_medication`(`id`, `tb_name`, `tb_alias`, `tb_where`, `tb_efficacy`, `tb_plantMorphology`, `tb_originDistribution`) VALUES(?, ?, ?, ?, ?, ?, ?)")
		if err != nil {
			log.Fatal("准备stmt失败", err)
		}
		defer stmt.Close()
		_, err = stmt.Exec("123123123", 
		"黄芪", 
		"黄耆、木耆、绵黄芪", 
		"豆科植物蒙古黄芪Astragalus membranaceus Bge. var.mongholicus(Bge.) Hsiao和膜荚黄芪A.membranaceus (Fisch.) Bge.的根。", 
		"补气固表、利尿、托毒排脓、生肌。属补虚药下属分类的补气药。", "野生黄芪春秋两季均可采挖，除净泥土及须根，切去根头，晒至七八成干，按粗细、长短不同分级。栽培黄芪应3年以后采收",
		"蒙古黄芪生于向阳草地及山坡;膜荚黄芪生于林缘、灌丛、林间草地及疏林下。分布于黑龙江、吉林、辽宁、河北、内蒙古等地。")
		if err != nil {
			log.Println("插入数据失败", err)
		} else {
			log.Println("出入数据成功")
		}
	}
	
}
