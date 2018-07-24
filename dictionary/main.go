package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/boltdb/bolt"
)

func main() {
	// Open the my.db data file in your current directory.
	// It will be created if it doesn't exist.
	fmt.Println("Parsing dictionary file and bootstrapping BoltDB")
	bootstrapDb()
}

type DictEntry struct {
	Traditional  string
	Pinyin       string
	Translations []string
}

func bootstrapDb() {
	db, err := bolt.Open("dictionary/dict.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	file, err := os.Open("dictionary/dict.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	db.Update(func(tx *bolt.Tx) error {
		bkt, err := tx.CreateBucketIfNotExists([]byte("dict"))
		if err != nil {
			return err
		}
		for scanner.Scan() {
			line := scanner.Text()
			if line[:1] != "#" {
				simplified, traditional, pinyin, translations := parseLine(line)
				json, jsonerr := json.Marshal(DictEntry{traditional, pinyin, translations})
				if jsonerr != nil {
					log.Fatal(jsonerr)
				}
				err = bkt.Put([]byte(simplified), json)

			}
		}
		return err
	})
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Sueccessfully created dictionary file")
}

func parseLine(line string) (traditional string, simplified string, pinyin string, translations []string) {
	//fmt.Println(line)
	simplified = strings.Split(line, " ")[0]
	traditional = strings.Split(line, " ")[1]

	pinyin = strings.Split(strings.Split(line, "[")[1], "]")[0]
	translations = strings.Split(strings.Split(line, "]")[1], "/")
	translations = translations[1 : len(translations)-1]
	return
}
