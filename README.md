# Weather History API

![Go Version](https://img.shields.io/github/go-mod/go-version/Cre4T3Tiv3/weather-history-api)
![License](https://img.shields.io/github/license/Cre4T3Tiv3/weather-history-api)
![Project Status](https://img.shields.io/badge/status-active-brightgreen)

🚀 **Project Status**: This RESTful API project is actively under development, with frequent updates and enhancements.

The Weather History API, developed in Go, archives and presents historical weather data alongside significant events. This project serves as a portfolio piece to showcase Golang skills and expertise in RESTful API development.

---

## 📚 Table of Contents
- [Features](#features)
- [Quick Start](#quick-start)
- [Detailed Setup Guide](#detailed-setup-guide)
- [API Endpoints](#api-endpoints)
- [Advanced Documentation](#advanced-documentation)
- [Contributing](#contributing)
- [Future Enhancements](#future-enhancements)
- [License](#license)

---

## 🔥 Features

- CRUD operations for weather data on specific dates.
- Integration of historical events with weather data.
- Search and analytics capabilities.
- Ongoing development for additional features.

---

## 🚀 Quick Start

### Prerequisites
- **Golang**: Latest stable version.
- **PostgreSQL**: Database storage.

### Steps
1. **Clone the Repository:**
   ```bash
   git clone https://github.com/Cre4T3Tiv3/weather-history-api.git
   cd weather-history-api
   ```

2. **Install Dependencies:**
   ```bash
   go mod download
   ```

3. **Set Up PostgreSQL Database:**
   - Create a PostgreSQL database using `pgAdmin` or CLI.
   - Example database name: `weather_history`.

4. **Create `config.json`:**
   Inside the `configs` directory, create a `config.json` file with the following structure:
   ```json
   {
     "DBHost": "localhost",
     "DBPort": 5432,
     "DBUser": "your_username",
     "DBPassword": "your_password",
     "DBName": "weather_history"
   }
   ```

5. **Run the Server:**
   ```bash
   go run cmd/weather-api/main.go
   ```

For troubleshooting or advanced setup details, refer to the [Wiki Setup Guide](https://github.com/Cre4T3Tiv3/weather-history-api/wiki/Setup-Guide).

---

## 📡 API Endpoints

### Add Weather Data
```bash
curl -X POST http://localhost:8080/weather -d '{
  "date": "2023-10-19",
  "temperature": 20.5,
  "description": "Sunny"
}'
```

### Get Weather Data by Date
```bash
curl http://localhost:8080/weather?date=2023-10-19
```

For a complete list of API endpoints and advanced usage examples, visit the [API Documentation in the Wiki](https://github.com/Cre4T3Tiv3/weather-history-api/wiki/API-Documentation).

---

## 📖 Detailed Setup Guide
For an in-depth setup tutorial, including PostgreSQL configuration and troubleshooting, visit the [Setup Guide in the Wiki](https://github.com/Cre4T3Tiv3/weather-history-api/wiki/Setup-Guide).

---

## 📘 Advanced Documentation

Explore detailed documentation for:
- [Error Handling](https://github.com/Cre4T3Tiv3/weather-history-api/wiki/Error-Handling)
- [Testing](https://github.com/Cre4T3Tiv3/weather-history-api/wiki/Testing)
- [CI/CD Integration](https://github.com/Cre4T3Tiv3/weather-history-api/wiki/CI-CD-Integration)
- [FAQs](https://github.com/Cre4T3Tiv3/weather-history-api/wiki/FAQs)

---

## 🤝 Contributing
Contributions are welcome! Here’s how to get started:
1. Fork this repository.
2. Create a new branch:
   ```bash
   git checkout -b feature/add-weather-endpoint
   ```
3. Commit your changes using the following convention:
   - For adding features:
     ```bash
     git commit -m "feat(api): add weather data endpoint"
     ```
   - For fixing bugs:
     ```bash
     git commit -m "fix(db): resolve connection timeout"
     ```
   - For updating documentation:
     ```bash
     git commit -m "docs(readme): update setup instructions"
     ```
4. Push to the branch:
   ```bash
   git push origin feature/add-weather-endpoint
   ```
5. Open a Pull Request.

For detailed contribution guidelines, visit the [Wiki Contributing Section](https://github.com/Cre4T3Tiv3/weather-history-api/wiki/Contributing).

---

## 💡 Future Enhancements
- Unit and integration tests.
- CI/CD pipelines with Jenkins or GitHub Actions.
- Additional endpoints for analytics (e.g., temperature trends).

---

## 📜 License
This project is licensed under the [MIT License](LICENSE).

---