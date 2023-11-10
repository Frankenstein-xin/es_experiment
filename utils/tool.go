package utils

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"es_experiment/biz/model"
	"gorm.io/gorm"
	"io"
	"io/fs"
	"log"
	"os"
	"time"

	es7 "github.com/elastic/go-elasticsearch/v7"
)

const (
	IndexRawBag = "raw_bag"
)

func FileJsonRead(srcName string, mode fs.FileMode) []byte {
	fp, err := os.OpenFile(srcName, os.O_RDONLY, mode)
	if err != nil {
		log.Fatal(err)
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)
	var byteData bytes.Buffer
	for scanner.Scan() {
		line := scanner.Bytes()
		byteData.Write(line)
	}
	return byteData.Bytes()
}

func ESQuery(cli *es7.Client, query io.Reader) map[string]interface{} {

	res, err := cli.Search(
		cli.Search.WithContext(context.Background()),
		cli.Search.WithIndex(IndexRawBag),
		cli.Search.WithBody(query),
		cli.Search.WithTrackTotalHits(true),
		cli.Search.WithPretty(),
	)

	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

	if res.IsError() {
		var e map[string]interface{}
		if err := json.NewDecoder(res.Body).Decode(&e); err != nil {
			log.Fatalf("Error parsing the response body: %s", err)
		} else {
			// Print the response status and error information.
			log.Fatalf("[%s] %s: %s",
				res.Status(),
				e["error"].(map[string]interface{})["type"],
				e["error"].(map[string]interface{})["reason"],
			)
		}
	}

	var result map[string]interface{}

	if err := json.NewDecoder(res.Body).Decode(&result); err != nil {
		log.Fatalf("Error parsing the response body: %v", err)
	}
	// Print the response status, number of results, and request duration.
	log.Printf(
		"[%s] %d ES hits; took: %dms",
		res.Status(),
		int(result["hits"].(map[string]interface{})["total"].(map[string]interface{})["value"].(float64)),
		int(result["took"].(float64)),
	)

	return result
}

func MysqlQuery(cli *gorm.DB, fileIds []string) []*model.RawBagStruct {

	start := time.Now()

	rawBags := make([]*model.RawBagStruct, 0)
	err := cli.Where("file_id in (?) ", fileIds).Find(&rawBags).Error
	if err != nil {
		log.Fatalf("MysqlQuery err %v", err)
	}

	tc := time.Since(start)

	log.Printf("Mysql query take: %v with %d records", tc, len(rawBags))

	return rawBags
}
