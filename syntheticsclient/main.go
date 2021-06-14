package main

// import (
// 	"fmt"
// 	"os"

// )

// func main() {

// 	token := os.Getenv("API_ACCESS_TOKEN")

// 	c := NewClient(token)

	
// 	//res, _, err := c.GetCheck(175389) //175389

// 	// o := GetChecksOptions{
// 	// 	Page: 0,
// 	// 	PerPage: 0,
// 	// }
// 	// res, _, err := c.GetChecks(&o)

// 	//HTTP CHECK Endpoints
// 	// res, _, err := c.GetHttpCheck(175389)
// 	// res, _, err := c.CreateHttpCheck(`{"name":"Clippie Homepage","tags":["test"],"frequency":5,"round_robin":true,"auto_retry":false,"enabled":true,"response_time_monitor_milliseconds":500,"http_request_headers":{"User-Agent":"Mozilla/5.0 (compatible; Rigor/1.0.0; http://rigor.com)"},"notifications":{"sms":false,"email":true,"call":false,"notify_who":[{"sms":false,"custom_email":"jeremyrhicks@splunk.com","call":false}],"notify_after_failure_count":2,"notify_on_location_failure":true,"muted":false},"url":"https://www.google.com","http_method":"GET","success_criteria":[{"action_type":"presence_of_text","comparison_string":"About"}],"connection":{"upload_bandwidth":5000,"download_bandwidth":20000,"latency":28,"packet_loss":0}}`)
// 	// res, _, err := c.UpdateHttpCheck(191484,`{"name":"ttt","tags":["test"],"frequency":5,"round_robin":true,"auto_retry":false,"enabled":true,"response_time_monitor_milliseconds":500,"http_request_headers":{"User-Agent":"Mozilla/5.0 (compatible; Rigor/1.0.0; http://rigor.com)"},"notifications":{"sms":false,"email":true,"call":false,"notify_who":[{"sms":false,"custom_email":"jeremyrhicks@splunk.com","call":false}],"notify_after_failure_count":2,"notify_on_location_failure":true,"muted":false},"url":"https://www.google.com","http_method":"GET","success_criteria":[{"action_type":"presence_of_text","comparison_string":"About"}],"connection":{"upload_bandwidth":5000,"download_bandwidth":20000,"latency":28,"packet_loss":0}}`)
// 	// res, err := c.DeleteHttpCheck(191496)

// 	//Browser CHECK Endpoints
// 	// res, _, err := c.CreateBrowserCheck(`{"name":"test","type":"real_browser","frequency":10,"round_robin":true,"auto_retry":false,"enabled":true,"integrations":[],"http_request_headers":{"User-Agent":"Mozilla/5.0 (X11; Linux x86_64; Rigor) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36"},"notifications":{"sms":false,"call":false,"email":true,"notify_after_failure_count":2,"notify_on_location_failure":true,"muted":false,"notify_who":[{"sms":false,"call":false,"email":true,"type":"user","links":{"self_html":"https://monitoring.rigor.com/admin/users/15649"},"id":15649}],"notification_windows":[],"escalations":[]},"url":"https://www.google.com/","user_agent":"Mozilla/5.0 (X11; Linux x86_64; Rigor) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36","auto_update_user_agent":true,"viewport":{"width":1366,"height":768},"enforce_ssl_validation":true,"browser":{"type":"chrome"},"dns_overrides":{},"wait_for_full_metrics":true,"tags":[],"blackout_periods":[{"timezone":"Eastern Time (US & Canada)","start_time":"2000-01-01T07:00:00.000Z","end_time":"2000-01-01T14:00:00.000Z","repeat_type":"daily","duration_in_minutes":420,"is_repeat":true,"created_at":"2020-11-11T21:54:32.000Z","updated_at":"2021-03-30T19:02:54.000Z"}],"steps":[],"javascript_files":[],"threshold_monitors":[],"excluded_files":[],"cookies":[],"connection":{"download_bandwidth":20000,"upload_bandwidth":5000,"latency":28,"packet_loss":0}}`)
// 	// res, _, err := c.UpdateBrowserCheck(191619, `{"name":"testtaste","type":"real_browser","frequency":10,"round_robin":true,"auto_retry":false,"enabled":true,"integrations":[],"http_request_headers":{"User-Agent":"Mozilla/5.0 (X11; Linux x86_64; Rigor) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36"},"notifications":{"sms":false,"call":false,"email":true,"notify_after_failure_count":2,"notify_on_location_failure":true,"muted":false,"notify_who":[{"sms":false,"call":false,"email":true,"type":"user","links":{"self_html":"https://monitoring.rigor.com/admin/users/15649"},"id":15649}],"notification_windows":[],"escalations":[]},"url":"https://www.google.com/","user_agent":"Mozilla/5.0 (X11; Linux x86_64; Rigor) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36","auto_update_user_agent":true,"viewport":{"width":1366,"height":768},"enforce_ssl_validation":true,"browser":{"type":"chrome"},"dns_overrides":{},"wait_for_full_metrics":true,"tags":[],"blackout_periods":[{"timezone":"Eastern Time (US & Canada)","start_time":"2000-01-01T07:00:00.000Z","end_time":"2000-01-01T14:00:00.000Z","repeat_type":"daily","duration_in_minutes":420,"is_repeat":true,"created_at":"2020-11-11T21:54:32.000Z","updated_at":"2021-03-30T19:02:54.000Z"}],"steps":[],"javascript_files":[],"threshold_monitors":[],"excluded_files":[],"cookies":[],"connection":{"download_bandwidth":20000,"upload_bandwidth":5000,"latency":28,"packet_loss":0}}`)
// 	// res, err := c.DeleteBrowserCheck(191619)

// 	if err != nil {
// 		fmt.Println(err)
// 	} else {
// 		JsonPrint(res)
// 	}
// }