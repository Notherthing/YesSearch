package leveldbOp

import (
	"YesSearch/basic"
	"YesSearch/rbmap"
	"log"
	"unicode"
	"os"
	"github.com/hyperjumptech/beda"
	"github.com/syndtr/goleveldb/leveldb"
)

//每次根据word，url，caption，初始化进cache
func InsertWordUrlcache(word string, url string, caption string) error{
	dbPath := getDbPathcache(word)
	//生成path
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		log.Fatalf("open dbfile failed, dbPath is %+v, err is %+v", 
		dbPath, err)
	}
	defer db.Close()
	err = db.Put([]byte(url), makeValue(caption), nil)
	//db当中，key为url的[]byte类型，value为caption的[]byte类型
	if err != nil {
		log.Fatalf("db put failed, dbpath is %+v, url is %+v",
		dbPath, err)	
	}
	return nil
}

//每次根据word，url，caption,初始化进主数据库
func InsertWordUrl(word string, url string, caption string) error{
	dbPath := getDbPath(word)
	//生成path
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		log.Fatalf("open dbfile failed, dbPath is %+v, err is %+v", 
		dbPath, err)
	}
	defer db.Close()
	err = db.Put([]byte(url), makeValue(caption), nil)
	//db当中，key为url的[]byte类型，value为caption的[]byte类型
	if err != nil {
		log.Fatalf("db put failed, dbpath is %+v, url is %+v",
		dbPath, err)	
	}
	return nil
}

//把一个word对应的所有url读入指定rbmap,from main
func ReadWordUrlMain(word string, mp rbmap.Map, searchCaption string) {
	dbPath := getDbPath(word)
	//生成path
	_, err:= os.Stat(dbPath)
	if err != nil {
		return
	}
	//如果无结果则直接返回
	
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		log.Fatalf("open dbfile failed, dbPath is %+v, err is %+v", 
		dbPath, err)
	}
	defer db.Close()
	//将一个word_db里的url与caption全部打入rbmap
	//rbmap需要以相似度进行排序
	//key 以一个float32的相似度和一个剩余url进行排序，value为caption
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := basic.BytesCombine(
			basic.Float32ToByte(beda.NewStringDiff(searchCaption, 
			string(iter.Value())).JaroWinklerDistance(0.1)),
			iter.Key() )
		mp.Put(string(key), iter.Value())
	}
}

//把一个word对应的所有url读入指定rbmap,from cache
func ReadWordUrlCache(word string, mp rbmap.Map, searchCaption string) {
	dbPath := getDbPathcache(word)
	//生成path
	_, err:= os.Stat(dbPath)
	if err != nil {
		return
	}
	//如果无结果则直接返回
	
	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		log.Fatalf("open dbfile failed, dbPath is %+v, err is %+v", 
		dbPath, err)
	}
	defer db.Close()
	//将一个word_db里的url与caption全部打入rbmap
	//rbmap需要以相似度进行排序
	//key 以一个float32的相似度和一个剩余url进行排序，value为caption
	iter := db.NewIterator(nil, nil)
	for iter.Next() {
		key := basic.BytesCombine(
			basic.Float32ToByte(beda.NewStringDiff(searchCaption, 
			string(iter.Value())).JaroWinklerDistance(0.1)),
			iter.Key() )
		mp.Put(string(key), iter.Value())
	}
}

func getDbPath(word string) string {
	path := "datadbmain/" + word + "/db"
	return path
}

func getDbPathcache(word string) string {
	path := "datadbcache/" + word + "/db"
	return path
}

func makeValue(dsc string) []byte{
	//直接将 caption 转为 bytes
	return basic.String2Bytes(dsc)
}

func Ishan(word string) bool {
	for _, v := range word {
		if unicode.Is(unicode.Han, v) {
			return true
		}
	}
	return false
}