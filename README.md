# googleTranslate
Repository used to take an input of a text string via the command line and will print corresponding translations to console

requires a working installation of go https://golang.org/doc/install


1 - Inside of main.go on lines 20ish, apply the corresponding language codes you wish to convert to, comment out the ones you dont want to convert to

2 - run the following in the root folder, where the translate flag corresponds to the text you wish to translate  ``$ go run main.go -translate="This is the text we wish to translate""``

3 - Providing nothing is broken, the translations will be printed to console


*** There are some weird edge cases by having to use this as a free version, such that the contents of the passed string will do weird stuff
if there is the following "," anywhere in the file. 

There is also the risk of being rate limited due to the free status of this api being used. It's unclear what the limit on the rates are, so dont use more than you need to
