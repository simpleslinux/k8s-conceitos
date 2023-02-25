package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"
)

type key int

const (
	requestIDKey key = 0
)

var (
	listenAddr string
	healthy    int32
)

func main() {
	flag.StringVar(&listenAddr, "listen-addr", ":5000", "server listen address")
	flag.Parse()

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("Server is starting...")

	router := http.NewServeMux()
	router.Handle("/", index())
	router.Handle("/healthz", healthz())

	nextRequestID := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}

	server := &http.Server{
		Addr:         listenAddr,
		Handler:      tracing(nextRequestID)(logging(logger)(router)),
		ErrorLog:     logger,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		logger.Println("Server is shutting down...")
		atomic.StoreInt32(&healthy, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
		}
		close(done)
	}()

	logger.Println("Server is ready to handle requests at", listenAddr)
	atomic.StoreInt32(&healthy, 1)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}

	<-done
	logger.Println("Server stopped")
}

func index() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}
		msg := os.Getenv("CONTEUDO")
		if msg == "" {
			msg = "Teste Container"
		}
		hostname,_ := os.Hostname()
		html := `<!DOCTYPE html><html lang="pt_BR" dir="ltr"><head><style>@import url(https://fonts.googleapis.com/css2?family=Poppins:wght@200;300;400;500;600;700&display=swap);*{margin:0;padding:0;box-sizing:border-box;font-family:Poppins,sans-serif}::selection{color:#000;background:#fff}nav{position:fixed;background:#1b1b1b;width:100%;padding:10px 0;z-index:12}nav .menu{max-width:1250px;margin:auto;display:flex;align-items:center;justify-content:space-between;padding:0 20px}.menu .logo a{text-decoration:none;color:#fff;font-size:35px;font-weight:600}.menu ul{display:inline-flex}.img{width:100%;height:100vh;background-size:cover;background-position:center;position:relative}.img::before{content:'';position:absolute;height:100%;width:100%;background:rgba(0,0,0,.8)}.center{position:absolute;top:52%;left:50%;transform:translate(-50%,-50%);width:100%;padding:0 20px;text-align:center}.center .title{color:#fff;font-size:55px;font-weight:600}.center .sub_title{color:#fff;font-size:52px;font-weight:600}</style><meta charset="UTF-8"><meta name="viewport" content="width=device-width,initial-scale=1"><title>Website | K8S Conceitos</title></head><body><nav><div class="menu"><div class="logo"><a href="#">K8S Conceitos</a></div></div></nav><div class="img"></div><div class="center"><div class="title">` + msg + `</div><div class="sub_title">Host: ` + hostname + `</div></div></body></html>`
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintln(w, html)
	})
}

func healthz() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if atomic.LoadInt32(&healthy) == 1 {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		w.WriteHeader(http.StatusServiceUnavailable)
	})
}

func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				requestID, ok := r.Context().Value(requestIDKey).(string)
				if !ok {
					requestID = "unknown"
				}
				logger.Println(requestID, r.Method, r.URL.Path, r.RemoteAddr, r.UserAgent())
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func tracing(nextRequestID func() string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get("X-Request-Id")
			if requestID == "" {
				requestID = nextRequestID()
			}
			ctx := context.WithValue(r.Context(), requestIDKey, requestID)
			w.Header().Set("X-Request-Id", requestID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}