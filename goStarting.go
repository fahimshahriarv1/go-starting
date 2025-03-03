package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	const formulaMsg = "BMI = weight (kg) / height m²"
	fmt.Println(formulaMsg)

	var rd = bufio.NewReader(os.Stdin)

	fmt.Print("Enter weight (kg): ")
	var w, _ = rd.ReadString('\n')
	w = strings.TrimSpace(w)
	
	fmt.Print("Enter your height in F' In\" format \n")

	fmt.Print("Enter foot portion of ur height the : ")
	var hF, _ = rd.ReadString('\n')

	fmt.Print("Enter inch portion of ur height the : ")
	var hIn, _ = rd.ReadString('\n')

	hIn = strings.TrimSpace(hIn)
	hF = strings.TrimSpace(hF)

	wW, _ := strconv.ParseFloat(w, 32)
	hInn, _ := strconv.ParseFloat(hIn, 32)
	hFf, _ := strconv.ParseFloat(hF, 32)

	var bmi = wW/(((hFf*12+hInn)/39.37) * ((hFf*12+hInn)/39.37))

	fmt.Println("Your BMI is = ", bmi)

	if bmi < 18.5 {
		fmt.Println("Health Status: Underweight")
	} else if bmi >= 18.5 && bmi < 24.9 {
		fmt.Println("Health Status: Normal weight")
	} else if bmi >= 24.9 && bmi < 29.9 {
		fmt.Println("Health Status: Overweight")
	} else {
		fmt.Println("Health Status: Obese")
	}
}
