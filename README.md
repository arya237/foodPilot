# 🍽️ FoodPilot  

**FoodPilot** is an automation tool that helps students of **Bu-Ali Sina University** automatically reserve meals without the hassle of manual booking. Built with **Go 1.24**, it’s lightweight, fast, and easy to run both locally and in containers.  

---

## ✨ Features  

- 🔄 **Automatic reservation** of daily meals  
- ⏰ **Scheduled booking** at the right time  
- 🔑 **Secure login** with university credentials  
- 🐳 **Docker-ready** for quick deployment  
- ⚡ Built with modern **Go 1.24** concurrency  

---

## 🛠️ Tech Stack  

- **Language:** [Go 1.24](https://go.dev/)  
- **Containerization:** Docker  
- **CI/CD:** GitHub Actions (for build and image publishing)  

---

## 🚀 Getting Started  

### Prerequisites  
- [Go 1.24+](https://go.dev/dl/) installed  
- Or [Docker](https://docs.docker.com/get-docker/) installed  

---

### 🔧 Installation (from source)  

Clone the repository:  

```bash
git clone https://github.com/<your-username>/foodpilot.git
cd foodpilot
```

Build the project:

```
go build -o foodpilot ./...
```

Run:

```
./foodpilot
```
