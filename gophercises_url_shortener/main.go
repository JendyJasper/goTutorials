package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/boltdb/bolt"
	yaml "gopkg.in/yaml.v2"
)

//decleare a struct that will check the yaml
type pathURL struct {
	Path string
	URL  string
}

func readBoltHandler(urlMapdb string) map[string]string {
	db, err := bolt.Open(urlMapdb, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	urlDB := make(map[string]string)
	db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("urlBucket"))
		if bucket == nil {
			log.Fatal(err) // does not exist
		}
		c := bucket.Cursor()

		for k, v := c.First(); k != nil; k, v = c.Next() {
			urlDB[string(k)] = string(v)
		}
		return nil
	})
	return urlDB
}

func writeBoltHandler(urlMapdb string, pathUrls []pathURL) []pathURL {
	db, err := bolt.Open(urlMapdb, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte("urlBucket"))
		if bucket == nil {
			// Create bucket
			bucket, err = tx.CreateBucket([]byte("urlBucket"))
			if err != nil {
				log.Fatal(err)
				return err
			}
		}
		for _, path := range pathUrls {
			bucket.Put([]byte(path.Path), []byte(path.URL))
		}
		return nil
	})
	return pathUrls
}

func mpHandler(urlMap map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//check the incoming path
		path := r.URL.Path
		urlMap = readBoltHandler("url.db")
		if dest, ok := urlMap[path]; ok {
			//redirect if the incoming path matches the value in the map
			http.Redirect(w, r, dest, http.StatusFound)
		}
		fallback.ServeHTTP(w, r)
	}
}

func yamlHandler(yamlData string, fallback http.Handler) http.HandlerFunc {
	var pathUrls []pathURL
	//write the yml file into it slice of struct
	file, err := ioutil.ReadFile(yamlData)
	if err != nil {
		fmt.Println(err)
	}
	er := yaml.Unmarshal(file, &pathUrls)
	if er != nil {
		fmt.Println(er)
	}
	writeBoltHandler("url.db", pathUrls)
	mpYaml := readBoltHandler("url.db")

	//range through the slice of structs and use the values to make a map
	//return map handler and use
	return mpHandler(mpYaml, fallback)
}

func jsonHandler(jsonData []byte, fallback http.Handler) http.HandlerFunc {
	var pathUrls []pathURL
	err := json.Unmarshal(jsonData, &pathUrls)
	if err != nil {
		fmt.Println(err)
	}
	writeBoltHandler("url.db", pathUrls)
	mpJSON := readBoltHandler("url.db")
	return mpHandler(mpJSON, fallback)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Welcome to homepage</h1>")
}

func main() {
	yamlFile := flag.String("yaml", "paths.yaml", "Pass in the yaml that has the file configurations")
	flag.Parse()
	mux := http.NewServeMux()

	urlsdbStruct := []pathURL{{Path: "/jw", URL: "https://jw.org"}, {Path: "/quizra", URL: "https://quizragame.com"}}

	writeBoltHandler("url.db", urlsdbStruct)
	urlDb := readBoltHandler("url.db")

	jsonData := `
	[{
		"Path":"/jw",
		"URL":"https://jw.org"
	},
	{
		"Path":"/smart",
		"URL":"https://smartalaba.com"
	}]
	`

	mph := mpHandler(urlDb, mux)
	yh := yamlHandler(*yamlFile, mph)
	jsh := jsonHandler([]byte(jsonData), yh)

	mux.HandleFunc("/", index)

	fmt.Println("Server sterting on port 3000")
	http.ListenAndServe(":3000", jsh)
}

// package main

// import (
// 	"fmt"
// 	"net/http"

// 	yaml "gopkg.in/yaml.v2"
// )

// type pathURL struct {
// 	Path string `yaml:"path"`
// 	URL  string `yaml:"url"`
// }

// func defaultMux() *http.ServeMux {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", hello)
// 	return mux
// }

// func hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Hello, world!")
// }

// func mpHandler(urlMap map[string]string, fallback http.Handler) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		path := r.URL.Path
// 		if dest, ok := urlMap[path]; ok {
// 			http.Redirect(w, r, dest, http.StatusFound)
// 		} //else if r.URL.Path == "/" {
// 		// 	fmt.Fprintf(w, "<h1>Welcome to the homepage</h1>")
// 		// }
// 		fallback.ServeHTTP(w, r)
// 	}
// }

// func yamlHandler(yamlData []byte, fallback http.Handler) http.HandlerFunc {
// 	var pathUrls []pathURL
// 	mpYaml := make(map[string]string)
// 	err := yaml.Unmarshal(yamlData, &pathUrls)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	for _, v := range pathUrls {
// 		mpYaml[v.Path] = v.URL
// 	}
// 	return mpHandler(mpYaml, fallback)
// }

// func main() {
// 	mux := defaultMux()

// 	urlPaths := map[string]string{"/goog": "https://google.com", "/fb": "https://facebook.com"}

// 	yaml := `
//     - path: /urlshort
//       url: https://github.com/gophercises/urlshort
//     - path: /urlshort-final
//       url: https://github.com/gophercises/urlshort/tree/solution
// `

// 	mph := mpHandler(urlPaths, mux)
// 	yh := yamlHandler([]byte(yaml), mph)

// 	//mux.HandleFunc("/", hello)

// 	fmt.Println("Server sterting on port 8000")
// 	http.ListenAndServe(":8000", yh)
// }
