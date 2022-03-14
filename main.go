package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var usageInformation = ` Usage: go run ./main.go [command...] [additionalParams...]

Commads:
	-list command 
		go run main.go list

	-search command
		go run main.go search <bookName>
		go run main.go search Return
		go run main.go search Lord 
		go run main.go search RiNg

	-get command
		go run main.go get <bookID>
		go run main.go get 5

	-delete command
	go run main.go delete <bookID>
	go run main.go delete 5

	-buy command
	go run main.go buy <bookID> <quantity>
	go run main.go buy 5 2
`

type author struct {
	AuthorID string
	name     string
}

type book struct {
	id            string
	bookName      string
	numberOfPages int
	stockCount    int
	price         int
	ISBN          string
	stockCode     string
	author
}

//convert book params to string to print
func (v book) ToString() string {

	return "name: " + v.bookName + " id: " + v.id + " stockCount: " + fmt.Sprint(v.stockCount) + " stockCode: " + fmt.Sprint(v.stockCode) + " price: " + fmt.Sprint(v.price) + "₺"
}

// check is bookname contains searchText
func (b book) isNameContains(searchText string) bool {
	//contains  is caseSensitive
	return strings.Contains(strings.ToLower(b.bookName), strings.ToLower(searchText))
}

func main() {
	//available books , mocked with test data
	bookLiblary := getBooks()
	//can be list or search
	firstQuery := os.Args[1]
	//understand command and redirect to handler. Lower case to make it not case sensitive
	switch strings.ToLower(os.Args[1]) {
	//if the request is list
	case "list":
		//print all books
		listHandler(bookLiblary)
	case "search":
		//is book name provided?

		searchHandler(bookLiblary)
	case "get":
		getHandler(bookLiblary)
	case "delete":

		deleteHandler(bookLiblary)

	case "buy":
		buyHandler(bookLiblary)

	default:
		printUsageAndExit(firstQuery + " is not supported.")

	}

}

//buy book from liblary if there is enought book in stock
func buyHandler(bookLiblary []book) {
	// arg lenght should be 4
	if len(os.Args) != 4 {
		printUsageAndExit("invalid query")
	}
	// os.Args[0] is not needed, os.Args[1] is command aka "buy"
	id := os.Args[2]
	quantity, err := strconv.Atoi(os.Args[3])

	//if its not convert string value to int
	//example
	//buy 2 bookName is false
	//buy 2 3 is correct
	if err != nil {
		printUsageAndExit("please type int value as for <quantity> like 'buy <id> <quantity>")
	}
	//for each book in liblary check is avaiable then check is quantity is okey
	for _, book := range bookLiblary {

		if book.id == id {

			if book.stockCount >= quantity {
				book.stockCount = book.stockCount - quantity
				fmt.Println("Great you succesfuly ordered !")
				book.ToString()
				fmt.Println(book.ToString())
				os.Exit(0)
			} else {
				book.ToString()
				printUsageAndExit("Sory , we dont have enought stock")

			}
		}
	}
	printUsageAndExit("book id is not correct , verify book id is exist")
}

// delete book by id if exist
func deleteHandler(bookLiblary []book) {

	// check is command correct
	// must be equal to 3  	delete		 <id>
	//                		os.Args[2]   [3]
	if len(os.Args) != 3 {
		printUsageAndExit("is unsupported command")
	}
	//get id of the book
	queryID := os.Args[2]

	for _, book := range bookLiblary {

		if book.id == queryID {
			//this is dummy delete operation
			fmt.Println("Book succesuly removed, book information : " + book.ToString())
			os.Exit(0)
		}
	}
	printUsageAndExit(" we dont have a book with ID: " + queryID)
}

//search by word
func searchHandler(bookLiblary []book) {
	// [1] = app tep location , [2] = query , [2:] book name
	if len(os.Args) == 3 {
		//join array with space

		searchText := strings.Join(os.Args[2:], " ")

		//to display not available meesage at the end of the search
		notAvailableFlag := true

		//iterate over array
		for _, book := range bookLiblary {

			if book.isNameContains(searchText) {
				fmt.Println(book.ToString())
				notAvailableFlag = false
			}
		}

		if notAvailableFlag {
			fmt.Println("The book '" + searchText + "' is not available. You can get all book name with 'list' command")
		}
		os.Exit(0)
	} else {
		//looks like its unsupported request
		//go run .\main.go search sasdas asdsa
		printUsageAndExit("")
	}
}

//print all books
func listHandler(bookLiblary []book) {
	//just print all of the books
	for _, b := range bookLiblary {
		fmt.Println(b.ToString())
	}

	os.Exit(0)
}

//get book by id
func getHandler(bookLiblary []book) {

	// must be equal to 3  	delete		 <id>
	//                		os.Args[2]   [3]

	if len(os.Args) != 3 {
		printUsageAndExit("unsupported parameters")
	}
	queryID := os.Args[2]

	for _, book := range bookLiblary {

		if book.id == queryID {
			fmt.Println(book.ToString())
			os.Exit(0)
		}
	}
	printUsageAndExit(" we dont have a book with ID: " + queryID)

}

//If there is an error occured like invalid param .. just print usage information with optional error msg
func printUsageAndExit(optionalText string) {
	fmt.Println(optionalText)
	fmt.Println(usageInformation)
	os.Exit(1)

}

// Get books from source , currenlty its just return dummy values
func getBooks() []book {

	return []book{
		{id: "0", bookName: "The Lord of the Rings: The Return of the King", numberOfPages: 355, stockCount: 1, stockCode: "123456", price: 50, ISBN: "ISBN1", author: author{AuthorID: "0", name: "J.R.R. Tolkien"}},
		{id: "1", bookName: "Hobbit", numberOfPages: 665, stockCount: 14, stockCode: "23456", price: 41, ISBN: "ISBN523", author: author{AuthorID: "0", name: "J.R.R. Tolkien"}},
		{id: "2", bookName: "The Unix Programming Environment", numberOfPages: 375, stockCount: 55, stockCode: "3456", price: 11, ISBN: "ISBN1", author: author{AuthorID: "1", name: "Rob Pike"}},
		{id: "3", bookName: "Beyaz Diş", numberOfPages: 285, stockCount: 3, stockCode: "456", price: 5, ISBN: "ISBN523", author: author{AuthorID: "2", name: "Jack London"}},
		{id: "4", bookName: "Palto", numberOfPages: 474, stockCount: 10, stockCode: "56", price: 10, ISBN: "ISBN51615", author: author{AuthorID: "3", name: "Vasilyeviç Gogol"}},
		{id: "5", bookName: "The Lord of The Rings: The Fellowship of the Ring", numberOfPages: 1558, stockCount: 32, stockCode: "65432", price: 24, ISBN: "ISBN5162615", author: author{AuthorID: "0", name: "J.R.R. Tolkien"}},
	}
}
