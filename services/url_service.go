package services

import "fmt"

var (
	UrlService urlServiceInterface = &urlService{}
)

var malwareList = map[string]string{
	"1": "whatever.com",
	"2": "mail.google.com",
}

type urlService struct{}

type urlServiceInterface interface {
	MalwareExists(domain string) bool
}

func (us *urlService) MalwareExists(domain string) bool {
	for _, v := range malwareList {
		if domain == v {
			fmt.Println(v)
			return true
		}
	}
	return false
}
