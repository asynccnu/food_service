package model

import (
	"strconv"
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/spf13/viper"
)

func AddNewSearchRecord(name string) error {
	exists, err := redis.Bool(RedisDb.Self.Do("exists", now()))
	if err != nil {
		return err
	}
	_, err = RedisDb.Self.Do("zincrby", now(), 1, name)
	if err != nil {
		return err
	}
	if !exists {
		_, err := RedisDb.Self.Do("expireat", now(), nextWeek())
		if err != nil {
			return err
		}
	}
	return nil
}

func now() string {
	now := time.Now().Format("2006-01-02")
	//fmt.Println(now)
	return now
}

func nextWeek() string {
	now := time.Now()
	next := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	result := strconv.FormatInt(next.Unix(), 10)
	return result
}

func nextDay() string {
	now := time.Now()
	next := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, now.Location())
	result := strconv.FormatInt(next.Unix(), 10)
	return result
}

func pastTime(day int) string {
	currentTime := time.Now()
	time1 := strconv.Itoa(24 * day)
	m, _ := time.ParseDuration("-" + time1 + "h")
	result := currentTime.Add(m)
	return result.Format("2006-01-02")
}

func GetHotSearch() ([]string, error) {
	exists, err := redis.Bool(RedisDb.Self.Do("exists", now()+"hot"))
	if err != nil {
		return []string{}, err
	}

	_, err = RedisDb.Self.Do("zunionstore", now()+"hot", 7, pastTime(0), pastTime(1), pastTime(2), pastTime(3), pastTime(4), pastTime(5), pastTime(6))
	if err != nil {
		return []string{}, err
	}

	result, err := redis.Strings(RedisDb.Self.Do("zrevrange", now()+"hot", 0, viper.GetInt("redis.max_num")-1))
	if err != nil {
		return []string{}, err
	}

	if !exists {
		_, err := RedisDb.Self.Do("expireat", now()+"hot", nextDay())
		if err != nil {
			return []string{}, err
		}
	}
	return result, nil
}
