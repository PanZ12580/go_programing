/**
 * @Author $
 * @Description //TODO $
 * @Date $ $
 * @Param $
 * @return $
 **/
package handler

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

type Database map[string]dollars

type dollars float64

var temp = `
<table>
	<tr style="border:1px solid #C1437A;">
		<th style="width: 10px;border:1px solid #C1437A;">item</td>
		<th style="width: 10px;border:1px solid #C1437A;">price</td>
	</tr>
	{{range $key, $value := .}}
	<tr style="border:1px solid #C1437A;">
		<td style="width: 10px;border:1px solid #C1437A;">{{$key}}</td>
		<td style="width: 10px;border:1px solid #C1437A;">{{$value}}</td>
	</tr>
	{{end}}
</table>`

func (d dollars) String() string {
	return fmt.Sprintf("$%.2f\n", d)
}

func (db Database) List(w http.ResponseWriter, r *http.Request) {
	t, err := template.
		New("list").
		Parse(temp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "internal error: %s\n", err)
	} else {
		err = t.Execute(w, db)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "internal error: %s\n", err)
		}
	}
}

func (db Database) QueryPrice(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "invalid parameter\n")
	} else {
		p, ok := db[item]
		if ok {
			fmt.Fprintf(w, "item: %s, price: %s\n", item, p)
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "item not found: %s\n", item)
		}
	}
}

func (db Database) DeleteItem(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	if item == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "invalid parameter\n")
	} else {
		_, ok := db[item]
		if ok {
			delete(db, item)
			fmt.Fprintf(w, "delete item ok: %s\n", item)
		} else {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "item not found: %s\n", item)
		}
	}
}

func (db Database) AddItem(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")
	if item == "" || price == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "invalid parameter\n")
	} else {
		fPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "invalid price value: %s\n", price)
		} else {
			db[item] = dollars(fPrice)
			fmt.Fprintf(w, "add new item successfully!---------item: %s, price: %s\n", item, dollars(fPrice))
		}
	}
}

func (db Database) UpdateItem(w http.ResponseWriter, r *http.Request) {
	item := r.URL.Query().Get("item")
	price := r.URL.Query().Get("price")
	if item == "" || price == "" {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "invalid parameter\n")
	} else {
		fPrice, err := strconv.ParseFloat(price, 64)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "invalid price value: %s\n", price)
		} else {
			_, ok := db[item]
			if !ok {
				w.WriteHeader(http.StatusNotFound)
				fmt.Fprintf(w, "item not found: %s\n", item)
			} else {
				db[item] = dollars(fPrice)
				fmt.Fprintf(w, "update item successfully!---------item: %s, price: %s\n", item, dollars(fPrice))
			}
		}
	}
}


