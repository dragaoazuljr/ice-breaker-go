package scrapper

import (
	"os"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func CreateWebDriver() (selenium.WebDriver, error) {
	service, err := selenium.NewChromeDriverService(os.Getenv("CHOMEDRIVER"), 4444)
	if err != nil {
		return nil, err
	}
	defer service.Stop()

	caps := selenium.Capabilities{"browserName": "chrome"}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--headless",
		"user-data-dir=chrome-data",
	}})

	wd, err := selenium.NewRemote(caps, "")
	return wd, err
}
