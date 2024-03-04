package model

type ClubActivity struct {
	Athlete            ClubAthlete  `json:"athlete"`
	Name               string       `json:"name"`
	Distance           float64      `json:"distance"`
	MovingTime         int          `json:"moving_time"`
	ElapsedTime        int          `json:"elapsed_time"`
	TotalElevationGain float64      `json:"total_elevation_gain"`
	Type               ActivityType `json:"type"`
	SportType          ActivityType `json:"sport_type"`
	RiderName          string       `json:"-"`
}

type ClubAthlete struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type ActivityType string

var ActivityTypes = struct {
	Ride           ActivityType
	AlpineSki      ActivityType
	BackcountrySki ActivityType
	Hike           ActivityType
	IceSkate       ActivityType
	InlineSkate    ActivityType
	NordicSki      ActivityType
	RollerSki      ActivityType
	Run            ActivityType
	Walk           ActivityType
	Workout        ActivityType
	Snowboard      ActivityType
	Snowshoe       ActivityType
	Kitesurf       ActivityType
	Windsurf       ActivityType
	Swim           ActivityType
	VirtualRide    ActivityType
	EBikeRide      ActivityType

	WaterSport         ActivityType
	Canoeing           ActivityType
	Kayaking           ActivityType
	Rowing             ActivityType
	StandUpPaddling    ActivityType
	Surfing            ActivityType
	Crossfit           ActivityType
	Elliptical         ActivityType
	RockClimbing       ActivityType
	StairStepper       ActivityType
	WeightTraining     ActivityType
	Yoga               ActivityType
	WinterSport        ActivityType
	CrossCountrySkiing ActivityType
}{
	"Ride", "AlpineSki", "BackcountrySki", "Hike", "IceSkate", "InlineSkate", "NordicSki", "RollerSki",
	"Run", "Walk", "Workout", "Snowboard", "Snowshoe", "Kitesurf", "Windsurf", "Swim", "VirtualRide", "EBikeRide",

	"WaterSport", "Canoeing", "Kayaking", "Rowing", "StandUpPaddling", "Surfing",
	"Crossfit", "Elliptical", "RockClimbing", "StairStepper", "WeightTraining", "Yoga", "WinterSport", "CrossCountrySkiing",
}
