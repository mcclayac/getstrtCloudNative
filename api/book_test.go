package api

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBook_toJSON(t *testing.T) {
	book := Book{
		Title:"Cloud Native Go",
		Author:"M.L.Reimer",
		ISBN:"0123456",
	}
	json := book.toJSON()

	//bookeq := Book{Ti\}

	//log.Println(string(json))
	assert.Equal(t,`{"title":"Cloud Native Go","author":"M.L.Reimer","isbn":"0123456"}`,
		string(json), "Book JSON marshalling is wrong")
}

func TestBookFromJSON(t *testing.T) {

	json := []byte(`{"title":"Cloud Native Go","author":"M.L.Reimer","isbn":"0123456"}`)
	book := FromJson(json)

	bookeq := Book{
		Title:"Cloud Native Go",
		Author:"M.L.Reimer",
		ISBN:"0123456",
	}
	//bookeq_json := bookeq.toJSON()
	//log.Println(bookeq_json)
	//bookeq := Book{Ti\}

	//log.Println(string(json))
	assert.Equal(t,bookeq,
		book, "Book JSON marshalling is wrong")

	assert.Equal(t, Book{Title:"Cloud Native Go", Author:"M.L.Reimer",ISBN:"0123456"},
		book, "Book JSON Unmarshalling wrong")
}






/*func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"Title":"Cloud Native Go","Author":"M.L.Reimer","ISBN":"0123456"}`)
	Book := FromJson(json)



	//assert.Equal(t, )
	bookeq := Book{}

	bookeq = bookeq
	//bookeq := Book{
	//	Title:"Cloud Native Go",
	//	Author:"M.L.Reimer",
	//	ISBN:"0123456",
	//}
}*/

