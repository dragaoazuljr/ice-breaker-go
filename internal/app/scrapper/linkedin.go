package scrapper

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func ScrapeLinkedInProfileData(profileUrl string) (string, error) {
	service, err := selenium.NewChromeDriverService(os.Getenv("CHOMEDRIVER"), 4444)
	if err != nil {
		return "", err
	}
	defer service.Stop()

	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--headless",
		"user-data-dir=chrome-data",
	}})

	wd, err := selenium.NewRemote(caps, "")
	if err != nil {
		return "", err
	}

	err = wd.Get("https://www.linkedin.com/login")
	if err != nil {
		return "", err
	}

	time.Sleep(1 * time.Second)

	currentUrl, err := wd.CurrentURL()
	if err != nil {
		return "", err
	}

	if strings.Contains(currentUrl, "login") {
		var email = os.Getenv("EMAIL")
		var password = os.Getenv("PASS")

		user, err :=wd.FindElement(selenium.ByID, "username")
		if err != nil {
			return "", err
		}

		if err := user.SendKeys(email); err != nil {
			return "", err
		}

		pass, err := wd.FindElement(selenium.ByID, "password")
		if err != nil {
			return "", err
		}

		if err := pass.SendKeys(password); err != nil {
			return "", err
		}

		button, err := wd.FindElement(selenium.ByCSSSelector, ".btn__primary--large.from__button--floating")
		if err != nil {
			return "", err
		}

		button.Click()

		time.Sleep(1 * time.Second)
	}


	if err := wd.Get(profileUrl); err != nil {
		return "", err
	}

	time.Sleep(3 * time.Second)

	page, _ := wd.PageSource()

	doc := soup.HTMLParse(page)

	nameElem := doc.FindStrict("h1", "class", "text-heading-xlarge inline t-24 v-align-middle break-words")
	titleElem := doc.FindStrict("div", "class", "text-body-medium break-words")
	locationElem := doc.FindStrict("span", "class", "text-body-small inline t-black--light break-words")
	aboutElem:= doc.Find("div", "class", "pv-shared-text-with-see-more")
	experienceElem:= doc.Find("div", "id", "experience")
	publicationsElem:= doc.Find("div", "id", "publications")

	var name string;
	var title string;
	var location string;
	var about string;
	var experience string;
	var publications string;

	fmt.Print(nameElem.Error, titleElem.Error, locationElem.Error, aboutElem.Error)
	re := regexp.MustCompile(`[\s\n\r]+`)

	if nameElem.Error == nil {
		name = re.ReplaceAllString(nameElem.FullText(), " ")
	}

	if titleElem.Error == nil {
		title = re.ReplaceAllString(titleElem.FullText(), " ")
	}

	if locationElem.Error == nil {
		location = re.ReplaceAllString(locationElem.FullText(), " ")
	}

	if aboutElem.Error == nil {
		about = re.ReplaceAllString(aboutElem.FullText(), " ")
	}

	if experienceElem.Error == nil {
		text := experienceElem.FindNextElementSibling().FindNextElementSibling().FullText() 
		experience = re.ReplaceAllString(text, " ")
	}

	if publicationsElem.Error == nil {
		text := publicationsElem.FindNextElementSibling().FindNextElementSibling().FullText() 
		publications = re.ReplaceAllString(text, " ")
	}

	return fmt.Sprintf("Name = %s, Title = %s, Location = %s, About = %s, Experience = %v, Publications: %v", name, title, location, about, experience, publications), nil
}
