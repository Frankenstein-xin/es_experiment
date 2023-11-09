package main

import (
	"es_experiment/utils"
	"log"
	"os"
	"path/filepath"
)

func init() {
	utils.MustInitConfig()
	utils.MustInitESClient(&utils.GetConfig().ESConfig)
	utils.MustInitMysqlClient(&utils.GetConfig().MysqlConfig)
}

func main() {
	currDir, _ := os.Getwd()
	queryPath := filepath.Join(currDir, "query.json")

	q, err := os.Open(queryPath)

	if err != nil {
		panic(err)
	}

	result := utils.ESQuery(utils.GetESCli(), q)

	hitsData := result["hits"].(map[string]interface{})["hits"].([]interface{})

	log.Printf("Current ES query len: %d", len(hitsData))

	var bagIds []string

	for _, bag := range hitsData {
		bagId := bag.(map[string]interface{})["_source"].(map[string]interface{})["file_id"].(string)
		bagIds = append(bagIds, bagId)
	}

	bags := utils.MysqlQuery(utils.GetMysqlCli(), bagIds)

	log.Printf("bag size: %d\n", len(bags))
}
