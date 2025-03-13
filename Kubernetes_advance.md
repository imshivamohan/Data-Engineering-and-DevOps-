# **Kubernetes Advanced Course** 🚀

## **1. Role-Based Access Control (RBAC)**

### **What is RBAC?**
RBAC (Role-Based Access Control) is used in Kubernetes to manage user permissions.

### **Why Use RBAC?**
- Provides **fine-grained** access control.
- Ensures **security** by restricting access to cluster resources.
- Prevents unauthorized users from making changes.

### **Create a Role (YAML)**
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: default
  name: pod-reader
rules:
  - apiGroups: [""]
    resources: ["pods"]
    verbs: ["get", "list", "watch"]
```

### **Bind Role to User**
```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: read-pods-binding
  namespace: default
subjects:
  - kind: User
    name: example-user
    apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role
  name: pod-reader
  apiGroup: rbac.authorization.k8s.io
```

Apply RBAC settings:
```sh
kubectl apply -f role.yaml
kubectl apply -f rolebinding.yaml
```

---

## **2. Helm (Kubernetes Package Manager)**

### **What is Helm?**
Helm is a package manager for Kubernetes that allows deploying applications using reusable templates called **charts**.

### **Why Use Helm?**
- Simplifies **application deployment**.
- Manages **versioning** of Kubernetes resources.
- Supports **rolling updates**.

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

## **3. Monitoring with Prometheus & Grafana**

### **What is Prometheus?**
Prometheus is an open-source monitoring system that collects metrics from Kubernetes workloads.

### **What is Grafana?**
Grafana is a visualization tool that helps display monitoring data from Prometheus.

### **Deploy Prometheus Using Helm**
```sh
helm install prometheus prometheus-community/kube-prometheus-stack
```

### **Expose Prometheus UI**
```sh
kubectl port-forward svc/prometheus-kube-prometheus-prometheus 9090
```

### **Deploy Grafana**
```sh
kubectl apply -f grafana.yaml
```

---

## **4. Ingress Controller for Load Balancing**

### **What is an Ingress Controller?**
An Ingress Controller manages external access to Kubernetes services via **HTTP(S) routing**.

### **Why Use an Ingress Controller?**
- Routes **multiple services** under a single domain.
- Provides **TLS termination** for secure HTTPS traffic.

### **Install Ingress-Nginx**
```sh
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/cloud/deploy.yaml
```

### **Create an Ingress Rule (YAML)**
```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: example-ingress
spec:
  rules:
    - host: example.local
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: nginx-service
                port:
                  number: 80
```

### **Apply the Ingress**
```sh
kubectl apply -f ingress.yaml
```

---

## **5. Network Policies for Security**

### **What are Network Policies?**
Network Policies define how **Pods** communicate with each other in a Kubernetes cluster.

### **Why Use Network Policies?**
- Restricts traffic **between Pods**.
- Prevents **unauthorized** access.

### **Create a Network Policy (YAML)**
```yaml
apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: deny-all
  namespace: default
spec:
  podSelector: {}
  policyTypes:
    - Ingress
```

### **Apply the Network Policy**
```sh
kubectl apply -f network-policy.yaml
```

---

## **6. StatefulSets for Stateful Applications**

### **What is a StatefulSet?**
A **StatefulSet** is used for deploying **stateful applications** like databases (e.g., MySQL, Cassandra).

### **Why Use StatefulSets?**
- Ensures **persistent storage**.
- Maintains **Pod identity** across restarts.

### **Create a StatefulSet (YAML)**
```yaml
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: mysql-db
spec:
  selector:
    matchLabels:
      app: mysql
  serviceName: "mysql"
  replicas: 2
  template:
    metadata:
      labels:
        app: mysql
    spec:
      containers:
        - name: mysql
          image: mysql:latest
          env:
            - name: MYSQL_ROOT_PASSWORD
              value: "password"
```

### **Apply the StatefulSet**
```sh
kubectl apply -f statefulset.yaml
```

---

## **7. Kubernetes Security Best Practices**

### **1. Use Role-Based Access Control (RBAC)**
```sh
kubectl create role developer --verb=get,list --resource=pods
```

### **2. Enable Network Policies**
```sh
kubectl apply -f network-policy.yaml
```

### **3. Run Containers as Non-Root**
```yaml
securityContext:
  runAsNonRoot: true
```

### **4. Scan Container Images for Vulnerabilities**
```sh
trivy image nginx:latest
```

---

## **Conclusion**
This advanced Kubernetes course includes **security, networking, Helm, monitoring, and RBAC** with explanations and practical examples. 🚀
