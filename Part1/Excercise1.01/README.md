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

k3d --help 
k3d cluster create -a 2 #Create cluster
k3d cluster stop/start  #Start/Stop cluster
kubectl config use-context k3d-k3s-default #Set Context

administrator@ASC:~/Escritorio/DevOpsWithKubernetes2025-PaulaJuliaBalza/Part1/Excercise1.01$ kubectl create deployment log-output --image=paulajuliabalza/log-output:latest
deployment.apps/log-output created


administrator@ASC:~/Escritorio/DevOpsWithKubernetes2025-PaulaJuliaBalza/Part1/Excercise1.01$ kubectl get pods
NAME                         READY   STATUS    RESTARTS   AGE
log-output-fdc46c986-z57bg   1/1     Running   0          55s

administrator@ASC:~/Escritorio/DevOpsWithKubernetes2025-PaulaJuliaBalza/Part1/Excercise1.01$ kubectl logs -f log-output-fdc46c986-z57bg
2025-04-07T23:30:11.041Z: 5af7594d-80a2-4888-84f8-98dace9c1895
2025-04-07T23:30:16.046Z: 5af7594d-80a2-4888-84f8-98dace9c1895
2025-04-07T23:30:21.048Z: 5af7594d-80a2-4888-84f8-98dace9c1895
2025-04-07T23:30:26.052Z: 5af7594d-80a2-4888-84f8-98dace9c1895
2025-04-07T23:30:31.058Z: 5af7594d-80a2-4888-84f8-98dace9c1895
2025-04-07T23:30:36.061Z: 5af7594d-80a2-4888-84f8-98dace9c1895
2025-04-07T23:30:41.064Z: 5af7594d-80a2-4888-84f8-98dace9c1895
2025-04-07T23:30:46.065Z: 5af7594d-80a2-4888-84f8-98dace9c1895
^Cadministrator@ASC:~/Escritorio/DevOpsWithKubernetes2025-PaulaJuliaBalza/Part1/Excercise1.01$ 
```
