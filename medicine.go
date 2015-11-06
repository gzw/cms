/*
* @Author: gzw
* @Date:   2015-11-04 13:26:41
* @Last Modified by:   gzw
* @Last Modified time: 2015-11-04 16:59:32
 */

package main

type Medicine struct {
	Id                      string `json:"id" sql:"id"`
	Name                    string `json:"name" sql:"tb_name"`
	Alias                   string `json:"alias,omitempty" sql:"tb_alias"`
	Where                   string `json:"where,omitempty" sql:"tb_where"`
	Efficacy                string `json:"efficacy,omitempty" sql:"tb_efficacy"`                               //功效
	PlantMorphology         string `json:"plantMorphology,omitempty" sql:"tb_plantMorphology"`                 // 植物形态
	OriginDistribution      string `json:"originDistribution,omitempty" sql:"tb_originDistribution"`           // 产地分布
	HarvestingAndProcessing string `json:"harvestingAndProcessing,omitempty" sql:"tb_harvestingAndProcessing"` // 采收加工
	MedicinalProperties     string `json:"medicinalProperties,omitempty" sql:"tb_medicinalProperties"`         // 药物性状
	Taste                   string `json:"taste,omitempty" sql:"tb_taste"`                                     // 性味归经
	Application             string `json:"application,omitempty" sql:"tb_application"`                         // 临床应用
	Basis                   string `json:"basis,omitempty" sql:"tb_basis"`                                     // 主要成分
	Taboo                   string `json:"taboo,omitempty" sql:"tb_taboo"`                                     // 使用禁忌
	Processing              string `json:"processing,omitempty" sql:"tb_processing"`                           // 加工方法
	Other                   string `json:"other,omitempty" sql:"tb_other"`                                     // 其他
	ImageUrl                string `json:"imageUrl,omitempty" sql:"tb_imageurl"`                               // 图像地址
}

func NewMedicine() *Medicine {
	return &Medicine{Id: "123123123",
		Name:                    "",
		Alias:                   "",
		Where:                   "",
		OriginDistribution:      "",
		HarvestingAndProcessing: "",
		PlantMorphology:         "",
		MedicinalProperties:     "",
		Taste:                   "",
		Efficacy:                "",
		Application:             "",
		Basis:                   "",
		Taboo:                   "",
		Processing:              "",
		Other:                   "",
		ImageUrl:                "",
	}
}
