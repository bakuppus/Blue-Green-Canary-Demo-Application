---
title: Blue-Green & Canary Deployment
---

import LabStep from '@site/src/components/LabStep';

# 🚀 Blue-Green & Canary Deployment

---

## 🚀 Implementation

<LabStep number="1" title="Prerequisites">

<div>

```bash
aws --version
kubectl version
docker --version
```
</div> 
</LabStep>

<LabStep number="2" title="Deploy Blue"> 
<div>

```bash
kubectl apply -f blue-deployment.yaml
kubectl apply -f service.yaml
```
</div> </LabStep>

<LabStep number="3" title="Deploy Green">
 <div>

 ```bash
kubectl apply -f green-deployment.yaml
```
</div> </LabStep>

<LabStep number="4" title="Switch Traffic"> 
<div>

```bash
kubectl patch service my-app \
-p '{"spec":{"selector":{"version":"green"}}}'
```
</div> </LabStep>

<LabStep number="5" title="Canary Deployment"> 
<div>

```bash
apiVersion: apps/v1
kind: Deployment
metadata:
  name: canary
spec:
  replicas: 1
  ```
</div> </LabStep>