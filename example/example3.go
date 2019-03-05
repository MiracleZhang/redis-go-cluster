package main

import (
    "log"
    "fmt"
    "time"

    "github.com/MiracleZhang/redis-go-cluster"
)

func main() {
    cluster, err := redis.NewCluster(
	&redis.Options{
	    StartNodes: []string{"159.138.28.39:7001", "159.138.28.39:7002", "159.138.25.20:7003", "159.138.25.20:7004", "159.138.27.102:7005", "159.138.27.102:7006"},
		InternalToOutStartNodes:[]string{"10.250.0.174:159.138.28.39","10.250.0.109:159.138.25.20","10.250.0.78:159.138.27.102"},
		//StartNodes: []string{"159.138.2.170:6379", "159.138.28.39:7001"},
	    ConnTimeout: 50 * time.Millisecond,
	    ReadTimeout: 50 * time.Millisecond,
	    WriteTimeout: 50 * time.Millisecond,
	    KeepAlive: 16,
	    AliveTime: 60 * time.Second,
	    Debug:true,
	})

    if err != nil {
	log.Fatalf("redis.New error: %s", err.Error())
    }

    _, err = cluster.Do("MSET", "myfoo1", "mybar1", "myfoo2", "mybar2", "myfoo3", "mybar3")
    if err != nil {
	log.Fatalf("MSET error: %s", err.Error())
    }

    values, err := redis.Strings(cluster.Do("MGET", "myfoo1", "myfoo5", "myfoo2", "myfoo3", "myfoo4"))
    if err != nil {
	log.Fatalf("MGET error: %s", err.Error())
    }

    for i := range values {
	fmt.Printf("reply[%d]: %s\n", i, values[i])
    }
}
