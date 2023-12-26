package main

import (
	"github.com/k0new/storage"
	"log"
	"time"
)

func main() {
	s := storage.New()
	key := "example1"
	val := "some value"
	log.Printf("setting value %q to key %q\n", val, key)
	s.Set(key, val, 5*time.Second)
	log.Printf("retrieving value for key %q\n", key)
	log.Println(s.Get(key))
	log.Printf("deleting value for key %q\n", key)
	s.Delete(key)
	log.Printf("retrieving value for key %q\n", key)
	log.Println(s.Get(key))
}
