package initdata

import (
	"YesSearch/leveldbOp"
	"YesSearch/tokenizer"
	"encoding/csv"
	"io"
	"log"
	"os"
)

var datafile string

func init() {
	datafile = "configs/src_cache/wukong_100m_0.csv"
}


func Read_data() {
	df, err := os.Open(datafile)
	if err != nil {
		log.Fatalf("read failed, err is %+v", err)
	}
	defer df.Close()
	//数据集较大，一行一行读取
	reader := csv.NewReader(df)

	//越过第一行
	reader.Read() 
	if err != nil{
		log.Fatalf("row read failed, err is %+v", err)
	}

	for {
		row, err := reader.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("row read failed, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		words := tokenizer.Cut_word(row[1])
		for _, word := range words {
			if !leveldbOp.Ishan(word) {
				continue
			}
			//fmt.Println(word)
			leveldbOp.InsertWordUrl(word, row[0], row[1])
		}
	}
}

func Read_data_cache() {
	df, err := os.Open(datafile)
	if err != nil {
		log.Fatalf("read failed, err is %+v", err)
	}
	defer df.Close()
	//数据集较大，一行一行读取
	reader := csv.NewReader(df)

	//越过第一行
	reader.Read() 
	if err != nil{
		log.Fatalf("row read failed, err is %+v", err)
	}

	for {
		row, err := reader.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("row read failed, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		words := tokenizer.Cut_word(row[1])
		for _, word := range words {
			if !leveldbOp.Ishan(word) {
				continue
			}
			//fmt.Println(word)
			leveldbOp.InsertWordUrlcache(word, row[0], row[1])
		}
	}
}
