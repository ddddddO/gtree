package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

var (
	host        string
	redisHost   string
	redisClient *redis.Client
)

func init() {
	host = os.Getenv("SHORTENED_URLS_SVC_HOST")
	if host == "" {
		host = "http://localhost"
	}

	redisHost = os.Getenv("REDIS_HOST")
	if redisHost == "" {
		redisHost = "localhost:6379"
	}

	redisClient = newRedisClient()
}

func newRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     redisHost,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

const (
	shortenedURLsPath = "/surls/"
)

func main() {
	log.Print("Start service")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc(shortenedURLsPath, shortenedurlsHandler)

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatalf("Failed to start service:\n%v", err)
	}
}

const indexHtml = `
<!DOCTYPE html>
<html>
<head>
	<title>Shortened URL Service</title>
</head>
<body>
	<form action="/surls/" method="post">
		<p>
			URL：<input type="text" name="url">
		</p>
	</form>
</body>
</html>
`

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	if r.Method != http.MethodGet {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	fmt.Fprint(w, indexHtml)
}

func shortenedurlsHandler(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, shortenedURLsPath) {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodGet {
		// KVSからsURLで検索して、リダイレクト
		sURL := strings.TrimPrefix(r.URL.Path, shortenedURLsPath)
		ctx := context.Background()
		realURL, err := redisClient.Get(ctx, sURL).Result()
		if err != nil {
			if err == redis.Nil {
				http.Error(w, "Not found shortened url", http.StatusNotFound)
				return
			}

			fmt.Fprint(w, err)
			return
		}

		http.Redirect(w, r, realURL, http.StatusSeeOther)
		return
	} else if r.Method == http.MethodPost {
		realURL := r.PostFormValue("url")
		if realURL == "" {
			http.Error(w, "Empty url value", http.StatusBadRequest)
			return
		}

		// 短縮URL用の文字列を生成して、KVSに登録し、短縮URLを返却
		uid, err := uuid.NewRandom()
		if err != nil {
			http.Error(w, "Could not generate shortened url", http.StatusInternalServerError)
			return
		}

		ust := uid.String()
		genPath := strings.Split(ust, "-")[0]
		ctx := context.Background()
		if err := redisClient.Set(ctx, genPath, realURL, 0).Err(); err != nil {
			http.Error(w, "Could not generate shortened url", http.StatusInternalServerError)
			return
		}

		fmt.Fprint(w, host+shortenedURLsPath+genPath+"\n")
		return
	} else {
		http.Error(w, "Bad request", http.StatusBadRequest)
	}
}
