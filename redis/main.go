package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool  //创建redis连接池

func string_op()  {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	defer c.Close()
	_, err = c.Do("Set", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(c.Do("Get", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	fmt.Println(r)
}

func mstring_op()  {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	defer c.Close()
	_, err = c.Do("MSet", "abc", 100, "efg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Ints(c.Do("MGet", "abc", "efg"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	for _, v := range r {
		fmt.Println(v)
	}
}

func expire_op()  {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	defer c.Close()
	_, err = c.Do("expire", "abc", 10)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func list_op()  {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	defer c.Close()
	_, err = c.Do("lpush", "book_list", "abc", "ceg", 300)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.String(c.Do("lpop", "book_list"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}

	fmt.Println(r)
}

func hash_op()  {
	c, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		fmt.Println("conn redis failed,", err)
		return
	}

	defer c.Close()
	_, err = c.Do("HSet", "books", "abc", 100)
	if err != nil {
		fmt.Println(err)
		return
	}

	r, err := redis.Int(c.Do("HGet", "books", "abc"))
	if err != nil {
		fmt.Println("get abc failed,", err)
		return
	}
	fmt.Println(r)
}

func init()  {
	pool = &redis.Pool{     //实例化一个连接池
		MaxIdle:16,    //最初的连接数量
		// MaxActive:1000000,    //最大连接数量
		MaxActive:0,    //连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout:300,    //连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn ,error){     //要连接的redis数据库
			return redis.Dial("tcp","localhost:6379")
		},
	}
}

func pool_op()  {
	c := pool.Get() //从连接池，取一个链接
	defer c.Close() //函数运行结束 ，把连接放回连接池

	_,err := c.Do("Set","abc",200)
	if err != nil {
		fmt.Println(err)
		return
	}

	r,err := redis.Int(c.Do("Get","abc"))
	if err != nil {
		fmt.Println("get abc faild :",err)
		return
	}
	fmt.Println(r)
	pool.Close() //关闭连接池
}

func main() {
	//mstring_op()
	//expire_op()
	//list_op()
	//hash_op()
	pool_op()
}
