
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("\n -- Starting program -- \n")

	// Comment/Uncomment the languages you want to translate to
	// the location of the language codes... https://cloud.google.com/translate/docs/languages
	languagesSupported := []string{
		"Afrikaans",
		"Amharic",
		"Armenian",
		"Azerbaijani",
		"Basque",
		"Belarusian",
		"Bengali",
		"Bosnian",
		"Bulgarian",
		"Catalan",
		"Cebuano",
		"Chinese",
		"Corsican",
		"Croatian",
		"Czech",
		"Danish",
		"Dutch",
		"English",
		"Esperanto",
		"Estonian",
		"Finnish",
		"French",
		"Frisian",
		"Galician",
		"Georgian",
		"German",
		"Greek",
		"Gujarati",
		"Haitian Creole",
		"Hausa",
		"Hawaiian",
		"Hebrew",
		"Hindi",
		"Hmong",
		"Hungarian",
		"Icelandic",
		"Igbo",
		"Indonesian",
		"Irish",
		"Italian",
		"Japanese",
		"Javanese",
		"Kannada",
		"Kazakh",
		"Khmer",
		"Korean",
		"Kurdish",
		"Kyrgyz",
		"Lao",
		"Latin",
		"Latvian",
		"Lithuanian",
		"Luxembourgish",
		"Macedonian",
		"Malagasy",
		"Malay",
		"Malayalam",
		"Maltese",
		"Maori",
		"Marathi",
		"Mongolian",
		"Myanmar",
		"Nepali",
		"Norwegian",
		"Nyanja",
		"Pashto",
		"Persian",
		"Polish",
		"Portuguese",
		"Punjabi",
		"Romanian",
		"Russian",
		"Samoan",
		"Scots Gaelic",
		"Serbian",
		"Sesotho",
		"Shona",
		"Sindhi",
		"Sinhala",
		"Slovak",
		"Slovenian",
		"Somali",
		"Spanish",
		"Sundanese",
		"Swahili",
		"Swedish",
		"Tagalog",
		"Tajik",
		"Tamil",
		"Telugu",
		"Thai",
		"Turkish",
		"Ukrainian",
		"Urdu",
		"Uzbek",
		"Vietnamese",
		"Welsh",
		"Xhosa",
		"Yiddish",
		"Yoruba",
		"Zulu",
	}

	sourceLanguage := "en"

	wordPtr := flag.String("translate", "foo", "a string")
	flag.Parse()


	fmt.Print("The text we are translating is:  " + *wordPtr + "\n")

	for i := 0; i < len(languagesSupported); i++ {
		err := translateText(languagesSupported[i], *wordPtr, sourceLanguage)

		if err != nil {
		fmt.Print("There was an issue translating the text")
		}
	}

	fmt.Println("\n\n -- Closing Program --")
}

func translateText(language string, text string, sourceLanguage string) error {

	// Get the corresponding language code for the languageString requested
	languageCode := getLanguageCode(language)

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

	// Hack the weird string response we get
	split := strings.Split(bodyString, `","`)
	stringToWrite := split[0][4:]


	fmt.Printf("\n%v:   %v", language, stringToWrite)

	//file, err := os.Create("translated/thisIsNewTranslation.txt")
	//defer func(){
	//	err := file.Close()
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//}()

	//_, err = file.WriteString(stringToWrite);
	//if err != nil {
	//	log.Fatalln("error writing record to txt file:", err)
	//}

	return nil
}

