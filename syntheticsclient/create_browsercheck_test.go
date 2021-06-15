package main

import (
	"net/http"
	"testing"
	"time"
)

var (
	createBrowserBody = `{"id":191689,"name":"testtaste","type":"real_browser","frequency":10,"paused":false,"muted":false,"created_at":"2021-06-15T18:48:23.425Z","updated_at":"2021-06-15T18:48:23.425Z","links":{"self":"https://monitoring-api.rigor.com/v2/checks/191689","self_html":"https://monitoring.rigor.com/checks/real-browsers/191689","metrics":"https://monitoring-api.rigor.com/v2/checks/191689/metrics","last_run":null,"runs":"https://monitoring-api.rigor.com/v2/checks/real_browsers/191689/runs","share_link":"https://monitoring.rigor.com/share/350c659ff38a4781d0c9355252326a1beb8bac01ba4154a98267ac807daf7bbb*MTk2NTsyOzE5MTY4OQ=="},"status":{"last_code":200,"last_message":"","last_response_time":1939,"last_run_at":"2021-06-15T19:06:02.000Z","last_failure_at":null,"last_alert_at":null,"has_failure":false,"has_location_failure":null},"notifications":{"sms":false,"call":false,"email":true,"notify_after_failure_count":2,"notify_on_location_failure":true,"muted":false,"notify_who":[{"sms":false,"call":false,"email":true,"custom_user_email":null,"type":"user","links":{"self_html":"https://monitoring.rigor.com/admin/users/15649"},"id":15649}],"notification_windows":[],"escalations":[]},"response_time_monitor_milliseconds":null,"http_request_headers":{"User-Agent":"Mozilla/5.0 (X11; Linux x86_64; Rigor) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36"},"round_robin":true,"auto_retry":false,"enabled":true,"integrations":[],"url":"https://www.google.com/","user_agent":"Mozilla/5.0 (X11; Linux x86_64; Rigor) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.75 Safari/537.36","auto_update_user_agent":true,"viewport":{"width":1366,"height":768},"enforce_ssl_validation":true,"browser":{"label":"Google Chrome","code":"chrome"},"dns_overrides":{},"wait_for_full_metrics":true,"tags":[],"blackout_periods":[],"locations":[{"id":6,"world_region":"NA","region_code":"na-us-virginia","name":"N. Virginia, United States"}],"steps":[],"javascript_files":[],"threshold_monitors":[],"excluded_files":[],"cookies":[],"connection":{"download_bandwidth":20000,"upload_bandwidth":5000,"latency":28,"packet_loss":0}}`
)

func TestCreateBrowseCheck(t *testing.T) {
	setup()
	defer teardown()

	testMux.HandleFunc("/v2/checks/real_browsers", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "POST")
		w.Write([]byte(createBrowserBody))
	})

	resp, _, err := testClient.CreateBrowserCheck(&CreateBrowserCheck{
		ID:                   191689,
		Name:                 "testtaste",
		Type:                 "real_browser",
		Links:                struct {
			Self      string "json:\"self,omitempty\""
			SelfHTML  string "json:\"self_html,omitempty\""
			Metrics   string "json:\"metrics,omitempty\""
			LastRun   string "json:\"last_run,omitempty\""
			Runs      string "json:\"runs,omitempty\""
			ShareLink string "json:\"share_link,omitempty\""
		}{
			Self:      "https://monitoring-api.rigor.com/v2/checks/191689",
		},
		Status:               struct {
			LastCode           int       "json:\"last_code,omitempty\""
			LastMessage        string    "json:\"last_message,omitempty\""
			LastResponseTime   int       "json:\"last_response_time,omitempty\""
			LastRunAt          time.Time "json:\"last_run_at,omitempty\""
			LastFailureAt      time.Time "json:\"last_failure_at,omitempty\""
			LastAlertAt        time.Time "json:\"last_alert_at,omitempty\""
			HasFailure         bool      "json:\"has_failure,omitempty\""
			HasLocationFailure bool      "json:\"has_location_failure,omitempty\""
		}{
			LastCode:           200,
		},
		URL:                  "https://www.google.com/",
	})
	if err != nil {
		t.Fatal(err)
	}
	if resp.Name != "testtaste" {
		t.Errorf("\nreturned: %#v\n\n want: %#v\n", resp.Name, "testtaste")
	}
	if resp.Type != "real_browser" {
		t.Errorf("\nreturned: %#v\n\n want: %#v\n", resp.Type, "real_browser")
	}
	if resp.URL != "https://www.google.com/" {
		t.Errorf("\nreturned: %#v\n\n want: %#v\n", resp.URL, "https://www.google.com/")
	}
	if resp.ID != 191689 {
		t.Errorf("\nreturned: %#v\n\n want: %#v\n", resp.ID, 191689)
	}
	if resp.Links.Self != "https://monitoring-api.rigor.com/v2/checks/191689" {
		t.Errorf("\nreturned: %#v\n\n want: %#v\n", resp.Links.Self, "https://monitoring-api.rigor.com/v2/checks/191689")
	}
	if resp.Status.LastCode != 200 {
		t.Errorf("\nreturned: %#v\n\n want: %#v\n", resp.Status.LastCode, 200)
	}
}
