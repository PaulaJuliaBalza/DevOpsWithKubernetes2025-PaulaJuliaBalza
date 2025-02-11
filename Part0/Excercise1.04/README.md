# Excercise 1.04: Project v0.2
Create a deployment.yaml for the project.

You won't have access to the port yet but that'll come soon.

# Solution
```bash
   #Build and push the Docker image  
   docker build -t paulajuliabalza/todoapp:latest .
   docker push paulajuliabalza/todoapp:latest
   #Apply the Deployment 
   kubectl apply -f deployment.yaml
   administrator@ASC:~/Escritorio/DevOpsWithKubernetes2025-PaulaJuliaBalza/Part0/Excercise1.02$ kubectl get pods
   NAME                         READY   STATUS    RESTARTS   AGE
   todoapp-65fbb59b68-jvhvj     1/1     Running   0          8s
   administrator@ASC:~/Escritorio/DevOpsWithKubernetes2025-PaulaJuliaBalza/Part0/Excercise1.02$ kubectl logs -f todoapp-65fbb59b68-jvhvj
   2025/02/10 21:29:45 Server started in port 8080
   ```