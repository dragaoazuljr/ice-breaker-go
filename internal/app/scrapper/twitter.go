package scrapper

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func ScapeTwitterProfileData(profileUrl string) (string, error) {
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

	err = wd.Get("https://www.twitter.com/login")
	if err != nil {
		return "", err
	}

	time.Sleep(1 * time.Second)

	currentUrl, _ := wd.CurrentURL()
	if strings.Contains(currentUrl, "login") {
		var email = os.Getenv("EMAIL")
		var password = os.Getenv("PASS")
		var username = os.Getenv("USERNAME")

		wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
			_, err := wd.FindElement(selenium.ByCSSSelector, "input[autocomplete=username]")
			if err != nil {
				return false, nil
			}
			return true, nil
		}, 5*time.Second)

		user, err := wd.FindElement(selenium.ByCSSSelector, "input[autocomplete=username]")

		if err != nil {
			return "", err
		}

		if err := user.SendKeys(email); err != nil {
			return "", err
		}

		next, err := wd.FindElements(selenium.ByCSSSelector, "div[role=button]")
		if err != nil {
			return "", err
		}

		next[3].Click()


		validateUsername, err := wd.FindElement(selenium.ByCSSSelector, "input[name=text]")

		fmt.Println("Finded 2")

		if err != nil {
			fmt.Print("Skipped validation")
		} else {
			if err := validateUsername.SendKeys(username); err != nil {
				return "", err
			}

			next, err := wd.FindElements(selenium.ByCSSSelector, "div[role=button]")
			if err != nil {
				return "", err
			}

			next[1].Click()
			time.Sleep(2 * time.Second)
		}

		wd.WaitWithTimeout(func(wd selenium.WebDriver) (bool, error) {
			_, err := wd.FindElement(selenium.ByCSSSelector, "input[name=password]")
			if err != nil {
				return false, nil
			}

			fmt.Println("Finded")

			return true, nil
		}, 5*time.Second)

		pass, err := wd.FindElement(selenium.ByCSSSelector, "input[name=password]")
		if err != nil {
			return "", err
		}

		if err := pass.SendKeys(password); err != nil {
			return "", err
		}

		buttons, err := wd.FindElements(selenium.ByCSSSelector, "div[role=button]")
		if err != nil {
			return "", err
		}

		buttons[2].Click()

		time.Sleep(1 * time.Second)
	}

	if err := wd.Get(profileUrl); err != nil {
		fmt.Println(err)
		return "", err
	}

	time.Sleep(3 * time.Second)

	tweetsDiv, err := wd.FindElements(selenium.ByCSSSelector, "div[data-testid=tweetText]")
	if err != nil {
		return "", err
	}

	var tweets []string

	for _, tweet := range tweetsDiv {
		tweetText, err := tweet.Text()
		if err != nil {
			return "", err
		}

		tweets = append(tweets, tweetText)
	}

	wd.ExecuteScript("window.scrollTo(0, 1000)", nil)

	time.Sleep(1 * time.Second)

	tweetsDiv, err = wd.FindElements(selenium.ByCSSSelector, "div[data-testid=tweetText]")
	if err != nil {
		return "", err
	}

	for _, tweet := range tweetsDiv {
		tweetText, err := tweet.Text()
		if err != nil {
			return "", err
		}

		if !contains(tweets, tweetText) {
		  re := regexp.MustCompile(`[\s\n\r]+`)	
			tweets = append(tweets,re.ReplaceAllString(tweetText," "))
		}
	}
		
	return concat(tweets, "| "), nil
}

func contains(s []string, e string) bool {
    for _, a := range s {
        if a == e {
            return true
        }
    }
    return false
}

func concat(s []string, e string) string {
	var result string
	for _, a := range s {
		result = result + a + e
	}
	return result
}
