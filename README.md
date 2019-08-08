# googleTranslate
Repository used to take an input of a txt file and output the translation of it to a new txt file

requires a working installation of go https://golang.org/doc/install

1 - The "untranslated" folder contains a txt document, open that up and paste the text you want to convert there

2 - Inside of main.go on lines 20ish, apply the corresponding language codes you wish to convert to

3 - For any added languages, make sure to add to the function getLanguageCode the switch statement that returns the long-form of the language

4 - run the following in the root folder  ``$ go run main.go``

5 - Providing nothing is broken, you will find a version of the translated text inside of the "translated" folder
