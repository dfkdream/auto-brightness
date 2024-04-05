package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"time"
	"log"
	"strconv"
)

func main(){
	verbose := flag.Bool("v", false, "verbose output")
	filepath := flag.String("f", "", "output file path")
	min:=flag.Int("min", 0, "minimum brightness")
	max:=flag.Int("max", 100, "maximum brightness")

	flag.Parse()

	now := time.Now()
	if *verbose{
		fmt.Println(now.Format(time.RFC1123Z))
	}

	minutes := now.Hour()*60 + now.Minute()
	if *verbose{
		fmt.Printf("Current time: %d\n", minutes)
	}

	magd2 := float64(*max - *min)/2

	brightness := magd2 - magd2 * math.Cos(math.Pi*float64(minutes)/(12*60)) + float64(*min)

	if *verbose{
		fmt.Printf("Calculated brightness: %d\n", int(brightness))
	}

	if !*verbose && *filepath == ""{
		fmt.Print(int(brightness))
	}

	if *filepath != ""{
		f, err := os.OpenFile(*filepath, os.O_WRONLY|os.O_TRUNC, 664)
		if err!=nil{
			log.Fatal(err)
		}

		defer f.Close()

		_, err = f.WriteString(strconv.Itoa(int(brightness)))
		if err!=nil{
			log.Fatal(err)
		}
	}
}
