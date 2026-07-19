# Golang Services

A collection of simple RESTful web services developed using **Go (Golang)** and the built-in **net/http** package. Each service demonstrates the fundamentals of building HTTP servers, handling requests, processing query parameters, and returning responses through a web browser.

This project was developed as part of hands-on learning in **Go programming**, **Web Services**, and **Backend Development**.

---

## 🚀 Technologies Used

- Go (Golang)
- net/http Package
- HTTP Protocol
- RESTful Web Services
- Git & GitHub

---

## 📂 Project Structure

```
Golang-Services/
│
├── calculator.go
├── greeting.go
├── helloworld.go
├── squarenumber.go
├── temperature.go
├── weather.go
├── factorial.go
├── evenodd.go
├── randomnumber.go
└── currenttime.go
```

---

## 📋 Services Included

| Service | Description |
|---------|-------------|
| Hello World Service | Displays a simple "Hello, World!" message |
| Greeting Service | Returns a personalized greeting using a name parameter |
| Square Number Service | Calculates the square of a given number |
| Temperature Converter | Converts Celsius to Fahrenheit |
| Calculator Service | Performs basic arithmetic operations |
| Weather Prediction Service | Displays sample weather information |
| Factorial Service | Calculates the factorial of a given number |
| Even/Odd Service | Checks whether a number is even or odd |
| Random Number Service | Generates a random number |
| Current Time Service | Displays the current system date and time |

---

## ▶️ How to Run

Clone the repository:

```bash
git clone https://github.com/grazpaulx/Golang-Services.git
```

Navigate to the project directory:

```bash
cd Golang-Services
```

Run any Go service:

```bash
go run filename.go
```

Example:

```bash
go run greeting.go
```

Open the URL displayed in the terminal or visit the appropriate endpoint in your browser.

---

## 🌐 Example Endpoints

| Service | Example URL |
|----------|-------------|
| Hello World | http://localhost:8080/ |
| Greeting | http://localhost:8080/greet?name=Grace |
| Square Number | http://localhost:8080/square?num=5 |
| Temperature Converter | http://localhost:8080/convert?celsius=30 |
| Calculator | http://localhost:8080/add?a=10&b=20 |
| Weather | http://localhost:8080/weather |
| Factorial | http://localhost:8080/factorial?num=5 |
| Even/Odd | http://localhost:8080/check?num=10 |
| Random Number | http://localhost:8080/random |
| Current Time | http://localhost:8080/time |

> **Note:** Since each program starts its own web server, run one service at a time on the same port, or modify the port number if running multiple services simultaneously.

---

## 📚 Learning Outcomes

- Understanding Go syntax and programming fundamentals
- Creating HTTP servers using the `net/http` package
- Handling URL routing and query parameters
- Building REST-style web services
- Writing modular backend applications
- Using Git for version control
- Managing projects on GitHub

---

## 👩‍💻 Author

**Grace Paul**

B.Tech Computer Science and Engineering

Christ College of Engineering, Irinjalakuda

---

## 📄 License

This project is created for educational and learning purposes.
