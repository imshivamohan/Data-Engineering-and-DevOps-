# **Kubernetes Basic Course** 🚀

## **1. Introduction to Kubernetes**
Kubernetes (K8s) is a powerful container orchestration platform used to manage, scale, and automate the deployment of containerized applications.

### **Why Kubernetes?**
- Automates deployment, scaling, and management of applications.
- Ensures high availability and fault tolerance.
- Provides load balancing and self-healing features.
- Allows for declarative configuration using YAML.

### **Check Kubernetes Version**
```sh
kubectl version --short
```

### **View Cluster Info**
```sh
kubectl cluster-info
```

---

## **2. Understanding Kubernetes YAML Files**

A **Kubernetes manifest** (YAML file) defines how an application should run.

### **Key Fields in YAML**
- `apiVersion`: The Kubernetes API version (e.g., `v1`, `apps/v1`).
- `kind`: The type of resource (e.g., `Pod`, `Deployment`, `Service`).
- `metadata`: Defines the name and labels of the resource.
- `spec`: Specifies the desired state of the object.

Example YAML breakdown:
```yaml
apiVersion: v1     # Kubernetes API version
kind: Pod          # Type of resource
metadata:
  name: my-pod     # Name of the Pod
spec:
  containers:
    - name: nginx-container
      image: nginx:latest
```

---

## **3. Namespaces in Kubernetes**

Namespaces help in organizing Kubernetes resources within a cluster.

```sh
kubectl get namespaces
```

Create a new namespace:
```sh
kubectl create namespace dev-namespace
```

---

## **4. Pods: The Smallest Deployable Unit**

A **Pod** is a group of containers that share storage and networking.

### **Create a Simple Pod (YAML)**
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx-pod
spec:
  containers:
    - name: nginx
      image: nginx:latest
      ports:
        - containerPort: 80
```

Apply the Pod:
```sh
kubectl apply -f nginx-pod.yaml
```

View running Pods:
```sh
kubectl get pods
```

Delete a Pod:
```sh
kubectl delete pod nginx-pod
```

---

## **5. Deployments: Managing Application Scaling**

A **Deployment** ensures that a set number of Pods are always running.

### **Create a Deployment (YAML)**
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 2
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
        - name: nginx
          image: nginx:latest
          ports:
            - containerPort: 80
```

Apply the Deployment:
```sh
kubectl apply -f nginx-deployment.yaml
```

Scale the Deployment:
```sh
kubectl scale deployment nginx-deployment --replicas=5
```

---

## **6. Services: Exposing Applications**

A **Service** exposes a Kubernetes application to the network.

### **Types of Services**
- **ClusterIP**: Exposes the service internally.
- **NodePort**: Exposes the service on a static port.
- **LoadBalancer**: Provides external access via a cloud provider.

### **Create a Service (YAML)**
```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx-service
spec:
  selector:
    app: nginx
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
  type: LoadBalancer
```

Apply the Service:
```sh
kubectl apply -f nginx-service.yaml
```

Check running services:
```sh
kubectl get services
```

---

## **7. ConfigMaps & Secrets**

### **What is a ConfigMap?**
A **ConfigMap** allows you to store configuration data separately from application code.

Create a ConfigMap:
```sh
kubectl create configmap app-config --from-literal=ENV=production
```

### **What is a Secret?**
A **Secret** is used to store sensitive data like passwords and API keys.

Create a Secret:
```sh
kubectl create secret generic db-secret --from-literal=PASSWORD=mysecretpass
```

---

## **8. Persistent Storage (Volumes & PVCs)**

### **Why Use Persistent Volumes?**
- Allows data to persist beyond Pod lifecycles.
- Helps in database applications where data must be stored.

### **Create a Persistent Volume (YAML)**
```yaml
apiVersion: v1
kind: PersistentVolume
metadata:
  name: my-pv
spec:
  capacity:
    storage: 1Gi
  accessModes:
    - ReadWriteOnce
  hostPath:
    path: "/mnt/data"
```

### **Create a PersistentVolumeClaim (YAML)**
```yaml
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: my-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
```

---

## **9. Helm: The Kubernetes Package Manager**

### **What is Helm?**
Helm is a package manager for Kubernetes that helps deploy applications easily using predefined charts.

### **Install Helm**
```sh
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
```

### **Add a Helm Repo**
```sh
helm repo add stable https://charts.helm.sh/stable
```

### **Install Nginx Using Helm**
```sh
helm install mynginx stable/nginx-ingress
```

---

## **10. Scaling & Auto-Scaling**

### **Scale a Deployment**
```sh
kubectl scale deployment nginx-deployment --replicas=5
```

### **Enable Auto-Scaling**
```sh
kubectl autoscale deployment nginx-deployment --cpu-percent=50 --min=2 --max=10
```

---

## **Conclusion**
This updated course now includes **detailed explanations** and practical examples to help you understand Kubernetes basics. 🚀
