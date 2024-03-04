package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/morris-kelly/strava600mfm/model"
	"github.com/spf13/viper"
)

const (
	baseURL      string = "https://www.strava.com/api/v3"
	itemsPerPage int    = 30
)

// Prerequisites
// 1. Create a Strava account
// 2. Create a Strava API application
// 3. Get an access token
// 4. Create a .env.local file in the root of the project
// 5. Add the following line to the .env.local file
// TOKEN=your_access_token
// 6. Run the program
// 7. The program will create a CSV file with the activities of the Proserv600MFM club
// 8. The CSV file will be named activitiesYYYY-MM-DD-HH-MM-SS.csv
func main() {
	viper.SetConfigName(".env.local")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
	csv := "RiderName,ActivityName,Distance (miles),MovingTime,ElapsedTime,TotalElevationGain,SportType\n"
	for i := 1; i < 10; i++ {

		requestString := "/clubs/Proserv600MFM/activities?page=" + fmt.Sprint(i)

		resp, err := GetStravaRequest(requestString)
		if err != nil {
			log.Fatal(err)
		}
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(string(body))

		activities := make([]model.ClubActivity, itemsPerPage)
		err = json.Unmarshal(body, &activities)
		if err != nil {
			log.Fatal(err)
		}
		for i := range activities {
			activities[i].RiderName = activities[i].Athlete.FirstName + " " + activities[i].Athlete.LastName
		}

		// Write the collection to CSV
		for _, activity := range activities {
			activity.Name = strings.ReplaceAll(activity.Name, ",", " ")
			activity.Name = strings.ReplaceAll(activity.Name, "  ", " ")

			activity.Distance = activity.Distance / 1000 * 0.6214

			activity.MovingTime = activity.MovingTime / 60

			activity.ElapsedTime = activity.ElapsedTime / 60

			activity.TotalElevationGain = activity.TotalElevationGain * 3.28084
			csv += fmt.Sprintf("%s,%s,%f,%d,%d,%f,%s\n", activity.RiderName, activity.Name, activity.Distance, activity.MovingTime, activity.ElapsedTime, activity.TotalElevationGain, activity.SportType)
		}

	}
	now := time.Now()
	nowString := now.Format("2006-01-02-15-04-05")
	err = os.WriteFile("activities"+nowString+".csv", []byte(csv), 0o644)
	if err != nil {
		log.Fatal(err)
	}
}

func GetStravaRequest(path string) (*http.Response, error) {
	req, err := http.NewRequest(http.MethodGet, baseURL+path, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+viper.GetString("TOKEN"))
	return http.DefaultClient.Do(req)
}
