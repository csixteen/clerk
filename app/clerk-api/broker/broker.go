// MIT License
//
// Copyright (c) 2020 Pedro Rodrigues
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package broker

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	d "github.com/csixteen/clerk/internal/database"
	"github.com/gorilla/mux"
)

// Broker responsible for binding the business logic of
// managing tasks and notes with the HTTP logic
type Broker struct {
	Db *sql.DB

	Router *mux.Router
}

func New() *Broker {
	db, err := d.SetupDatabase()
	if err != nil {
		log.Fatal(err)
	}

	r := mux.NewRouter().StrictSlash(true)

	b := &Broker{
		Router: r,
		Db:     db,
	}

	return b
}

func (b *Broker) Start(binder func(b *Broker)) {
	log.Println("Starting the broker. Listening on port 8888...")

	binder(b)

	srv := &http.Server{
		Handler: b.Router,
		Addr:    "127.0.0.1:8888",
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)

	// Graceful shutdown when CTRL+C (SINGINT)
	signal.Notify(c, os.Interrupt)

	// Block until we receive the signal
	<-c

	ctx, cancel := context.WithTimeout(
		context.Background(),
		15*time.Second,
	)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline
	srv.Shutdown(ctx)

	log.Println("shutting down...")
	os.Exit(0)
}
