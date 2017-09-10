package alfredV3

import (
	"bytes"
	"errors"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	plist "github.com/DHowett/go-plist"
	"github.com/boltdb/bolt"
)

const (
	//CacheDir ...
	CacheDir = "Library/Caches/com.runningwithcrayons.Alfred-3/Workflow Data/"
	//InfoFile ...
	InfoFile = "info.plist"
	//
	defaultExpireTime = 24 * time.Hour
)

var (
	cachePath string
)

//InfoPlist defines the plist data format
type InfoPlist struct {
	Bundleid    string `plist:"bundleid"`
	Category    string `plist:"category"`
	Name        string `plist:"name"`
	Description string `plist:"description"`
	Author      string `plist:"createdby"`
	Webaddress  string `plist:"webaddress"`
}

//Cache defines cache object
type Cache struct {
	Expire int64
	Res    []string
	Input  string
}

func init() {
	rootPath, err := os.Getwd()
	if err != nil {
		log.Fatalln(err)
	}
	rootInfoPath := path.Join(rootPath, InfoFile)
	if _, err = os.Stat(rootInfoPath); os.IsNotExist(err) {
		log.Fatalln(err)
	}
	info, err := ioutil.ReadFile(rootInfoPath)
	if err != nil {
		log.Fatalln(err)
	}
	var infoData InfoPlist
	infoBuf := bytes.NewReader(info)
	decoder := plist.NewDecoder(infoBuf)
	err = decoder.Decode(&infoData)
	if err != nil {
		log.Fatalln(err)
	}
	cachePath = os.Getenv("alfred_workflow_cache")
	if cachePath == "" {
		cachePath = path.Join(rootPath, CacheDir, infoData.Bundleid)
	}
	//init db
	initDB()
}

func initDB() {
	db, err := bolt.Open(".cache", 0600, nil)
	if err != nil {
		log.Fatalln(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("alfredItems"))
		if err != nil {
			log.Fatalln(err)
			return err
		}
		return nil
	})

	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte("expire"))
		if err != nil {
			log.Fatalln(err)
			return err
		}
		return nil
	})

}

//StoreData store data
func StoreData(cache *Cache) error {
	if cache.Input == "" {
		return errors.New("invalid input")
	}
	db, err := bolt.Open(".cache", 0600, nil)
	if err != nil {
		log.Fatalln(err)
	}
	db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("alfredItems"))
		err := b.Put([]byte(cache.Input), []byte(strings.Join(cache.Res, ",")))
		return err
	})

	db.Update(func(tx *bolt.Tx) error {
		if cache.Expire == 0 {
			//use default
			cache.Expire = int64(defaultExpireTime)
		}
		b := tx.Bucket([]byte("expire"))
		err := b.Put([]byte(cache.Input), []byte(strconv.FormatInt(cache.Expire, 10)))
		return err
	})
	return nil
}

//FetchData get the cache by the key
func FetchData(key string) *Cache {
	cache := new(Cache)
	cache.Input = key
	db, err := bolt.Open(".cache", 0600, nil)
	if err != nil {
		log.Fatalln(err)
	}

	db.View(func(tx *bolt.Tx) error {
		bk := tx.Bucket([]byte("expire"))
		var err error
		cache.Expire, err = strconv.ParseInt(string(bk.Get([]byte(key))), 10, 64)
		if err != nil {
			return err
		}
		if cache.Expire < time.Now().Unix() {
			cache.Res = []string{}
			return nil
		}
		b := tx.Bucket([]byte("alfredItems"))
		cache.Res = strings.SplitN(string(b.Get([]byte(key))), ",", -1)
		return nil
	})

	return cache
}
