package helper

import (
	"crypto/rand"
	"fmt"
	"os"
	"regexp"
	"time"
)

// default location
var loc *time.Location

func init() { //to make sure that it will be initialized
	var err error
	loc, err = time.LoadLocation(os.Getenv("TIME_ZONE"))
	if err != nil {
		loc = time.UTC // fallback to UTC if loading fails
	}
}

// UserInfo contains both the unique ID and current timestamp
type UtilInfo struct {
	ID        string
	CreatedAt string
}

type Response struct {
	Status   bool
	Response interface{} `json:"result"`
}

// GetCurrentTime returns the current time in Africa/Kigali in YYYY-MM-DD HH:MM:SS
func GetCurrentTime() string {
	return time.Now().In(loc).Format("2006-01-02 15:04:05")
}

func GetResponse(statusPar bool, responsePar interface{}) Response {

	return Response{
		Status:   statusPar,
		Response: responsePar,
	}
}

// GenerateUserID generates a unique ID with username + timestamp + random suffix
func GenerateID(userName string) string {

	timestamp := time.Now().In(loc).Format("20060102150405") // YYYYMMDDHHMMSS

	// Random 4-digit hex suffix
	b := make([]byte, 2)
	rand.Read(b)
	randomSuffix := fmt.Sprintf("%X", b)

	return fmt.Sprintf("%s_%s_%s", userName, timestamp, randomSuffix)
}

func CleanString(input string) string {
	// Keep only letters and numbers
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	cleaned := re.ReplaceAllString(input, "")
	return cleaned
}
func GenerateInfos(userName string) UtilInfo {
	timestamp := time.Now().In(loc).Format("2006-01-02 15:04:05") //this is how go managed format

	// Random 4-digit hex suffix
	b := make([]byte, 2)
	rand.Read(b)
	randomSuffix := fmt.Sprintf("%X", b)
	name := CleanString(userName)

	id := fmt.Sprintf("%s_%s_%s", name, time.Now().In(loc).Format("20060102150405"), randomSuffix)

	return UtilInfo{
		ID:        id,
		CreatedAt: timestamp,
	}
}
