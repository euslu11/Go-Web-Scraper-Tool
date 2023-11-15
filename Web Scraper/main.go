package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/common-nighthawk/go-figure"
)

func TheHackerNews() ([]string, []string, []string, error) {
	url := "https://thehackernews.com/"

	res, err := http.Get(url)
	if err != nil {
		return nil, nil, nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, nil, nil, err
	}

	var titles, descriptions, dates []string

	doc.Find(".body-post").Each(func(i int, s *goquery.Selection) {
		title := s.Find(".home-title").Text()
		titles = append(titles, title)

		description := s.Find(".home-desc").Text()
		date := s.Find(".h-datetime").Text()

		descriptions = append(descriptions, description)
		dates = append(dates, date)
	})

	return titles, descriptions, dates, nil
}

func WebAslan() ([]string, []string, []string, error) {
	url := "https://www.sondakika.com/webaslan/"

	res, err := http.Get(url)
	if err != nil {
		return nil, nil, nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, nil, nil, err
	}

	var titles, descriptions, dates []string

	doc.Find(".nws").Each(func(i int, s *goquery.Selection) {
		titleSelection := s.Find(".title")
		title := titleSelection.Text()

		descriptionSelection := s.Find(".news-detail.news-column")
		description := descriptionSelection.Text()

		date := s.Find(".date").Text()
		

		titles = append(titles, title)
		descriptions = append(descriptions, description)
		dates = append(dates, date)
	})

	return titles, descriptions, dates, nil
}

func extractDateFromWebAslan(dateLink string) string {
	parts := strings.Split(dateLink, "/")
	if len(parts) >= 4 {
		return parts[3]
	}

	return "Tarih Bulunamadı"
}

func SecurityIntelligance() ([]string, []string, []string, error) {
	url := "https://securityintelligence.com/"

	res, err := http.Get(url)
	if err != nil {
		return nil, nil, nil, err
	}
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, nil, nil, err
	}

	var titles, descriptions, dates []string

	doc.Find(".article__text_container").Each(func(i int, s *goquery.Selection) {
		titleSelection := s.Find(".article__title")
		title := titleSelection.Text()

		descriptionSelection := s.Find(".article__excerpt")
		description := descriptionSelection.Text()

		dateSelection := s.Find(".article__date")
		date := extractDatefromIntelligance(dateSelection)

		titles = append(titles, title)
		descriptions = append(descriptions, description)
		dates = append(dates, date)
	})

	return titles, descriptions, dates, nil
}

func extractDatefromIntelligance(selection *goquery.Selection) string {
	dateText := selection.Text()

	if dateText == "" {
		dateText = selection.Find(".article_date").Parent().Text()
	}

	if dateText != "" {
		return dateText
	}

	return "Tarih Bulunamadı"
}

func main() {
	myFigure := figure.NewFigure("Yavuzlar", "", true)
	myFigure.Print()
	myFigure2 := figure.NewFigure("Web Scaper Tool", "small", true)
	myFigure2.Print()
	websiteFlag := flag.Int("website", 1, "Haber sitesi seçimi (1: TheHackerNews, 2: WebAslan, 3: SecurityIntelligance)")
	dateFlag := flag.Bool("date", true, "Tarih bilgisini gösterme")
	helpFlag := flag.Bool("h", false, "Tool yardımını göster")
	flag.Parse()

	if *helpFlag {
		printHelp()
		return
	}

	var titles, descriptions, dates []string
	var err error

	switch *websiteFlag {
	case 1:
		titles, descriptions, dates, err = TheHackerNews()
	case 2:
		titles, descriptions, dates, err = WebAslan()
	case 3:
		titles, descriptions, dates, err = SecurityIntelligance()
	default:
		log.Fatal("Geçersiz haber sitesi seçimi. Lütfen 1, 2 veya 3 girin.")
	}

	if err != nil {
		log.Fatal("Web sitesi çekme hatası:", err)
	}

	if !*dateFlag {
		dates = nil
	}

	for i, title := range titles {
		fmt.Printf("Başlık %d: %s\n", i+1, title)
		if descriptions != nil {
			fmt.Printf("Açıklama %d: %s\n", i+1, descriptions[i])
		}
		if *dateFlag {
			fmt.Printf("Tarih %d: %s\n", i+1, dates[i])
		}
	}
}
func printHelp() {
	fmt.Println("Web Scraper Tool")
	fmt.Println("Usage:")
	fmt.Println("  go run main.go [flags]")
	fmt.Println()
	fmt.Println("Flags:")
	flag.PrintDefaults()
}
