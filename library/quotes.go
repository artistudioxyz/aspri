package library

import (
	"encoding/json"
	"flag"
	"fmt"
)

var (
	/** Extras */
	QuoteofTheDayFlag = flag.Bool("quote-of-the-day", false, "show quote of the day")
)

/** Initiate Quote Function */
func InitiateQuoteFunction() {
	if *QuoteofTheDayFlag {
		QuoteofTheDay()
		return
	}
}

/** Show Quote of The Day */
func QuoteofTheDay() {
	/** Quote Data Type */
	type quote struct {
		Text   string `json:"text"`
		Author string `json:"author"`
	}

	/** Get Quootes */
	responseText := getDataFromAPI("https://type.fit/api/quotes")

	/** Decode */
	quotes := []quote{}
	err := json.Unmarshal(responseText, &quotes)
	if err != nil {
		fmt.Println(err)
		return
	}

	/** Print The Quote */
	fmt.Println("----------------------------------------")
	fmt.Println(quotes[0].Text)
	fmt.Println(quotes[0].Author)
	fmt.Println("----------------------------------------")
}
