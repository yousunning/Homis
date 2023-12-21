package golang

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// Book 구조체 정의
type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

// 초기 데이터 설정
func init() {
	books = append(books, Book{ID: 1, Title: "Book 1", Author: "Author 1"})
	books = append(books, Book{ID: 2, Title: "Book 2", Author: "Author 2"})
	// ... 추가 데이터
}

// 모든 책 조회 (Read)
func getAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// 특정 ID의 책 조회 (Read)
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range books {
		if fmt.Sprintf("%d", item.ID) == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{}) // 해당 ID의 책이 없을 경우 빈 응답 반환
}

// 새로운 책 생성 (Create)
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = len(books) + 1
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// 책 정보 업데이트 (Update)
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if fmt.Sprintf("%d", item.ID) == params["id"] {
			var updatedBook Book
			_ = json.NewDecoder(r.Body).Decode(&updatedBook)
			updatedBook.ID = item.ID
			books[index] = updatedBook
			json.NewEncoder(w).Encode(updatedBook)
			return
		}
	}
	json.NewEncoder(w).Encode(&Book{}) // 해당 ID의 책이 없을 경우 빈 응답 반환
}

// 책 삭제 (Delete)
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if fmt.Sprintf("%d", item.ID) == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

func main() {
	// Gorilla Mux 라우터 생성
	router := mux.NewRouter()

	// 엔드포인트 및 핸들러 등록
	router.HandleFunc("/books", getAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/books", createBook).Methods("POST")
	router.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	// 서버 시작
	fmt.Println("Server is running on :8080")
	http.ListenAndServe(":8080", router)
}

// 모든 책 조회  : curl http://localhost:8080/books
// 특정 책 조회 : curl http://localhost:8080/books/1
// 책 생성 : curl -X POST -H "Content-Type: application/json" -d '{"title": "New Book", "author": "New Author"}' http://localhost:8080/books
// 책 업데이트 : curl -X PUT -H "Content-Type: application/json" -d '{"title": "Updated Book", "author": "Updated Author"}' http://localhost:8080/books/1
// 책 삭제 : curl -X DELETE http://localhost:8080/books/1
