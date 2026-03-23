---
title: Blue-Green & Canary Deployment on Kubernetes
tags: [Kubernetes, Deployment, Canary, DevOps]
---

# 🚀 Blue-Green & Canary Deployment on Kubernetes

<div className="lab-actions">

<a href="https://github.com/bakuppus/Blue-Green-Canary-Demo-Application" target="_blank" className="lab-btn">
💻 View Source Code
</a>

<a href="https://github.com/bakuppus/Blue-Green-Canary-Demo-Application" target="_blank" className="lab-btn primary">
🚀 Open Playground
</a>

</div>

---

## ⚡ Lab Quick Insight

Implement **Blue-Green and Canary deployment strategies** in Kubernetes to achieve:

- Zero downtime deployments  
- Controlled traffic shifting  
- Safe production releases  

---

## 🏗️ Architecture Diagram

![architecture](./images/architecture.png)

---

## 🧰 Tools Used

- Kubernetes (EKS / Minikube)
- kubectl
- Docker
- NGINX / Service routing
- GitHub

---

## 🚀 Quick Start: Deploy in 15 Minutes

<div className="lab-timeline">

  {/* STEP 1 */}
  <div className="lab-step">
    <div className="lab-step-number">1</div>

    <div className="lab-step-card">
      <h3>Prerequisites</h3>

      <p><strong>Required Tools:</strong></p>

```bash
# Verify installations
aws --version
kubectl version
docker --version
  <p><strong>AWS Setup:</strong></p>

  <ul>
    <li>IAM permissions: <code>AdministratorAccess</code></li>
    <li>AWS configured: <code>aws configure</code></li>
    <li>Region: <code>us-east-1</code></li>
  </ul>
</div>
</div>

{/* STEP 2 */}

<div className="lab-step"> <div className="lab-step-number">2</div>
<div className="lab-step-card">
  <h3>Deploy Blue Version</h3>
kubectl apply -f blue-deployment.yaml
kubectl apply -f service.yaml
  <p>Blue version is now live.</p>
</div>
</div>

{/* STEP 3 */}

<div className="lab-step"> <div className="lab-step-number">3</div>
<div className="lab-step-card">
  <h3>Deploy Green Version</h3>
kubectl apply -f green-deployment.yaml
  <p>Green version deployed alongside Blue.</p>
</div>
</div>

{/* STEP 4 */}

<div className="lab-step"> <div className="lab-step-number">4</div>
<div className="lab-step-card">
  <h3>Switch Traffic (Blue → Green)</h3>
kubectl patch service my-app \
-p '{"spec":{"selector":{"version":"green"}}}'
  <p>Traffic is now routed to Green deployment.</p>
</div>
</div>

{/* STEP 5 */}

<div className="lab-step"> <div className="lab-step-number">5</div>
<div className="lab-step-card">
  <h3>Canary Deployment</h3>
apiVersion: apps/v1
kind: Deployment
metadata:
  name: canary
spec:
  replicas: 1
  <p>Gradually shift traffic and monitor behavior.</p>
</div>
</div> </div>