package dbs

import (
	"log"

	"github.com/go-redis/redis"
)

// Redis Connect function
// this function is mainly use to connect to Redis database
// return err indicate the success or failure function execute
func (db *TRedis) Connect() (err error) {
	// create one new redis client
	db.DB = redis.NewClient(&redis.Options{
		Addr:     db.Host + ":" + db.Port,
		Password: db.Password,
		DB:       0,
	})
	// try to ping redis server
	err = db.ping()
	if err != nil {
		db.DB.Close()
		log.Println("Error connect to redis server.")
		return err
	}
	return err
}

// Redis Get function
// this function is mainly use to get a value though a key
// return err indicate the success or failure function execute
func (db *TRedis) Get(key string) (value string, err error) {
	value, err = db.DB.Get(key).Result()
	if err == redis.Nil {
		log.Printf("Error get redis key(%v):key do not exist.", key)
	} else if err != nil {
		log.Printf("Error get redis key(%v):%v", key, err)
	} else {
		log.Printf("Success get redis key(%v) value(%v).", key, value)
	}
	return value, err
}

// Redis Set function
// this function is mainly use to set a value though a key
// return err indicate the success or failure function execute
func (db *TRedis) Set(key string, value interface{}) (err error) {
	err = db.DB.Set(key, value, 0).Err()
	if err != nil {
		log.Printf("Error set redis key(%v) value(%v):%v", key, value, err)
		return err
	}
	return err
}

// Redis ping function
// this function is mainly use to test whether database exist or not
// return err indicate the success or failure function execute
func (db *TRedis) ping() (err error) {
	pong, err := db.DB.Ping().Result()
	if err != nil {
		log.Println("Error ping redis server:", err)
		return err
	}
	log.Println("Ping redis server:", pong)
	return err
}
