# Excercise 1.02: Project v0.1

Project can be done with any language and framework you want

The project will be a simple todo application with the familiar features of create, read, update, and delete (CRUD). We'll develop it during all parts of this course. Check the title of the exercise for "Project vX.Y" to know it is about building the project.

Todo is a text like "I need to clean the house" that can be in state of not-done or done.

![My project](project.png)

Dashed lines separate major differences across the course. Some exercises are not included in the picture. The connections between most pods are not included as well. You are free to do them however you want.

Keep this in mind if you want to avoid doing more work than necessary.

Let's get started!

Create a web server that outputs "Server started in port NNNN" when it is started and deploy it into your Kubernetes cluster. Please make it so that an environment variable PORT can be used to choose that port. You will not have access to the port when it is running in Kubernetes yet. We will configure the access when we get to networking.


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