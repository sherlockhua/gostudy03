
package main

import (
  // "encoding/json"
   "fmt"
   "log"
   "time"

   "github.com/influxdata/influxdb/client/v2"
)

const (
   MyDB          = "sys_info"
   username      = "admin"
   password      = ""
   MyMeasurement = "cpu_usage"
)

func main() {
   conn := connInflux()
   fmt.Println(conn)

   //insert
   WritesPoints(conn)

   //获取10条数据并展示
   qs := fmt.Sprintf("SELECT * FROM %s LIMIT %d", MyMeasurement, 10)
   res, err := QueryDB(conn, qs)
   if err != nil {
	   log.Fatal(err)
   }

   for i, row := range res[0].Series[0].Values {
	   for j,  value := range row {
		   /*
		t, err := time.Parse(time.RFC3339, row[j].(string))
		if err != nil {
			log.Fatal(err)
		}
		//fmt.Println(reflect.TypeOf(row[1]))
		valu := row[j].(json.Number)
		log.Printf("[%2d] %s: %s\n", i, t.Format(time.Stamp), valu)
		*/
		log.Printf("i:%d j:%d value:%#v\n", i, j, value)
		}
   }
}

func connInflux() client.Client {
   cli, err := client.NewHTTPClient(client.HTTPConfig{
	   Addr:     "http://127.0.0.1:8086",
	   Username: username,
	   Password: password,
   })
   if err != nil {
	   log.Fatal(err)
   }
   return cli
}

//query
func QueryDB(cli client.Client, cmd string) (res []client.Result, err error) {
   q := client.Query{
	   Command:  cmd,
	   Database: MyDB,
   }
   if response, err := cli.Query(q); err == nil {
	   if response.Error() != nil {
		   return res, response.Error()
	   }
	   res = response.Results
   } else {
	   return res, err
   }
   return res, nil
}

//Insert
func WritesPoints(cli client.Client) {
   bp, err := client.NewBatchPoints(client.BatchPointsConfig{
	   Database:  MyDB,
	   Precision: "s",
   })
   if err != nil {
	   log.Fatal(err)
   }

   tags := map[string]string{"cpu": "ih-cpu"}
   fields := map[string]interface{}{
	   "idle":   20.1,
	   "system": 43.3,
	   "user":   86.6,
   }

   pt, err := client.NewPoint(
	   "cpu_usage",
	   tags,
	   fields,
	   time.Now(),
   )
   if err != nil {
	   log.Fatal(err)
   }
   bp.AddPoint(pt)

   if err := cli.Write(bp); err != nil {
	   log.Fatal(err)
   }
}
