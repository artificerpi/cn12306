package main

var (
	cache Cache
)

func init() {
	cache = &SimpleCache{}
}
