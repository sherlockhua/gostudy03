package main

import (
	"context"
	"fmt"
	elastic "github.com/olivere/elastic"
)

type Person struct {
	Id int   `json:"id"`
	Name string `json:"name"`
	Age int    `json:"age"`
	City string  `json:"city"`
	Desc string `json:"desc"`
}

func main() {
	client, err := elastic.NewClient(elastic.SetURL("http://192.168.20.200:9200"))
	if err != nil {
		fmt.Printf("new client faile, err:%v\n", err)
		return
	}

	for i := 0; i < 10000; i++ {
		a := &Person{
			Id:i,
			Name:fmt.Sprintf("hello %d", i),
			Age:28,
			Desc: `中新网兰州11月3日电 (记者 魏建军)兰州广通新能源汽车有限公司首批新能源汽车2日在兰州新区投产，该公司董事长兼总经理杨健表示，生产线的正式投产，标志着这个82岁的“汽车老兵”翻开了崭新的一页。
			
			兰州广通新能源汽车有限公司，其前身是甘肃驼铃客车厂，成立于1936年，是甘肃省唯一具有客车生产资质的企业。`,
		}
		_, err :=client.Index().Index("account2").Type("Person").BodyJson(a).Do(context.Background())
		if err != nil {
			fmt.Printf("do insert es failed, err:%v\n", err)
		}
	}
}