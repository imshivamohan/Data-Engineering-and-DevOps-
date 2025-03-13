# **Docker Complete Course with Practical Examples**

## **1. Introduction to Docker**
Docker is a platform for developing, shipping, and running applications inside lightweight containers.

### **Installing Docker**
- Follow the official guide: [Docker Installation Guide](https://docs.docker.com/get-docker/)

### **Check Docker Version**
```sh
docker --version
```

---

## **2. Basic Docker Commands**

### **Run a Container**
```sh
docker run alpine echo "Hello, Docker!"
```

### **Run an Interactive Shell in a Container**
```sh
docker run -it alpine sh
```

### **List Running Containers**
```sh
docker ps
```

### **List All Containers (Including Stopped)**
```sh
docker ps -a
```

### **Stop a Running Container**
```sh
docker stop <container_id>
```

### **Start a Stopped Container**
```sh
docker start <container_id>
```

### **Remove a Container**
```sh
docker rm <container_id>
```

### **Remove all Container**
```sh
docker container prune
```

### **List Available Images**
```sh
docker images
```

### **Remove an Image**
```sh
docker rmi <image_id>
```

---

## **3. Working with Ports**

### **Run Nginx and Expose Port 8080**
```sh
docker run -d -p 8080:80 --name mynginx nginx
```

Access at:  
👉 `http://localhost:8080`

### **Check Container Ports**
```sh
docker port mynginx
```

---

## **4. Working with Volumes**

### **List Available Volumes**
```sh
docker volume ls
```

### **Create a Named Volume**
```sh
docker volume create mydata
```

### **Mount Volume in a Container (`--mount`)**
```sh
docker run -d --name alpine-volume --mount source=mydata,target=/data alpine sleep 600
```

### **Mount a Host Directory as a Volume (`-v`)**
```sh
docker run -d -v $(pwd)/html:/usr/share/nginx/html nginx
```

### **Check Data Inside a Volume**
```sh
docker run --rm -it --mount source=mydata,target=/data alpine sh
ls -l /data
```

### **Inspect a Volume**
```sh
docker volume inspect mydata
```

### **Remove a Volume**
```sh
docker volume rm mydata
```

---

## **5. Docker Networking**

### **List Docker Networks**
```sh
docker network ls
```

### **Create a Custom Network**
```sh
docker network create mynetwork
```

### **Run Containers on a Custom Network**
```sh
docker run -d --name alpine1 --network mynetwork alpine sleep 600
docker run -d --name alpine2 --network mynetwork alpine sleep 600
```

### **Test Connectivity (Inside a Container)**
```sh
docker exec -it alpine1 sh
ping alpine2
```

### **Connect a Running Container to a Network**
```sh
docker network connect mynetwork alpine1
```

---

## **6. Building Custom Images**

### **Create a Simple `Dockerfile`**
```dockerfile
FROM alpine
RUN echo "Hello from a custom Docker image!" > /message.txt
CMD cat /message.txt
```

### **Build the Image**
```sh
docker build -t mycustomimage .
```

### **Run the Custom Image**
```sh
docker run mycustomimage
```

---

## **7. Tagging and Pushing Images to Docker Hub**

### **Login to Docker Hub**
```sh
docker login
```

### **Tag an Image**
```sh
docker tag mycustomimage mydockerhubusername/mycustomimage:v1
```

### **Push the Image to Docker Hub**
```sh
docker push mydockerhubusername/mycustomimage:v1
```

### **Pull an Image from Docker Hub**
```sh
docker pull mydockerhubusername/mycustomimage:v1
```

---

## **8. Running Commands Inside Containers**

### **Execute a Command Inside a Running Container**
```sh
docker exec -it mynginx sh
```

### **Copy Files Between Host and Container**
```sh
docker cp index.html mynginx:/usr/share/nginx/html/index.html
```

---

## **9. Docker Compose Basics**

### **Basic `docker-compose.yml` File**
```yaml
version: '3.8'
services:
  alpine:
    image: alpine
    container_name: myalpine
    command: sleep 600
```

Run it:
```sh
docker compose up -d
```

### **Check Running Containers**
```sh
docker ps
```

### **Stop and Remove All Containers**
```sh
docker compose down
```

---

## **10. Multi-Container Setup (Nginx + Alpine Client)**

### **`docker-compose.yml` File**
```yaml
version: '3.8'
services:
  web:
    image: nginx
    container_name: mynginx
    ports:
      - "8080:80"
    networks:
      - mynetwork

  client:
    image: alpine
    container_name: myalpine
    command: sleep 600
    networks:
      - mynetwork

networks:
  mynetwork:
```

### **Start Both Containers**
```sh
docker compose up -d
```

### **Test Connection from Alpine**
```sh
docker exec -it myalpine sh
wget -qO- mynginx
```

---

## **11. Building a Custom Image with Compose**

### **Create `Dockerfile`**
```dockerfile
FROM nginx
COPY index.html /usr/share/nginx/html/
```

### **Create `docker-compose.yml`**
```yaml
version: '3.8'
services:
  web:
    build: .
    container_name: customnginx
    ports:
      - "8080:80"
```

### **Build and Run the Image**
```sh
docker compose up --build -d
```

Access at:  
👉 `http://localhost:8080`

---

## **12. Using Environment Variables in Compose**

### **Create `.env` File**
```sh
MESSAGE="Hello from .env file!"
```

### **Update `docker-compose.yml`**
```yaml
version: '3.8'
services:
  alpine:
    image: alpine
    container_name: myalpine
    environment:
      - MESSAGE=${MESSAGE}
    command: sh -c "echo $MESSAGE && sleep 600"
```

### **Run Compose**
```sh
docker compose up
```

You should see:  
```
Hello from .env file!
```

---

## **Conclusion**
This course covers all key Docker concepts, from running containers to networking and Compose. 🚀

Would you like to add more advanced topics? Let me know! 😃
