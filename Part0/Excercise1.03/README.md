# Excercise 1.03: Declarative approach

In your "Log output" application create a folder for manifests and move your deployment into a declarative file.

Make sure everything still works by restarting and following logs.

# Solution
```bash
   #Build and add a new tag
   docker build -t paulajuliabalza/log-output:1.0.0 .
   docker push paulajuliabalza/log-output:1.0.0
   #Specify the new tag in the deployment
   containers:
        - name: log-output
          image: paulajuliabalza/log-output:1.0.0
          imagePullPolicy: Always
  #Apply the new manifest
  kubectl apply -f manifests/deployment.yaml
  deployment.apps/log-output configured
  #Restart the pod
  kubectl delete pod -l app=log-output
  pod "log-output-67cb547fd-pxvwm" deleted
   kubectl get pods
   NAME                         READY   STATUS    RESTARTS   AGE
   log-output-67cb547fd-s4g5h   1/1     Running   0          6s
```

The output of kubectl describe shows how the tag updated:

```bash
administrator@ASC:~/Escritorio/DevOpsWithKubernetes2025-PaulaJuliaBalza/Part0/Excercise1.03$ kubectl describe pod log-output-67cb547fd-s4g5h
Name:             log-output-67cb547fd-s4g5h
Namespace:        default
Priority:         0
Service Account:  default
Node:             k3d-k3s-default-server-0/172.18.0.2
Start Time:       Mon, 10 Feb 2025 19:24:47 -0300
Labels:           app=log-output
                  pod-template-hash=67cb547fd
Annotations:      <none>
Status:           Running
IP:               10.42.1.7
IPs:
  IP:           10.42.1.7
Controlled By:  ReplicaSet/log-output-67cb547fd
Containers:
  log-output:
    Container ID:   containerd://0aabd38e501e609e74dfc7d1982f7001e403c40175d86634589432d9c9bb7a9d
    Image:          paulajuliabalza/log-output:1.0.0
    Image ID:       docker.io/paulajuliabalza/log-output@sha256:4c5b76683f3c5abb720da6f61995a7d0ee5e4f5293f7214122056eb2c30e7a61
    Port:           <none>
    Host Port:      <none>
    State:          Running
      Started:      Mon, 10 Feb 2025 19:24:49 -0300
    Ready:          True
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-2vv2p (ro)
Conditions:
  Type              Status
  Initialized       True 
  Ready             True 
  ContainersReady   True 
  PodScheduled      True 
Volumes:
  kube-api-access-2vv2p:
    Type:                    Projected (a volume that contains injected data from multiple sources)
    TokenExpirationSeconds:  3607
    ConfigMapName:           kube-root-ca.crt
    ConfigMapOptional:       <nil>
    DownwardAPI:             true
QoS Class:                   BestEffort
Node-Selectors:              <none>
Tolerations:                 node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                             node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  118s  default-scheduler  Successfully assigned default/log-output-67cb547fd-s4g5h to k3d-k3s-default-server-0
  Normal  Pulling    118s  kubelet            Pulling image "paulajuliabalza/log-output:1.0.0"
  Normal  Pulled     116s  kubelet            Successfully pulled image "paulajuliabalza/log-output:1.0.0" in 1.480108967s
  Normal  Created    116s  kubelet            Created container log-output
  Normal  Started    116s  kubelet            Started container log-output
```