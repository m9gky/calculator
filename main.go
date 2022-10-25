package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var mainTemplate = `
<html>
	<body>
		<form action='result'>
			<input name='a'>
			<select name='op'>
				<option value='add'>+</option>
				<option value='sub'>-</option>
				<option value='mul'>*</option>
				<option value='div'>/</option>
			</select>
			<input name='b'>
			<input type='submit'>
		</form>
	</body>
</html>
`

var resultTemplate = `
<html>
	<body>
		<p>Ответ: %v</p>
        <a href='/'>Ещё раз</a>
	</body>
</html>
`
var errorTemplate = `
<html>
    <body>
        <p>Ошибка в данных</p>
        <a href='/'>Назад</a>
    </body>
</html>
`

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/result", resultHandler)
	http.ListenAndServe("localhost:8000", nil)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, mainTemplate)
}

func resultHandler(w http.ResponseWriter, r *http.Request) {
	urlValues := r.URL.Query()
	op := urlValues.Get("op")
	a, errA := strconv.Atoi(urlValues.Get("a"))
	b, errB := strconv.Atoi(urlValues.Get("b"))
	if errA != nil || errB != nil {
		fmt.Fprintf(w, errorTemplate)
	} else if op == "add" {
		fmt.Fprintf(w, resultTemplate, a+b)
	} else if op == "sub" {
		fmt.Fprintf(w, resultTemplate, a-b)
	} else if op == "mul" {
		fmt.Fprintf(w, resultTemplate, a*b)
	} else if op == "div" {
		div := float32(a) / float32(b)
		fmt.Fprintf(w, resultTemplate, div)
	}
}
