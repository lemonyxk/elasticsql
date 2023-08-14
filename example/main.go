/**
* @program: elasticsql
*
* @description:
*
* @author: lemo
*
* @create: 2023-08-14 18:50
**/

package main

import (
	"github.com/lemonyxk/elasticsql"
	"log"
)

func main() {

	var sql = "select * from a where id is null"

	log.Println(elasticsql.Convert(sql))

}
