## Go Web Application
A practical DevOps project to containerize, deploy, and operationalize a simple Golang web app using Docker, Kubernetes (EKS), Helm, GitHub Actions, and ArgoCD.

### Project Overview
This project showcases the DevOps implementation of a basic Golang web server built using the net/http package. The app is served on /courses endpoint.

The goal is to take an app without any existing DevOps practices and implement:

* Docker-based containerization
* Kubernetes deployment (EKS)
* Ingress controller and custom DNS mapping
* Helm for multi-environment support
* CI with GitHub Actions
* CD with GitOps using ArgoCD

### Project Workflow

**STAGE 1: Containerize the Application using Multi-Stage Dockerfile**

This stage focuses on understanding the Go application and creating an optimized Docker image for it.

**Step 1:** Understand and Run the Application Locally

```
git clone https://github.com/Dzennieva/go-app.git
cd go-app
```

Run the application locally:

```
go build -o main .
./main # Access via http://localhost:8080/courses
```

**Step 2:** Build the Docker image and push to Docker Hub:
 ```
docker build -t <your-username>/go-app:latest .
docker push <your-username>/go-app:latest
```

**STAGE 2: Deploy to Kubernetes Cluster using EKS and Ingress Controller (Nginx)**

This stage covers the deployment of the containerized application to an EKS cluster and exposing it via an Ingress Controller.

**Step 1:** Kubernetes Deployment and Service
Create Kubernetes Manifests

**Install Prerequisites**

``eksctl``: For creating and managing EKS clusters. <br>
``AWS CLI``: For interacting with AWS services.<br>
``kubectl``: For interacting with Kubernetes clusters.

Authenticate with AWS
```
aws configure
```
Install EKS Cluster

```
eksctl create cluster --name <cluster-name> --region <aws-region>
```
Apply Manifest Files

```
kubectl apply -f k8s/deployment.yml
```

Expose Service via NodePort and Test

``Get node IP``
```
kubectl get nodes -o wide
```
Access the application via ``http://<nodeip>:<nodeport>/courses``

_Update the EC2 Security Group to allow traffic on NodePort range (30000-32767)._


**Step 2:** Install Ingress-Nginx Controller and Map DNS<br>
Install Ingress-NGINX

```
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.11.1/deploy/static/provider/aws/deploy.yaml
kubectl get pods -n ingress-nginx
```
Create Ingress resource and get the external IP:

```
kubectl get ing
nslookup <ingress-host>
```
Map IP to custom domain (for testing):
```
sudo nano /etc/hosts
```
``<LB-IP> go-app.local``

_Some browsers ignore custom host mappings. Use the DNS name provided by the Ingress controller instead._

**STAGE 3: Use Helm Chart**

This stage demonstrates how to package your Kubernetes manifests into a Helm chart for easier management and deployment across environments.

**Step 1:** Create and Customize Helm Chart
Create a helm directory

```
mkdir helm
cd helm
```
Create a new Helm chart:
```
helm create go-web-app
```
Delete default charts and modify parameters in values.yaml

Install Helm chart

```
helm install my-go-app ./go-web-app
```
Uninstall Helm chart (when done):
```
helm uninstall my-go-app
```

**STAGE 4: Implement CI/CD**

This stage focuses on setting up Continuous Integration (CI) with GitHub Actions and Continuous Delivery (CD) with ArgoCD.

Step 1: CI with GitHub Actions
- Linting with golangci-lint

- Docker build and push
- Update image tag in Values.yml

_Use GitHub Secrets for sensitive values_

**Step 2:** CD with GitOps/ArgoCD

Install ArgoCD on the Cluster
```
kubectl create namespace argocd

kubectl apply -n argocd -f https://raw.githubusercontent.com/argoproj/argo-cd/stable/manifests/install.yaml
```
Access ArgoCD UI
```
kubectl patch svc argocd-server -n argocd -p '{"spec": {"type": "LoadBalancer"}}'
```
Login to ArgoCD

Get initial admin password
```
kubectl get secret argocd-initial-admin-secret -n argocd -o jsonpath="{.data.password}" | base64 -d && echo
```
Default username: ``admin``



