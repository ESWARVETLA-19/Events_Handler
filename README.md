# 🚀 Go Project - REST API  

This is a **Go-based REST API** project. Follow the steps below to set up and run the application.  

---

## 📥 Clone the Repository  

Clone the repository and navigate into the project directory:  

```sh
git clone https://github.com/ESWARVETLA-19/Events_Handler.git
cd Events_Handler
```

---

## 🛠️ Build the Project  

Compile the Go application:  

```sh
go build -o app
```

---

## ▶️ Run the Project  

Start the application:  

```sh
go run main.go
```

---

## 🐳 Run Using Docker  

### **1️⃣ Build the Docker Image**  
To create a Docker image, run:  

```sh
docker build -t my-go-app .
```

- **`-t my-go-app`** → Assigns a tag (`my-go-app`) to the image. You can change it to any preferred name.  
- **`.` (dot)** → Specifies the **current directory** as the build context.  

---

### **2️⃣ Run the Docker Container**  
Start a container and expose port 8080:  

```sh
docker run -p 8080:8080 my-go-app
```

- **`-p 8080:8080`** → Maps port **8080** from the container to the **host machine**, allowing external access to the API.  

---
<!--
## 📝 License  

This project is licensed under the [MIT License](LICENSE).  

---

### 🎯 **Now your Go REST API is up and running!** 🚀  

---
-->


