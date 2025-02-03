package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Function to check if a number is prime
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// Function to check if a number is perfect
func isPerfect(n int) bool {
	sum := 1
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			sum += i
			if i != n/i {
				sum += n / i
			}
		}
	}
	return sum == n && n != 1
}

// Function to check if a number is Armstrong (Narcissistic)
func isArmstrong(n int) bool {
	temp, sum, digits := n, 0, len(strconv.Itoa(n))
	for temp > 0 {
		digit := temp % 10
		sum += int(math.Pow(float64(digit), float64(digits)))
		temp /= 10
	}
	return sum == n
}

// Function to calculate digit sum
func digitSum(n int) int {
	sum := 0
	for n > 0 {
		sum += n % 10
		n /= 10
	}
	return sum
}

// Function to get a fun fact from Numbers API
func getFunFact(n int) string {
	url := fmt.Sprintf("http://numbersapi.com/%d?json", n)
	resp, err := http.Get(url)
	if err != nil {
		return "Could not fetch fun fact."
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	var result map[string]interface{}
	json.Unmarshal(body, &result)

	if fact, ok := result["text"].(string); ok {
		return fact
	}
	return "No fun fact available."
}

// API handler function
func classifyNumber(c *gin.Context) {
	numberStr := c.Query("number")
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"number": numberStr,
			"error":  true,
		})
		return
	}

	// Determine number properties
	properties := []string{}
	if isPrime(number) {
		properties = append(properties, "prime")
	}
	if isPerfect(number) {
		properties = append(properties, "perfect")
	}
	if isArmstrong(number) {
		properties = append(properties, "armstrong")
	}
	if number%2 == 0 {
		properties = append(properties, "even")
	} else {
		properties = append(properties, "odd")
	}

	// Prepare response
	response := gin.H{
		"number":     number,
		"is_prime":   isPrime(number),
		"is_perfect": isPerfect(number),
		"properties": properties,
		"digit_sum":  digitSum(number),
		"fun_fact":   getFunFact(number),
	}

	c.JSON(http.StatusOK, response)
}

func main() {
	router := gin.Default()

	// Enable CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		c.Next()
	})

	router.GET("/api/classify-number", classifyNumber)

	// Start server on port 8080
	router.Run(":8080")
}
