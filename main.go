package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	verbose := flag.Bool("v", false, "verbose output")
	filepath := flag.String("f", "", "output file path")
	forceWrite := flag.Bool("force-write", false, "write to output file without any checks")
	min := flag.Int("min", 0, "minimum brightness")
	max := flag.Int("max", 100, "maximum brightness")

	flag.Parse()

	now := time.Now()
	if *verbose {
		fmt.Println(now.Format(time.RFC1123Z))
	}

	minutes := now.Hour()*60 + now.Minute()
	if *verbose {
		fmt.Printf("Current time: %d\n", minutes)
	}

	magd2 := float64(*max-*min) / 2

	floatBrightness := magd2 - magd2*math.Cos(math.Pi*float64(minutes)/(12*60)) + float64(*min)

	brightness := int(floatBrightness)

	if *verbose {
		fmt.Printf("Calculated brightness: %d\n", brightness)
	} else {
		fmt.Print(brightness)
	}

	if *filepath == "" {
		return
	}

	if *forceWrite {
		writeBrightness(*filepath, brightness)
		return
	}

	currentBrightness, err := getCurrentBrightness(*filepath)
	if err != nil {
		log.Println("Failed to read current brightness:", err)
		writeBrightness(*filepath, brightness)
		return
	}

	if *verbose {
		fmt.Println("Current brightness:", currentBrightness)
	}

	if currentBrightness == int(brightness) {
		if *verbose {
			fmt.Println("Skipping file write.")
		}
		return
	}

	writeBrightness(*filepath, brightness)
}

func writeBrightness(filepath string, brightness int) {
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0664)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	_, err = f.WriteString(strconv.Itoa(int(brightness)))
	if err != nil {
		log.Fatal(err)
	}
}

func getCurrentBrightness(filepath string) (int, error) {
	content, err := os.ReadFile(filepath)
	if err != nil {
		return -1, err
	}

	currentBrightness, err := strconv.Atoi(strings.TrimSpace(string(content)))
	if err != nil {
		return -1, err
	}

	return currentBrightness, nil
}
