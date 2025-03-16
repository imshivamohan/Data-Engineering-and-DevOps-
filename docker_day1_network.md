# Docker Networking and Its Types

Docker provides various networking options to connect containers. Each type serves different use cases based on container communication needs.

## 1️⃣ Bridge Network (Default)
The **bridge network** is the default network mode in Docker. Containers on the same bridge network can communicate with each other using container names, but they are isolated from external networks.

### 🔹 Example: Create a Custom Bridge Network
```bash
# Create a custom bridge network
docker network create my_bridge_network

# Run two containers in the custom bridge network
docker run -d --network my_bridge_network --name web nginx
docker run -d --network my_bridge_network --name app alpine sleep 3600
```
Now, the `app` container can reach `web` using its container name:
```bash
docker exec -it app ping web
```

---

## 2️⃣ Host Network
The **host network** makes the container share the same network as the host system. This removes network isolation.

### 🔹 Example: Run a Container on the Host Network
```bash
docker run -d --network host --name my_nginx nginx
```
Now, Nginx runs on the **same network as the host**, so you can access it directly via `http://localhost:80`.

🛑 **Limitation:** You **cannot** run multiple containers on the same port since they share the host’s network.

---

## 3️⃣ None Network
The **none network** disables networking for the container. The container has no external or internal communication.

### 🔹 Example: Run a Container with No Networking
```bash
docker run -d --network none --name isolated_container alpine sleep 3600
```
The container is completely isolated from other containers and the internet.

---

## 4️⃣ Macvlan Network
The **Macvlan network** assigns containers real IP addresses from the same subnet as the host. Containers act like physical devices on the network.

### 🔹 Example: Create a Macvlan Network
```bash
# Create a macvlan network
docker network create -d macvlan \
  --subnet=192.168.1.0/24 \
  --gateway=192.168.1.1 \
  -o parent=eth0 \
  my_macvlan

# Run a container with a specific IP
docker run -d --network my_macvlan --ip 192.168.1.100 --name macvlan_container nginx
```
Now, `192.168.1.100` is accessible from the local network.

🛑 **Limitation:** The host **cannot communicate** with Macvlan containers unless you create a special bridge.

---

## 5️⃣ Ipvlan Network
The **Ipvlan network** is similar to Macvlan but allows containers to share the same MAC address as the host. This reduces network complexity.

### 🔹 Example: Create an Ipvlan Network
```bash
# Create an ipvlan network
docker network create -d ipvlan \
  --subnet=192.168.1.0/24 \
  --gateway=192.168.1.1 \
  -o parent=eth0 \
  my_ipvlan

# Run a container with a specific IP
docker run -d --network my_ipvlan --ip 192.168.1.101 --name ipvlan_container nginx
```
Now, `192.168.1.101` can communicate with the network just like a normal device.

🛑 **Limitation:** Ipvlan is best for environments where **MAC address conflicts** are a concern.

---

## 6️⃣ Overlay Network (For Docker Swarm)
The **overlay network** allows containers on different Docker hosts to communicate as if they are on the same network. It is primarily used in **Docker Swarm**.

### 🔹 Example: Create an Overlay Network in Swarm Mode
```bash
# Initialize Docker Swarm
docker swarm init

# Create an overlay network
docker network create -d overlay my_overlay_network

# Deploy a service to the overlay network
docker service create --name web --network my_overlay_network nginx
```
Now, all containers in `my_overlay_network` can communicate across different hosts in the Swarm.

---

## 🎯 Summary Table

| Network Type | Communication Scope | Use Case |
|-------------|----------------|---------|
| **Bridge** | Containers on the same bridge | Default for container-to-container communication |
| **Host** | Shares host’s network | Performance-critical apps (e.g., DNS, web servers) |
| **None** | No communication | Completely isolated containers |
| **Macvlan** | Containers get real network IPs | Requires full integration with physical network |
| **Ipvlan** | Shares host’s MAC address | Ideal when MAC conflicts are a concern |
| **Overlay** | Multi-host communication | Docker Swarm for scaling apps |

🚀 **Now you can choose the best network type for your Docker setup!**

