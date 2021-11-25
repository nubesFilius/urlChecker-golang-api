package services

import "fmt"

var (
	UrlService urlServiceInterface = &urlService{}
)

// Mocking a local db
var malwareList = map[string]string{
	"1": "whatever.com",
	"2": "mail.google.com",
}

type urlService struct{}

type urlServiceInterface interface {
	MalwareExists(domain string) bool
}

func (us *urlService) MalwareExists(domain string) bool {
	// TODO:Need to parse url to remove ports, need to add validation of url as well
	for _, v := range malwareList {
		if domain == v {
			fmt.Println(v)
			return true
		}
	}
	return false
}
