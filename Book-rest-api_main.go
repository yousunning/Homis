package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http" // https://pkg.go.dev/net/http
	"sync"
	//서버 및 클라이언트 기능을 제공하는 Go의 표준 라이브러리
	// net/http를 임포트
)

// Unmarshal 과정 : 특정 형식의 데이터를 프로그램이 사용할 수 있는 구조나 객체로 변환하는 과정

// Book 구조체 정의
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
} //*포인터 값 / &주소 /

var ( // 전역변수
	books = make(map[string]Book) //이 맵은 책들의 정보를 저장 // map 초기화
	mutex = &sync.Mutex{}         //string 키와 *Book 값으로 구성된 새로운 맵을 생성
)

//mutex는 sync.Mutex 타입의 변수로, Go 언어의 표준 라이브러리에 있는 동기화 기능을 제공
//&sync.Mutex{}는 sync.Mutex의 새 인스턴스를 생성하고 그 주소를 mutex 변수에 할당

// (r *http.Request)을 분석하여 필요한 정보를 추출
// http.ResponseWriter (w)를 사용하여 응답을 전송

func main() { //book 경로에 대한 요청을 처리할 핸들러 함수를 등록/book 경로에 대한 요청을 처리할 핸들러 함수를 등록
	// http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
	// 	w.Write([]byte("Hello, World!"))
	// 	//이 함수는 응답으로 "Hello, World!" 문자열을 클라이언트에게 보냄
	// })
	// //000번 포트에서 HTTP 서버를 시작합니다.
	// //nil은 기본 HTTP 핸들러를 사용하겠다는 의미
	// err := http.ListenAndServe(":8000", nil)
	// if err != nil {
	// 	panic(err)
	// }

	// data := `[
	// 								{"id":"001","title":"트렌드 코리아 2024","author":"김난도"},
	// 								{"id":"002","title":"뉴 트렌드 2024","author":"이영희"},
	// 								{"id":"003","title":"미래의 도전","author":"박철수"}
	// 								]`

	// //JSON 데이터가 배열 형태로 되어 있기 때문에 배열을 처리할 수 있는 슬라이스로 Book을 정의
	// var books []Book //
	// err := json.Unmarshal([]byte(data), &books)
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Printf("%+v\n", books)
	//https://pkg.go.dev/search?q=HandleFunc&m=symbol
	http.HandleFunc("/books", booksHandler)
	//책의 상세 정보를 조회하거나 수정하는 작업을 수행
	http.HandleFunc("/book/", bookHandler)
	//https://pkg.go.dev/net/http#hdr-Servers
	log.Println("Starting web server at http://localhost:8000/")
	//어떤 주소(http://localhost:8000)에서 서비스되고 있는지 콘솔에 출력
	log.Fatal(http.ListenAndServe(":8000", nil))
	//log.Fatal은 http.ListenAndServe 함수에서 반환되는 에러를 로깅하며,
	//에러가 있을 경우 프로그램을 종료
	// form>input*3+select>option*4^input*4+input
}

