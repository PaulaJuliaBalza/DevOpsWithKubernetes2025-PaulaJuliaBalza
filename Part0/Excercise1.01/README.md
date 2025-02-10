# Excercise 1.01: Getting started

Exercises can be done with any language and framework you want.

Create an application that generates a random string on startup, stores this string into memory, and outputs it every 5 seconds with a timestamp. e.g.

```bash2020-03-30T12:15:17.705Z: 8523ecb1-c716-4cb6-a044-b9e83bb98e43
2020-03-30T12:15:22.705Z: 8523ecb1-c716-4cb6-a044-b9e83bb98e43
```
Deploy it into your Kubernetes cluster and confirm that it's running with kubectl logs ...

You will keep building this application in the future exercises. This application will be called "Log output".

# Solution
```bash
#Build and push the image
docker build -t paulajuliabalza/log-output:latest .
docker login
docker push paulajuliabalza/log-output:latest

#Deploy to k3d Kubernetes cluster 
```bash
k3d --help 
k3d cluster create -a 2 #Create cluster
k3d cluster stop/start  #Start/Stop cluster
kubectl config use-context k3d-k3s-default #Set Context
kubectl apply -f deployment.yaml #Create Deployment

administrator@ASC:~/Escritorio/DevOpsWithKubernetes2025-PaulaJuliaBalza/Part0/Exercise1.01$ kubectl get pods
NAME                         READY   STATUS    RESTARTS   AGE
log-output-fdc46c986-stlhf   1/1     Running   0          2m44s

administrator@ASC:~/Escritorio/DevOpsWithKubernetes2025-PaulaJuliaBalza/Part0/Exercise1.01$ kubectl logs -f log-output-fdc46c986-stlhf
2025-02-10T20:42:23.237675669Z: 546f25b8-3ab4-4406-b082-95209dbda333
2025-02-10T20:42:28.237059428Z: 546f25b8-3ab4-4406-b082-95209dbda333
2025-02-10T20:42:33.239227169Z: 546f25b8-3ab4-4406-b082-95209dbda333
2025-02-10T20:42:38.238387998Z: 546f25b8-3ab4-4406-b082-95209dbda333
2025-02-10T20:42:43.237530551Z: 546f25b8-3ab4-4406-b082-95209dbda333
2025-02-10T20:42:48.237496212Z: 546f25b8-3ab4-4406-b082-95209dbda333
```
