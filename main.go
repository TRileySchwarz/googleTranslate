
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"os"
)

func main() {
	fmt.Println("\n -- Starting program -- \n")

	// Comment/Uncomment the languages you want to translate to
	// the location of the language codes... https://cloud.google.com/translate/docs/languages
	languagesSupported := []string{
		"af",
		"sq",
	}

	sourceLanguage := "en"

	file, err := os.Open("untranslated/textToTranslate")
	if err != nil {
		fmt.Print("There was an issue opening the text file")
		fmt.Printf("This is the error: %v", err)
		return
	}
	defer file.Close()

	text, err := ioutil.ReadAll(file)
	fmt.Print(string(text))

	for i := 0; i < len(languagesSupported); i++ {
		err := translateText(languagesSupported[i], string(text), sourceLanguage)

		if err != nil {
		fmt.Print("There was an issue translating the text")
		}
	}

	fmt.Println("\n\n -- Closing Program --")
}

// This function takes in a language code and returns the corresponding full string of that code
func getLanguage(languageCode string) (string, error) {
	switch languageCode{
		case "af":
		return "Afrikaans", nil

		case "sq":
		return "Albanian", nil

		// Add new language conversions here
	}

	return "", errors.New("An unsupported language was submited")
}

func translateText(languageCode string, text string, sourceLanguage string) error {
	// Google translate API url
	url := "https://translate.googleapis.com/translate_a/single?client=gtx&sl=" + sourceLanguage + "&tl=" + languageCode + "&dt=t&q=" + url.QueryEscape(text)

	//log.Printf("Querying Google translate API with the following url: \n%v\n\n", url)

	// Create httpGet Response
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer func(){
		err = resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	// Read the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Cast the body as a string
	bodyString := string(body)

	split := strings.Split(bodyString, `","`)

	stringToWrite := split[0][4:]

	language, err := getLanguage(languageCode)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("\n\nThis is the result of the %v translation: %v", language, stringToWrite)

	file, err := os.Create("translated/"+ language + ".txt")
	defer func(){
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	_, err = file.WriteString(stringToWrite);
	if err != nil {
		log.Fatalln("error writing record to txt file:", err)
	}

	return nil
}



