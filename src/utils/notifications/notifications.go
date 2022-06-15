package notifications

import (
	"fmt"
	"net/http"
	"pagerduty/src/models"
	"strings"
)

const (
	ERROR_MESSAGE = "Too many 5xx"
)

type NotificationStatus struct {
	Success []string
	Failed  []string
}

func SendNotifications(developers []models.Developer) NotificationStatus {

	var status NotificationStatus
	for _, developer := range developers {
		err := SendNotification(developer.PhoneNumber)
		if err != nil {
			status.Failed = append(status.Failed, developer.Name)
		} else {
			status.Success = append(status.Success, developer.Name)
		}
	}

	return status
}

func SendNotification(number string) error {
	url := "https://run.mocky.io/v3/fd99c100-f88a-4d70-aaf7-393dbbd5d99f"
	method := "POST"

	payload := strings.NewReader(fmt.Sprintf(`{"phone_number": "%v", "message": "%v"}`, number, ERROR_MESSAGE))

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return err
	}
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("error sending notification")
	}
	return nil
}
