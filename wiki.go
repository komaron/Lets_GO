package main

import (
	"fmt"
	"io/ioutil"
)

// Data Structur declarations
type Page struct {
	Title string
	Body []byte
}