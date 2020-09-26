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

const (
	sortenedURLsPath = "/surls/"
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

func main() {
	log.Print("Start service")

	http.HandleFunc("/", indexHandler)
	http.HandleFunc(sortenedURLsPath, shortenedurlsHandler)

	if err := http.ListenAndServe(":8888", nil); err != nil {
		log.Fatalf("Failed: %v", err)
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
		fmt.Fprint(w, "Bad request")
		return
	}

	if r.Method == http.MethodGet {
		fmt.Fprint(w, indexHtml)
		return
	}

	fmt.Fprint(w, "Bad request")
}

func shortenedurlsHandler(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, sortenedURLsPath) {
		fmt.Fprint(w, "Bad request")
		return
	}

	if r.Method == http.MethodGet {
		sURL := strings.TrimPrefix(r.URL.Path, sortenedURLsPath)

		// KVSからsURLで検索して、リダイレクト
		var ctx = context.Background()
		realURL, err := redisClient.Get(ctx, sURL).Result()
		if err != nil {
			if err == redis.Nil {
				fmt.Fprint(w, "key does not exists")
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
			fmt.Fprint(w, "empty url value")
			return
		}

		// 短縮URL用の文字列を生成して、KVSに登録し、短縮URLを返却
		uid, err := uuid.NewRandom()
		if err != nil {
			fmt.Fprint(w, err)
			return
		}
		ust := uid.String()
		genPath := strings.Split(ust, "-")[0]

		var ctx = context.Background()
		err = redisClient.Set(ctx, genPath, realURL, 0).Err()
		if err != nil {
			fmt.Fprint(w, err)
			return
		}

		fmt.Fprint(w, host+sortenedURLsPath+genPath+"\n")
		return
	}

	fmt.Fprint(w, "Bad request")
}