func booksHandler(w http.ResponseWriter, r *http.Request) { //*메모리주소를 가리킨다
	switch r.Method { //클라이언트의 요청 메소드(GET, POST 등)를 확인하는 부분
	case "GET":
		getBooks(w, r) //모든 책의 목록을 클라이언트에게 반환
	case "POST":
		createBook(w, r) //함수를 호출하여 새로운 책을 추가
	case "DELETE":
		deleteBooks(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	} //에러 상태 코드를 반환
}

func bookHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		// URL에서 책의 ID를 파싱
		id := r.URL.Path[len("/book/"):] // 예: URL이 /book/1234 라면, "1234"를 추출

		// 책 정보를 찾는 로직
		book, exists := books[id] // 'books'는 책 정보를 저장하는 데이터 구조
		if !exists {
			//만약 해당 ID에 해당하는 책이 존재하지 않으면,
			//HTTP 응답에 "Book not found" 메시지와 함께 404 상태 코드를 반환합니다.
			http.Error(w, "Book not found", http.StatusNotFound)
			return
		}
		// HTTP 응답의 헤더를 설정하여 클라이언트에게 전송할 데이터가 JSON 형식임을 알립니다.
		//해당 정보를 HTTP 응답으로 전송합니다. 'w'는 HTTP 응답 작성을 위한 ResponseWriter를 나타냅니다.
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(book)

		// 특정 책 ID에 해당하는 책을 찾아서 해당 정보를 JSON 형식으로
		//클라이언트에게 반환하는 웹 서버의 핸들러 부분입니다.
		//만약 책이 존재하지 않으면 404 상태 코드와 함께 "Book not found" 메시지를 반환

	case "PUT":
		updateBook(w, r) // PUT 요청 처리
	case "DELETE":
		deleteBook(w, r) // DELETE 요청 처리
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// CRUD 함수들 정의
func getBooks(w http.ResponseWriter, r *http.Request) {
	// 함수의 시작 부분에서 뮤텍스를 잠급니다.
	// 이는 books 맵에 대한 동시 접근을 방지하여 동시성 문제를 예방
	mutex.Lock() //https://pkg.go.dev/sync#Mutex.Lock
	//mutex.Lock()은 공유 자원에 대한 동시 액세스를 제어하기 위한 동기화 메커니즘 중 하나인 뮤텍스(Mutex)를 사용하는 함수

	//함수의 종료 시점에 뮤텍스를 자동으로 해제하기 위해 defer 키워드를 사용
	//이는 함수가 끝날 때(예를 들어, 모든 책 정보가 반환된 후) 반드시 뮤텍스가 해제
	defer mutex.Unlock()

	var bookList []Book          //책 포인터들을 저장할 슬라이스를 초기화
	for _, book := range books { //books 맵에 저장된 모든 책들을 순회
		bookList = append(bookList, book) //책을 bookList 슬라이스에 추가
	}

	//클라이언트에게 반환되는 내용이 JSON 형식 명시
	w.Header().Set("Content-Type", "application/json")

	//bookList 슬라이스를 JSON 형식으로 인코딩하고, 이를 HTTP 응답으로 클라이언트에게 전송
	json.NewEncoder(w).Encode(bookList)
}

func createBook(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	// var  Book
	var bookDatas []Book
	if err := json.NewDecoder(r.Body).Decode(&bookDatas); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// if _, exists := books[book.ID]; exists {
	// 	http.Error(w, "Book already exists", http.StatusBadRequest)
	// 	return
	// }
	//book.ID는 책의 고유 식별자로 사용
	for _, b := range bookDatas {
		books[b.ID] = b //&book는 해당 책 객체의 주소(포인터)
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(bookDatas)
	//books 맵에 새로운 책이 추가되거나, 기존의 책 정보가 업데이트

	// 새로운 책을 추가하고 그 정보를 클라이언트에게 반환하는 역할을 합니다.
	// 클라이언트는 이 응답을 통해 새로 생성된 책의 정보를 확인할 수 있습니다.

}
func updateBook(w http.ResponseWriter, r *http.Request) {
	// URL에서 책의 ID 추출
	id := r.URL.Path[len("/book/"):]

	// 요청 본문에서 업데이트할 책 정보를 디코딩
	var updatedBook Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 잠금을 사용하여 동시성 문제를 방지
	mutex.Lock()
	defer mutex.Unlock()

	// 책이 존재하는지 확인
	book, exists := books[id]
	if !exists { // 책이 맵에 없을 때 실행되는 코드
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// 책 정보 업데이트
	book.Title = updatedBook.Title
	book.Author = updatedBook.Author
	books[id] = book

	// 응답
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(book)
}
func deleteBook(w http.ResponseWriter, r *http.Request) {
	// URL에서 책의 ID를 추출
	id := r.URL.Path[len("/book/"):]

	// 잠금을 사용하여 동시성 문제를 방지
	mutex.Lock()
	defer mutex.Unlock()

	// 책이 존재하는지 확인
	_, exists := books[id]
	if !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// 책 삭제
	delete(books, id)

	// 성공적으로 삭제됨을 알림
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted")
}
func deleteBooks(w http.ResponseWriter, r *http.Request) {
	// URL에서 책의 ID를 추출
	// id := r.URL.Path[len("/book/"):]

	// 잠금을 사용하여 동시성 문제를 방지
	mutex.Lock()
	defer mutex.Unlock()

	// // 책이 존재하는지 확인
	// _, exists := books[id]
	// if !exists {
	// 	http.Error(w, "Book not found", http.StatusNotFound)
	// 	return
	// }
	var bookIds []string
	if err := json.NewDecoder(r.Body).Decode(&bookIds); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// 책 삭제
	for _, b := range bookIds {
		delete(books, b)
	}
	// delete(books, id)

	// 성공적으로 삭제됨을 알림
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Deleted")
}
