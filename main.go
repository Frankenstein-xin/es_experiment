package main

import (
	"es_experiment/utils"
	"log"
	"os"
	"path/filepath"
	"time"
)

func init() {
	utils.MustInitConfig()
	utils.MustInitESClient(&utils.GetConfig().ESConfig)
	utils.MustInitMysqlClient(&utils.GetConfig().MysqlConfig)
}

func timeCost(start time.Time) {
	tc := time.Since(start)
	log.Printf("Total time cost: %v\n", tc)
}

func main() {

	defer timeCost(time.Now())

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

	utils.MysqlQuery(utils.GetMysqlCli(), bagIds)
}
