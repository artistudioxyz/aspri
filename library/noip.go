package library

import (
	"fmt"
	"net/http"
)

/** Initiate NoIP Function */
func InitiateNoIPFunction(flags Flag) {
	/** removeFilesExceptExtensions */
	if *flags.NoIP && *flags.Update && *flags.Username != "" && *flags.Password != "" && *flags.Hostname != "" {
		UpdateNoIPHostName(*flags.Username, *flags.Password, *flags.Hostname)
	}
}

/** Update Hostname IP */
func UpdateNoIPHostName(username string, password string, hostname string) {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://dynupdate.no-ip.com/nic/update", nil)
	if err != nil {
		fmt.Println("❌ ", err)
		return
	}

	q := req.URL.Query()
	q.Add("hostname", hostname)
	req.URL.RawQuery = q.Encode()

	req.Header.Add("User-Agent", fmt.Sprintf("coppit docker no-ip/.1 %s", username))
	req.SetBasicAuth(username, password)

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("❌ ", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("✅ Success update NoIP hostname", hostname, resp.Status)
}
