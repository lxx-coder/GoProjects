package redisTest

import (
    "fmt"
    "github.com/gomodule/redigo/redis"
)

const (
    REDISKEY = "mobile:p30:n3000:1"
    Ip       = "127.0.0.1"
    Port     = "6379"
)


func lpush() {
    conn, err := redis.Dial("tcp",Ip + ":" + Port)
    if err != nil {
        fmt.Printf("connect to redis failed, err: %v\n", err)
        return
    }

    defer conn.Close()

    fmt.Println("connect redis success")

    keys, err := conn.Do("keys", REDISKEY)
    if err != nil {
        fmt.Printf("get key error: %v, key: %v", err, keys)
    }

    if keys, ok := keys.([]interface{}); ok {
        if len(keys) <= 0 {
            _, err = conn.Do("lpush",REDISKEY, 1,2,3,4,5,6,7,8,9,10)
            if err != nil {
                fmt.Printf("lpush to redis failed, err: %v\n", err)
            }
        }
    }

    item, err := redis.String(conn.Do("lpop", REDISKEY))
    if err != nil {
        fmt.Printf("lpop redis failed, err: %v\n", err)
        return
    }

    fmt.Printf("lpop from redis success, item: %v\n", item)

    //if keys, ok := keys.([]interface{});ok {
    //    for _, key := range keys {
    //        switch key := key.(type) {
    //        case []byte:
    //            fmt.Printf("key: %v\n", string(key))
    //        }
    //    }
    //}

}
