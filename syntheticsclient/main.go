package main

import (
	"fmt"
	"os"

)

func main() {

	token := os.Getenv("API_ACCESS_TOKEN")

	c := NewClient(token)

	//HTTP CHECK Endpoints
	//res, _, err := c.GetCheck(175389) //175389

	// o := GetChecksOptions{
	// 	Page: 0,
	// 	PerPage: 0,
	// }
	// res, _, err := c.GetChecks(&o)

	// res, _, err := c.GetHttpCheck(175389)
	// res, _, err := c.CreateHttpCheck(`{"name":"TFTest2","tags":["test"],"frequency":5,"round_robin":true,"auto_retry":false,"enabled":true,"response_time_monitor_milliseconds":500,"http_request_headers":{"User-Agent":"Mozilla/5.0 (compatible; Rigor/1.0.0; http://rigor.com)"},"notifications":{"sms":false,"email":true,"call":false,"notify_who":[{"sms":false,"custom_email":"jeremyrhicks@splunk.com","call":false}],"notify_after_failure_count":2,"notify_on_location_failure":true,"muted":false},"url":"https://www.google.com","http_method":"GET","success_criteria":[{"action_type":"presence_of_text","comparison_string":"About"}],"connection":{"upload_bandwidth":5000,"download_bandwidth":20000,"latency":28,"packet_loss":0}}`)
	// res, _, err := c.UpdateHttpCheck(191484,`{"name":"ttt","tags":["test"],"frequency":5,"round_robin":true,"auto_retry":false,"enabled":true,"response_time_monitor_milliseconds":500,"http_request_headers":{"User-Agent":"Mozilla/5.0 (compatible; Rigor/1.0.0; http://rigor.com)"},"notifications":{"sms":false,"email":true,"call":false,"notify_who":[{"sms":false,"custom_email":"jeremyrhicks@splunk.com","call":false}],"notify_after_failure_count":2,"notify_on_location_failure":true,"muted":false},"url":"https://www.google.com","http_method":"GET","success_criteria":[{"action_type":"presence_of_text","comparison_string":"About"}],"connection":{"upload_bandwidth":5000,"download_bandwidth":20000,"latency":28,"packet_loss":0}}`)
	res, err := c.DeleteHttpCheck(191494)

	if err != nil {
		fmt.Println(err)
	} else {
		JsonPrint(res)
	}
}