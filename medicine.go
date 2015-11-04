/*
* @Author: gzw
* @Date:   2015-11-04 13:26:41
* @Last Modified by:   gzw
* @Last Modified time: 2015-11-04 16:59:32
 */

package main

type Medicine struct {
	Id                      string `json:"id"`
	Name                    string `json:"name"`
	Alias                   string `json:"alias,omitempty"`
	Where                   string `json:"where,omitempty"`
	Efficacy                string `json:"efficacy,omitempty"`                //功效
	PlantMorphology         string `json:"plantMorphology,omitempty"`         // 植物形态
	OriginDistribution      string `json:"originDistribution,omitempty"`      // 产地分布
	HarvestingAndProcessing string `json:"harvestingAndProcessing,omitempty"` // 采收加工
	MedicinalProperties     string `json:"medicinalProperties,omitempty"`     // 药物性状
	Taste                   string `json:"taste,omitempty"`                   // 性味归经
	Application             string `json:"application,omitempty"`             // 临床应用
	Basis                   string `json:"basis,omitempty"`                   // 主要成分
	Taboo                   string `json:"taboo,omitempty"`                   // 使用禁忌
	Processing              string `json:"processing,omitempty"`              // 加工方法
	Other 					string `json:"other,omitempty"`				  // 其他
}
