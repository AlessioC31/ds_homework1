package main

import (
	"encoding/json"
	"encoding/xml"
	"net/http"

	pb "github.com/AlessioC31/ds_homework1_server/protos"
	"google.golang.org/protobuf/proto"
)

func main() {
	http.HandleFunc("/json/tiny", jsonTiny)
	http.HandleFunc("/json/medium", jsonMedium)
	http.HandleFunc("/json/large", jsonLarge)

	http.HandleFunc("/xml/tiny", xmlTiny)
	http.HandleFunc("/xml/medium", xmlMedium)
	http.HandleFunc("/xml/large", xmlLarge)

	http.HandleFunc("/protobuf/tiny", xmlTiny)
	http.HandleFunc("/protobuf/medium", xmlMedium)
	http.HandleFunc("/protobuf/large", xmlLarge)

	http.ListenAndServe(":8080", nil)
}

func jsonTiny(w http.ResponseWriter, r *http.Request) {
	a := Author{"Mario", "Rossi"}
	json, err := json.Marshal(a)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func jsonMedium(w http.ResponseWriter, r *http.Request) {
	b := Book{
		[]Author{
			Author{
				"Mario",
				"Rossi",
			},
		},
		Publisher{
			"Mondadori",
		},
		Date{
			5,
			4,
			2010,
		},
		500,
		Available,
	}

	json, err := json.Marshal(b)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func jsonLarge(w http.ResponseWriter, r *http.Request) {
	b1 := Book{
		[]Author{
			Author{
				"Mario",
				"Rossi",
			},
		},
		Publisher{
			"Mondadori",
		},
		Date{
			1,
			1,
			1972,
		},
		140,
		Available,
	}

	b2 := Book{
		[]Author{
			Author{
				"Maria",
				"Rossi",
			},
		},
		Publisher{
			"Springer",
		},
		Date{
			12,
			12,
			2012,
		},
		1500,
		NotAvailable,
	}

	b3 := Book{
		[]Author{
			Author{
				"Giuseppe",
				"Rossi",
			},
		},
		Publisher{
			"O'Reilly",
		},
		Date{
			5,
			4,
			2005,
		},
		123,
		Available,
	}

	l := Library{
		[]Book{
			b1,
			b2,
			b3,
		},
		Time{9, 0},
		Time{18, 0},
	}

	json, err := json.Marshal(l)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(json)
}

func xmlTiny(w http.ResponseWriter, r *http.Request) {
	a := Author{"Mario", "Rossi"}
	xml, err := xml.Marshal(a)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(xml)
}

func xmlMedium(w http.ResponseWriter, r *http.Request) {
	b := Book{
		[]Author{
			Author{
				"Mario",
				"Rossi",
			},
		},
		Publisher{
			"Mondadori",
		},
		Date{
			5,
			4,
			2010,
		},
		500,
		Available,
	}

	xml, err := xml.Marshal(b)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(xml)
}

func xmlLarge(w http.ResponseWriter, r *http.Request) {
	b1 := Book{
		[]Author{
			Author{
				"Mario",
				"Rossi",
			},
		},
		Publisher{
			"Mondadori",
		},
		Date{
			1,
			1,
			1972,
		},
		140,
		Available,
	}

	b2 := Book{
		[]Author{
			Author{
				"Maria",
				"Rossi",
			},
		},
		Publisher{
			"Springer",
		},
		Date{
			12,
			12,
			2012,
		},
		1500,
		NotAvailable,
	}

	b3 := Book{
		[]Author{
			Author{
				"Giuseppe",
				"Rossi",
			},
		},
		Publisher{
			"O'Reilly",
		},
		Date{
			5,
			4,
			2005,
		},
		123,
		Available,
	}

	l := Library{
		[]Book{
			b1,
			b2,
			b3,
		},
		Time{9, 0},
		Time{18, 0},
	}

	xml, err := xml.Marshal(l)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/xml")
	w.Write(xml)
}

func protobufTiny(w http.ResponseWriter, r *http.Request) {
	author := &pb.Author{
		Name:    "Mario",
		Surname: "Rossi",
	}

	out, err := proto.Marshal(author)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "x-protobuf")
	w.Write(out)
}

func protobufMedium(w http.ResponseWriter, r *http.Request) {
	book := &pb.Book{
		Authors: []*pb.Author{
			&pb.Author{
				Name:    "Mario",
				Surname: "Rossi",
			},
		},
		Date: &pb.Date{
			Day:   5,
			Month: 4,
			Year:  2010,
		},
		Pages: 500,
		Publisher: &pb.Publisher{
			Name: "Mondadori",
		},
		State: pb.Book_AVAILABLE,
	}

	out, err := proto.Marshal(book)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "x-protobuf")
	w.Write(out)
}

func protobufLarge(w http.ResponseWriter, r *http.Request) {
	book1 := &pb.Book{
		Authors: []*pb.Author{
			&pb.Author{
				Name:    "Mario",
				Surname: "Rossi",
			},
		},
		Date: &pb.Date{
			Day:   1,
			Month: 1,
			Year:  1972,
		},
		Pages: 140,
		Publisher: &pb.Publisher{
			Name: "Mondadori",
		},
		State: pb.Book_AVAILABLE,
	}

	book2 := &pb.Book{
		Authors: []*pb.Author{
			&pb.Author{
				Name:    "Maria",
				Surname: "Rossi",
			},
		},
		Date: &pb.Date{
			Day:   12,
			Month: 12,
			Year:  2012,
		},
		Pages: 1500,
		Publisher: &pb.Publisher{
			Name: "Springer",
		},
		State: pb.Book_AVAILABLE,
	}

	book3 := &pb.Book{
		Authors: []*pb.Author{
			&pb.Author{
				Name:    "Giuseppe",
				Surname: "Rossi",
			},
		},
		Date: &pb.Date{
			Day:   5,
			Month: 4,
			Year:  205,
		},
		Pages: 123,
		Publisher: &pb.Publisher{
			Name: "O'Reilly",
		},
		State: pb.Book_NOT_AVAILABLE,
	}

	library := &pb.Library{
		Books: []*pb.Book{
			book1,
			book2,
			book3,
		},
		OpeningTime: &pb.Time{
			Hours:   9,
			Minutes: 0,
		},
		ClosingTime: &pb.Time{
			Hours:   18,
			Minutes: 0,
		},
	}

	out, err := proto.Marshal(library)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "x-protobuf")
	w.Write(out)
}
