# Number Classification API

## Overview
The **Number Classification API** is a simple RESTful service that takes a number as input and returns mathematical properties about it, along with a fun fact. The API supports classification of numbers into prime, perfect, Armstrong, even, or odd.

### Example Usage:
```
GET /api/classify-number?number=371
```

### Example JSON Response:
```json
{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["armstrong", "odd"],
    "digit_sum": 11,
    "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
}
```

## Features
✅ Accepts a number as input via a GET request
✅ Determines if the number is prime, perfect, Armstrong, even, or odd
✅ Calculates the sum of its digits
✅ Fetches a fun fact about the number from [Numbers API](http://numbersapi.com/)
✅ Returns responses in JSON format
✅ CORS enabled for cross-origin requests

## Tech Stack
- **Language:** Go (Golang)
- **Framework:** Gin (for HTTP routing)
- **Hosting:** Render

## Installation & Setup
### 1️⃣ Clone the Repository
```sh
git clone https://github.com/nhzi/number-classifier.git
cd number-classifier
```

### 2️⃣ Install Dependencies
```sh
go mod tidy
```

### 3️⃣ Run the API Locally
```sh
go run main.go
```

### 4️⃣ Test the API
Visit:
```
http://localhost:8080/api/classify-number?number=371
```

## Deployment
### Deploy to Railway (Recommended)
1. Install Railway CLI:
   ```sh
   curl -fsSL https://railway.app/install.sh | sh
   railway login
   ```
2. Deploy:
   ```sh
   railway up
   ```

Alternatively, you can deploy using **Render, Fly.io, or Google Cloud Run**.

## API Endpoints
### ➤ Classify a Number
**Request:**
```
GET /api/classify-number?number=<number>
```

**Successful Response (200 OK):**
```json
{
    "number": 371,
    "is_prime": false,
    "is_perfect": false,
    "properties": ["armstrong", "odd"],
    "digit_sum": 11,
    "fun_fact": "371 is an Armstrong number because 3^3 + 7^3 + 1^3 = 371"
}
```

**Error Response (400 Bad Request):**
```json
{
    "number": "invalid",
    "error": true
}
```

## Contributing
1. Fork the repository
2. Create a feature branch (`git checkout -b feature-xyz`)
3. Commit your changes (`git commit -m "Add feature XYZ"`)
4. Push to the branch (`git push origin feature-xyz`)
5. Create a Pull Request

## License
This project is licensed under the MIT License.

---

🌟 **Developed by Ved.dev; HNG Internship** 🚀

