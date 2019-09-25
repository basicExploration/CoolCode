package main

import (
	"fmt"
	"testing"
)

func TestPackageOne(t *testing.T) {
	fmt.Print("one",PackageOne())// one 25
}

func TestPackageTwo(t *testing.T) {
	fmt.Println("two",PackageTwo()) // two 396
}