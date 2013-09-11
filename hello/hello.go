package hello

import (
	"appengine"
	"appengine/datastore"
	//	"appengine/user"
	"fmt"
	"net/http"
)

type Person struct {
	Name string
}

func init() {
	http.HandleFunc("/", RootHandler)
	http.HandleFunc("/add", AddHandler)
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	q := datastore.NewQuery("Person")
	var people []Person
	_, err := q.GetAll(c, &people)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var body = ""
	for _, p := range people {
		body += p.Name + "<br />"
	}
	body += `
<a href="/add">add</a><br />
`
	fmt.Fprint(w, body)
}

func AddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		c := appengine.NewContext(r)

		c.Infof(string(r.FormValue("name")))
		e1 := Person {
			Name: string(r.FormValue("name")),
		}
		_, err := datastore.Put(c, datastore.NewIncompleteKey(c, "Person", nil), &e1)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fmt.Fprint(w, `<a href="/">Go to root</a>`)
	} else {
		var body = `
<html><body>
<form action="/add" method="POST">
    <input type="text name="name" /><br />
    <input type="submit" />
</form>
</body></html>
`
		fmt.Fprint(w, body)
	}
}
