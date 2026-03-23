import LabStep from '@site/src/components/LabStep';

# 🚀 Blue-Green & Canary Deployment

---

## 🚀 Implementation

<LabStep number="1" title="Prerequisites">

```bash
aws --version
kubectl version
docker --version
</LabStep>
<LabStep number="2" title="Deploy Blue">
kubectl apply -f blue-deployment.yaml
kubectl apply -f service.yaml
</LabStep>
<LabStep number="3" title="Deploy Green">
kubectl apply -f green-deployment.yaml
</LabStep>
<LabStep number="4" title="Switch Traffic">
kubectl patch service my-app \
-p '{"spec":{"selector":{"version":"green"}}}'
</LabStep>
<LabStep number="5" title="Canary Deployment">
apiVersion: apps/v1
kind: Deployment
metadata:
  name: canary
spec:
  replicas: 1
</LabStep> ```