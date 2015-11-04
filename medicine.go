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
	Alias                   string `json:"alias"`
	Where                   string `json:"where"`
	Efficacy                string `json:"efficacy"`                //功效
	PlantMorphology         string `json:"plantMorphology"`         // 植物形态
	OriginDistribution      string `json:"originDistribution"`      // 产地分布
	HarvestingAndProcessing string `json:"harvestingAndProcessing"` // 采收加工
	MedicinalProperties     string `json:"medicinalProperties"`     // 药物性状
	Taste                   string `json:"taste"`                   // 性味归经
	Application             string `json:"application"`             // 临床应用
	Basis                   string `json:"basis"`                   // 主要成分
	Taboo                   string `json:"taboo"`                   // 使用禁忌
	Processing              string `json:"processing"`              // 加工方法
}
