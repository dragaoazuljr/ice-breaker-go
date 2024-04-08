package scrapper

import (
	"fmt"
	"os"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func GetFirstGoogleResultLink(query string, complement string) (string, error) {
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

	if err := wd.Get("https://www.google.com"); err != nil {
		return "", err
	}

	elem, err := wd.FindElement(selenium.ByCSSSelector, "textarea[name=q]")
	if err != nil {
		return "", err
	}

	if err := elem.SendKeys(fmt.Sprintf("%s %s", query, complement)); err != nil {
		return "", err
	}

	if err := elem.Submit(); err != nil {
		return "", err
	}

	wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
		_, err := wd.FindElement(selenium.ByCSSSelector, "#search")
		if err != nil {
			return false, nil
		}
		return true, nil
	}, 10000)

	elem, err = wd.FindElement(selenium.ByCSSSelector, "#search a")

	if err != nil {
		panic(err)
	}

	link, err := elem.GetAttribute("href")

	if err != nil {
		return "", err
	}

	return link, nil
}
