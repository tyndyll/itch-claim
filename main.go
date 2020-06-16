

package main

import (
	"fmt"
	"log"
	"time"

	"github.com/gocolly/colly"
	"github.com/spf13/cobra"
)

var (
	password string
	username string
	url      string
	cmd      *cobra.Command
)

func init() {
	cmd = &cobra.Command{
		Use:   "itch-claim",
		Short: "Add all items from the #BLM itch.io bundle to your account",
		RunE:  LoginAndAddItems,
	}
	pf := cmd.PersistentFlags()
	pf.StringVarP(&username, "username", "", "", "itch.io username")
	pf.StringVarP(&password, "password", "", "", "itch.io password")
	pf.StringVarP(&url, "url", "", "", "your unique bundle url")
	cobra.MarkFlagRequired(pf,"username")
	cobra.MarkFlagRequired(pf,"password")
	cobra.MarkFlagRequired(pf,"url")
}

func LoginAndAddItems(cmd *cobra.Command, args []string) error {
	c := colly.NewCollector()
	var collyErr error

	if err := c.Limit(&colly.LimitRule{DomainGlob: "itch.io", Delay: 5 * time.Second}); err != nil {
		return err
	}

	log.Println("logging in")

	c.OnHTML(`.user_login_page`, func(e *colly.HTMLElement) {
		// We have errors in the login submission. Need to exit
 		if e.ChildText(".form_errors") != "" {
			collyErr = fmt.Errorf("could not login with provided credentials")
			return
		}

		// No form errors so we can assume we are trying to log in
		data := map[string]string{
			"username":   username,
			"password":   password,
			"csrf_token": e.ChildAttr(`input[name="csrf_token"]`, "value"),
		}
		if err := c.Post("https://itch.io/login", data); err != nil {
			collyErr = err
		}
	})

	if err := c.Visit("https://itch.io/login"); err != nil {
		return err
	}

	if collyErr != nil {
		return collyErr
	}

	log.Println("log in successful. commencing download")

	// We have logged in. Now parse the game rows from the download page
	c.OnHTML("div.game_row", func(e *colly.HTMLElement) {
		// If we have already claimed the game the row will contain a link instead of a form
		if e.ChildAttr(".game_download_btn", "href") == "" {
			log.Println("adding", e.ChildText(".game_title"))
			formData := map[string]string{
				"csrf_token": e.ChildAttr(`input[name="csrf_token"]`, "value"),
				"game_id":    e.ChildAttr(`input[name="game_id"]`, "value"),
				"action":     "claim",
			}
			if err := c.Post(url, formData); err != nil {
				log.Println(err.Error())
			}
		}
	})

	// Look for the next page button and visit it
	c.OnHTML("div.pager", func(e *colly.HTMLElement) {
		nextPage := e.ChildAttr("a.next_page", "href")
		if err := c.Visit(url + nextPage); err != nil {
			fmt.Println(err.Error())
		}
	})

	if err := c.Visit(url); err != nil {
		log.Print(err.Error())
	}
	return nil
}


func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalln(err.Error())
	}
}
