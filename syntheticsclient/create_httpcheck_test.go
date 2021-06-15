package main

import (
	"net/http"
	"testing"
)

var (
	createHttpBody = `{"id":191588,"name":"Clippie Homepage","type":"http","frequency":5,"paused":false,"muted":false,"created_at":"2021-06-14T14:12:31.494Z","updated_at":"2021-06-14T14:12:31.494Z","links":{"self":"https://monitoring-api.rigor.com/v2/checks/191588","self_html":"https://monitoring.rigor.com/checks/http/191588","metrics":"https://monitoring-api.rigor.com/v2/checks/191588/metrics","last_run":""},"tags":[{"id":9,"name":"test"}],"status":{"last_code":0,"last_message":"","last_response_time":0,"last_run_at":"0001-01-01T00:00:00Z","last_failure_at":"0001-01-01T00:00:00Z","last_alert_at":"0001-01-01T00:00:00Z","has_failure":false,"has_location_failure":false},"round_robin":true,"auto_retry":false,"enabled":true,"blackout_periods":[],"locations":[{"id":1,"name":"N. Virginia, United States","world_region":"NA","region_code":"na-us-virginia"}],"integrations":[],"http_request_headers":{"User-Agent":"Mozilla/5.0 (compatible; Rigor/1.0.0; http://rigor.com)"},"notifications":{"sms":false,"email":true,"call":false,"notify_who":[{"sms":false,"email":true,"call":false,"links":{}}],"notify_after_failure_count":2,"notify_on_location_failure":true,"notification_windows":[],"escalations":[],"muted":false},"url":"https://www.google.com","http_method":"GET","success_criteria":[{"action_type":"presence_of_text","comparison_string":"About","created_at":"2021-06-14T14:12:31.526Z","updated_at":"2021-06-14T14:12:31.526Z"}],"connection":{"upload_bandwidth":5000,"download_bandwidth":20000,"latency":28,"packet_loss":0}}`
)

func TestCreateHttpCheck(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/v2/checks/http", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.Write([]byte(createHttpBody))
	})

	resp, _, err := testClient.CreateHttpCheck(&CreateHttpCheck{
		Name:               "Clippie Homepage",
		Type:               "http",
		Frequency:          5,
		URL:                "https://www.google.com",
	})

	if err != nil {
		t.Fatal(err)
	}
	if resp.Name != "Clippie Homepage"{
		t.Errorf("\nreturned: %#v\n\n want: %#v\n", resp.Name, "Clippie Homepage")
	}
	if resp.Type != "http" {
		t.Errorf("\nreturned: %#v\n\n want: %#v\n", resp.Type, "http")
	}
	if resp.Frequency != 5 {
		t.Errorf("\nreturned: %#v\n\n want: %#v\n", resp.Frequency, 5)
	}
	if resp.URL != "https://www.google.com" {
		t.Errorf("returned \n\n%#v want \n\n%#v", resp.URL, "https://www.google.com")
	}
	
}

