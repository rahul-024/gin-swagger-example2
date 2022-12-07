package db

import "github.com/rahul-024/gin-swagger-example2/models"

// Books slice to seed book data.
var Books = []models.Book{
	{ISBN: "9781612680194", Title: "Rich Dad Poor Dad", Author: "Robert Kiyosaki"},
	{ISBN: "9781781257654", Title: "The Daily Stotic", Author: "Ryan Holiday"},
	{ISBN: "9780593419052", Title: "A Mind for Numbers", Author: "Barbara Oklay"},
}