// This function takes in a language code and return the corresponding full string of that code
func getLanguageCode(languageString string) string {
	switch languageString {
	case "Afrikaans": return	"af"
	case "Albanian": return	"sq"
	case "Amharic": return	"am"
	case "Arabic": return	"ar"
	case "Armenian": return	"hy"
	case "Azerbaijani": return	"az"
	case "Basque": return	"eu"
	case "Belarusian": return	"be"
	case  "Bengali": return	"bn"
	case "Bosnian": return	"bs"
	case  "Bulgarian": return	"bg"
	case  "Catalan": return	"ca"
	case  "Cebuano": return	"ceb"
	case  "Chinese": return   "zh"
	case   "Corsican": return	"co"
	case   "Croatian": return	"hr"
	case  "Czech": return	"cs"
	case   "Danish": return	"da"
	case   "Dutch": return	"nl"
	case  "English": return	"en"
	case "Esperanto": return	"eo"
	case "Estonian": return	"et"
	case  "Finnish": return	"fi"
	case  "French": return	"fr"
	case  "Frisian": return	"fy"
	case  "Galician": return	"gl"
	case  "Georgian": return	"ka"
	case   "German": return	"de"
	case  "Greek": return	"el"
	case  "Gujarati": return	"gu"
	case   "Haitian Creole": return	"ht"
	case  "Hausa": return	"ha"
	case  "Hawaiian": return	"haw"
	case  "Hebrew": return	"he"
	case   "Hindi": return	"hi"
	case   "Hmong": return	"hmn"
	case   "Hungarian": return	"hu"
	case  "Icelandic": return	"is"
	case   "Igbo": return	"ig"
	case  "Indonesian": return	"id"
	case   "Irish": return	"ga"
	case  "Italian": return	"it"
	case   "Japanese": return	"ja"
	case  "Javanese": return	"jw"
	case  "Kannada": return	"kn"
	case   "Kazakh": return	"kk"
	case  "Khmer": return	"km"
	case  "Korean": return	"ko"
	case  "Kurdish": return	"ku"
	case  "Kyrgyz": return	"ky"
	case   "Lao": return	"lo"
	case   "Latin": return	"la"
	case   "Latvian": return	"lv"
	case  "Lithuanian": return	"lt"
	case  "Luxembourgish": return	"lb"
	case   "Macedonian": return	"mk"
	case   "Malagasy": return	"mg"
	case  "Malay": return	"ms"
	case   "Malayalam": return	"ml"
	case   "Maltese": return	"mt"
	case  "Maori": return	"mi"
	case   "Marathi": return	"mr"
	case  "Mongolian": return	"mn"
	case  "Myanmar (Burmese)": return	"my"
	case  "Nepali": return	"ne"
	case  "Norwegian": return	"no"
	case   "Nyanja (Chichewa)": return	"ny"
	case  "Pashto": return	"ps"
	case   "Persian": return	"fa"
	case  "Polish": return	"pl"
	case  "Portuguese": return	"pt"
	case  "Punjabi": return	"pa"
	case  "Romanian": return	"ro"
	case   "Russian": return	"ru"
	case   "Samoan": return	"sm"
	case  "Scots Gaelic": return	"gd"
	case   "Serbian": return	"sr"
	case  "Sesotho": return	"st"
	case  "Shona": return "sn"
	case   "Sindhi": return	"sd"
	case   "Sinhala (Sinhalese)": return	"si"
	case  "Slovak": return	"sk"
	case   "Slovenian": return	"sl"
	case  "Somali": return	"so"
	case  "Spanish": return	"es"
	case  "Sundanese": return	"su"
	case   "Swahili": return	"sw"
	case  "Swedish": return	"sv"
	case   "Tagalog (Filipino)": return	"tl"
	case  "Tajik": return	"tg"
	case   "Tamil": return	"ta"
	case  "Telugu": return	"te"
	case  "Thai": return	"th"
	case  "Turkish": return	"tr"
	case  "Ukrainian": return	"uk"
	case   "Urdu": return	"ur"
	case   "Uzbek": return	"uz"
	case   "Vietnamese": return	"vi"
	case  "Welsh": return	"cy"
	case  "Xhosa": return "xh"
	case   "Yiddish": return	"yi"
	case  "Yoruba": return	"yo"
	case   "Zulu": return	"zu"
	}

	return ""
}



