---
title: Blue-Green & Canary Deployment on Kubernetes
tags: [Kubernetes, Deployment, Canary, DevOps]
---

<style>{`
.lab-container {
  margin-top: 30px;
}

.lab-actions a {
  margin-right: 10px;
  padding: 10px 16px;
  border-radius: 8px;
  text-decoration: none;
  border: 1px solid #e5e7eb;
  display: inline-block;
}

.lab-actions .primary {
  background: #22c55e;
  color: white;
}

.lab-timeline {
  position: relative;
  padding-left: 40px;
  margin-top: 30px;
}

.lab-timeline::before {
  content: '';
  position: absolute;
  left: 18px;
  top: 0;
  bottom: 0;
  width: 3px;
  background: linear-gradient(to bottom, #6366f1, #22c55e);
}

.lab-step {
  position: relative;
  margin-bottom: 40px;
}

.lab-step-number {
  position: absolute;
  left: -4px;
  width: 36px;
  height: 36px;
  border-radius: 50%;
  background: linear-gradient(135deg, #6366f1, #8b5cf6);
  color: white;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
}

.lab-card {
  margin-left: 40px;
  padding: 20px;
  border-radius: 16px;
  border: 1px solid #e5e7eb;
  background: white;
  box-shadow: 0 10px 25px rgba(0,0,0,0.05);
}

.lab-card pre {
  background: #0f172a;
  color: #e5e7eb;
  padding: 12px;
  border-radius: 10px;
}

.lab-cta {
  margin-top: 40px;
  padding: 20px;
  border-radius: 12px;
  background: #f9fafb;
  border: 1px solid #e5e7eb;
}
`}</style>

# 🚀 Blue-Green & Canary Deployment on Kubernetes

<div className="lab-actions">
  <a href="https://github.com/bakuppus/Blue-Green-Canary-Demo-Application">💻 View Code</a>
  <a href="https://github.com/bakuppus/Blue-Green-Canary-Demo-Application" className="primary">🚀 Playground</a>
</div>

---

## ⚡ Lab Quick Insight

Deploy applications using **Blue-Green and Canary strategies** to achieve:

- Zero downtime  
- Safe releases  
- Gradual traffic control  

---

## 🏗️ Architecture

![architecture](./images/architecture.png)

---

## 🧰 Tools Used

- Kubernetes  
- kubectl  
- Docker  
- AWS / EKS  

---

## 🚀 Quick Start: Deploy in 15 Minutes

<div className="lab-timeline">

<div className="lab-step">
<div className="lab-step-number">1</div>

<div className="lab-card">

### Prerequisites

```bash
aws --version
kubectl version
docker --version