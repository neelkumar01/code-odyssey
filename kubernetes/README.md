<h1>Kubernetes <img width="40" height="40" alt="image" src="https://github.com/user-attachments/assets/de944b12-84fd-4480-aad3-481d32b39da5" />
</h1>


### Monolith

A **single, large, unified codebase or application** where all components (frontend, backend, business logic, database access, etc.) are tightly integrated and run as one unit.

### Microservices

A big application is divided into small services and each services handles a specific function

- node = server - stores data, run programs and provide services to clients over internet
- multi node = more that 1 server = cluster
- kubernetes manages cluster

### Master node servers

API server - communication gateway

- every operation pass through it
- handles request and communication b/w clusters
- update data in etcd

Scheduler

- schedule & run containers or pods at worker node

pod is where inside docker containers run

etcd

- data stored here - everything running in kubernetes cluster

Controller manager

- controls all nodes

### Servers of worker node

kubelet

- manages everything at worker node
- pods management
- give update to api server

service proxy

- as a user, we can access pods using api server

kubectl

- to give instructions to kubernetes cluster as an engineer

### CNI

CNI - computer network interface

- allow pods to connect to network
- so that pods can talk to each other

![image.png](attachment:5a4fabd0-e20a-4bb0-9299-d84ae29c011f:image.png)

Check docker is running or not

```bash
docker ps
```

Making my first yaml file

```yaml
key: value
name: Neel

skills:
  - Kubernetes
  - Ml
  - Golang

info:
  name: Neel
  age: 18
  university: dtu
  skills:
    - Kubernetes
    - Ml
    - Golang

```

Creating a cluster with a control plane and 3 worker nodes

Step1 : creating config.yml file

```yaml
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4

nodes:
- role: control-plane
  image: kindest/node:v1.34.0
- role: worker
  image: kindest/node:v1.34.0
- role: worker
  image: kindest/node:v1.34.0
- role: worker
  image: kindest/node:v1.34.0
  extraPortMappings:
  - containerPort: 80
    hostPort: 80
    protocol: TCP
  - containerPort: 443
    hostPort: 443
    protocol: TCP
   
```

- role - tells Kind what type of node to create
- Kind - helps creating kubernetes cluster on laptop using docker
- image - contains everything needed to run kubernetes node
- extraPortMappings

Without port mapping, your apps (like Nginx, Node.js, etc.) inside the cluster aren’t accessible from your local browser.

`extraPortMappings` exposes them

Step2 : creating cluster

In terminal,

```bash
kind create cluster --name=neel-cluster --config=config.yml
```

![image.png](attachment:f2647916-39b3-41e0-ae5f-738499a6a5dd:image.png)

To check all nodes

```bash
kubectl get nodes
```

kind: cluster 

this tells that kubernetes object type is cluster

means this represent a cluster or group of nodes working together

nodes is nothing but different servers

apiVersion: [kind.x-k8s.io/v1alpha4](http://kind.x-k8s.io/v1alpha4) 

this tells which version and rules kubernetes need to follow to create this object

here we have cluster as kubernetes object

 

role: control-plane

This helps in creating control plane

image: kindest/node:v1.34.0

Node is a server - physical or virtual

inside node - pods are running

pod contains 1 or more container

container are actual running thing

images are blueprint for container - dependencies, code, libraries

### Docker system memory management

```bash
docker system df
```

This shows space used by:

- Images
- Containers
- Volumes
- Build cache

🧹 One-Shot Full Cleanup (Safe for Kind + Docker)

```bash
# Delete all kind clusters
kind delete clusters --all

# Remove all images
docker rmi $(docker images -q)

# Remove only kind images
docker rmi kindest/node:v1.34.0

# Remove local volumes
docker volume rm $(docker volume ls -q)

# Clean up build cache & reclaim space
docker system prune -a --volumes -f

```

### Real world example

- we create a python web app
- Here, image is (python runtime + code + flask library)
- We execute this image to run our app - this is container - a running thing!
- container exists in pod
- pod exist in node or server (virtual or physical)

### Chatgpt example

`Image` is there which is like a ready package that contains (code + gpt logic + dependencies + libraries)

`Container` is where image is executing and running gpt model and code

`pod` is where 1 or more containers run

Lets assume each pod runs 2000 containers means for each pod, 2000 gpt models can be handled 

If chatgpt has 1,000,000 users then we need 1,000,000 / 2000 = 500 pods

`Deployment` - this is pod manager and it helps in creation of pods when

- there is crash
- we need to scale
- more users added and using the service

It makes sure required no. of pods keep running

`Service` - helps user in accessing pods 

All these things are resources and they run in node or server

### Namespace

```bash
kubectl get ns
```

For openai cluster:

| Namespace | App | Purpose |
| --- | --- | --- |
| chatgpt-dev | ChatGPT model dev version | Testing |
| chatgpt-prod | ChatGPT production version | Real users |
| billing | Payment APIs | Handle billing |
| monitoring | Prometheus, Grafana | Logs & metrics |

So ChatGPT is an app (pods, deployments, service)

It runs under namespace, like here it is `chatgpt-prod`

```bash
#output
NAME                 STATUS   AGE
default              Active   13h  # when no namespace created
kube-node-lease      Active   13h  # info of k8s nodes
kube-public          Active   13h  
kube-system          Active   13h
local-path-storage   Active   13h
```

```bash
kubectl get pods
>> No resources found in default namespace.
```

no pods are there in default

```bash
kubectl get pods -n kube-system
```

we get all pods in this namespace - “kube-system”

```bash
kubectl create ns nginx
>> namespace/nginx created
```

```bash
kubectl run nginx --image=nginx
>> pod/nginx created
```

pod crate in default namespace

```bash
kubectl delete pod nginx
>> pod "nginx" deleted from default namespace
```

creating pod in specific namespace

```bash
kubectl run nginx --image=nginx -n nginx
>> pod/nginx created

kubectl get pods -n nginx
>>
NAME    READY   STATUS              RESTARTS   AGE
nginx   0/1     ContainerCreating   0          3m24s
```

```bash
kubectl delete ns nginx
```

yml to create namespace

```yaml
kind: Namespace
apiVersion: v1
metadata:
  name: nginx
```

```bash
kubectl apply -f namespace.yml
```

### Create pod in namespace nginx

```yaml
kind: Pod
apiVersion: v1
metadata:
  name: nginx-pod
  namespace: nginx
spec:
  containers:
  - name: nginx
    image: nginx:latest
    ports:
    - containerPort: 80
```

```bash
kubectl apply -f pod.yml

kubectl get pods -n nginx
>>
NAME        READY   STATUS    RESTARTS   AGE
nginx-pod   1/1     Running   0          46s
```

Describing pod

```bash
kubectl describe pod/nginx-pod -n nginx
```

### Creating deployment

- helps in auto scaling when needed when a lot of users came to use our application
- helps in replication control so that whole traffic is distributed and managed by different pods

label - given to pod

selector - select that pod whose name given and replicate that pod

first delete our old pod

```bash
kubectl delete -f pod.yml
```

creating deployment.yml

```yaml
kind: Deployment
apiVersion: apps/v1
metadata:
  name: nginx-deployment
  namespace: nginx
spec:
  replica: 2
  selector:
    matchLabels:
      app: nginx
  
  template:
    metadata:
      labels:
        app: nginx

    spec:
      containers:
      - name: nginx
        image: nginx:latest
        ports:
        - containerPort: 80

```

Here template is the pod we want to replicate

now apply and make deployment

```bash
kubectl apply -f deployment.yml
>> deployment.apps/nginx-deployment created
```

checking

```bash
kubectl get deployment -n nginx
>>
NAME               READY   UP-TO-DATE   AVAILABLE   AGE
nginx-deployment   2/2     2            2           92s
```

Here 2 pods are ready

```bash
kubectl get pods -n nginx
>>
NAME                                READY   STATUS    RESTARTS   AGE
nginx-deployment-6f9664446b-xjs98   1/1     Running   0          3m55s
nginx-deployment-6f9664446b-z9bxq   1/1     Running   0          3m55s
```

Now if a traffic came and we want to scale then

```bash
kubectl scale deployment/nginx-deployment -n nginx --replicas=6
>> deployment.apps/nginx-deployment scaled
```

6 pods created

```bash
kubectl get pods -n nginx
>>
NAME                                READY   STATUS    RESTARTS   AGE
nginx-deployment-6f9664446b-fxcck   1/1     Running   0          108s
nginx-deployment-6f9664446b-jvwvv   1/1     Running   0          108s
nginx-deployment-6f9664446b-wdmsx   1/1     Running   0          108s
nginx-deployment-6f9664446b-xjs98   1/1     Running   0          7m35s
nginx-deployment-6f9664446b-xn2sz   1/1     Running   0          108s
nginx-deployment-6f9664446b-z9bxq   1/1     Running   0          7m35s
```

Now if we want to scale down and want only 3 pods then

```bash
kubectl scale deployment/nginx-deployment --replicas=3
>> deployment.apps/nginx-deployment scaled
```

### Rolling update

Imagine you have a **Deployment** running:

- 3 pods → all using image `myapp:v1`

Now you update the Deployment to use `myapp:v2`.

👉 What happens (rolling update style):

1. Kubernetes creates **1 new pod** (v2)
2. When it becomes **Ready**, it deletes **1 old pod** (v1)
3. Repeats until all v1 pods are replaced with v2 pods

✅ Result:

App keeps running — users never notice downtime

updating our container image to some previous version

```bash
kubectl set image deployment/nginx-deployment -n nginx nginx=nginx:1.28.0
```

This will update all containers to image `v1.28.0`

now updating container images to latest version

```bash
kubectl set image deployment/nginx-deployment -n nginx nginx=nginx:latest
```

### Replica set

delete previous deployment

```bash
kubectl delete -f deployment.yml
kubectl get pods -n nginx

>> No resources found in nginx namespace.
```

create `replicasets.yml`

```yaml
kind: ReplicaSet
apiVersion: apps/v1
metadata:
  name: nginx-replicasets
  namespace: nginx
spec:
  replicas: 2
  selector:
    matchLabels:
      app: nginx
  
  template:
    metadata:
      name: nginx-rep-pod
      labels:
        app: nginx

    spec:
      containers:
      - name: nginx
        image: nginx:latest
        ports:
        - containerPort: 80

```

now 

```bash
kubectl apply -f replicasets.yml
kubectl get replicasets -n nginx
>>
NAME                DESIRED   CURRENT   READY   AGE
nginx-replicasets   2         2         2       116s
```

A ReplicaSet makes sure that a fixed number of identical Pods are always running in your cluster.

Think of it like this:

> “I always want 3 nginx pods running — no matter what.”
> 

So Kubernetes will:

- Automatically create new Pods if one crashes.
- Delete extra Pods if there are too many.

### Daemon set

Run one pod per node, automatically — no matter how many nodes I have.

delete replica set

```bash
kubectl delete -f replicasets.yml
```

create daemonsets.yml

```yaml
kind: DaemonSet
apiVersion: apps/v1
metadata:
  name: nginx-daemonsets
  namespace: nginx
spec:
  selector:
    matchLabels:
      app: nginx
  
  template:
    metadata:
      name: nginx-dmn-pod
      labels:
        app: nginx

    spec:
      containers:
      - name: nginx
        image: nginx:latest
        ports:
        - containerPort: 80

```

now apply

```bash
kubectl apply -f daemonsets.yml

kubectl get pods -n nginx
>>
NAME                     READY   STATUS    RESTARTS   AGE
nginx-daemonsets-mqjk4   1/1     Running   0          12s
```

### Job

A Job in Kubernetes is used to run a task that should finish (run once or a set number of times) —

not something that runs forever like a web server.

completions - how many times to run job

parallelism - how many pods to run job

busybox - runs commands

create job.yml

```yaml
kind: Job
apiVersion: batch/v1
metadata:
  name: nginx-job
  namespace: nginx
spec:
  completions: 1
  parallelism: 1
  template:
    metadata:
      name: nginx-job-pod
      labels:
        app: pod-job
    spec:
      containers:
        - name: job-container
          image: busybox:latest
          command: ["sh", "-c", "echo hello kubernetes && sleep 10"]
      restartPolicy: Never
```

now

```bash
kubectl apply -f job.yml
>> job.batch/nginx-job created
```

### Cron jobs

task or job to schedule

```bash
touch cron-job.yml
```

now create yml

use “cron guru” for scheduling

```yaml
kind: CronJob
apiVersion: batch/v1
metadata:
  name: every-min-backup
  namespace: nginx

spec:
  schedule: "* * * * *"
  jobTemplate:
    spec:
      template:
        metadata:
          name: cron-job-pod
          labels:
            app: min-backup-pod
        spec:
          containers:
            - name: backup-container
              image: busybox
              command:
                - sh
                - -c
                - >
                  echo "Backup started";
                  mkdir -p /backups &&
                  mkdir -p /data &&
                  cp -r /data /backups &&
                  echo "Backup completed";
              volumeMounts:
                - name: data-volume
                  mountPath: /data
                - name: backup-volume
                  mountPath: /backups
          restartPolicy: OnFailure
          volumes:
            - name: data-volume
              hostPath:
                path: /data
                type: DirectoryOrCreate
            - name: backup-volume
              hostPath:
                path: /backups
                type: DirectoryOrCreate
```

- kind

tells kubernetes what type of object are we creating

- apiVersion: batch/v1

runs the current kubernetes api version to use

- kubernetes api

same thing that is there in master node

- metadata

tells info of kubernetes object

- spec

specifies how object should work and behave

- schedule

tell how and when job should run

- jobTemplate

blueprint of actual job work

we are creating a pod here

```yaml
    spec:
      template:
        metadata:
          name: cron-job-pod
          labels:
            app: min-backup-pod
```

volume - extra storage space used by containers

```yaml
          volumes:
            - name: data-volume
              hostPath:
                path: /data
                type: DirectoryOrCreate
            - name: backup-volume
              hostPath:
                path: /backups
                type: DirectoryOrCreate
```

hostPath - to directly use folder on server for storage

volumeMounts - attach volumes (storage spaces to containers)

```yaml
              volumeMounts:
                - name: data-volume
                  mountPath: /data
                - name: backup-volume
                  mountPath: /backups
```

now

```yaml
kubectl apply -f cron-job.yml
kubectl get cronjob -n nginx
```

### Storage

![image.png](attachment:8d2cc940-0554-45d7-8db5-d812bf8d6cf5:image.png)

- in servers lot if pods are running
- if one of the pod gets deleted, then whole data along with it also deletes
- persistent volume is the storage allocated on out host machine of some part from total storage
- we can use it to save our pod data forever even if pod gets deleted

creating persistent volume using yml file

```yaml
kind: PersistentVolume
apiVersion: v1
metadata:
  name: local-pv
  namespace: nginx
  labels:
    app: my-pv

spec:
  capacity:
    storage: 1Gi
  accessModes:
    - ReadWriteOnce
  persistentVolumeReclaimPolicy: Retain
  storageClassName: local-storage
  hostPath:
    path: /mnt/data

```

- created a persistent volume on our host machine of 1Gi (1Gi means 1Gb)
- accessModes here means,

| Access Mode | Short Code | Meaning |
| --- | --- | --- |
| 📝 **ReadWriteOnce (RWO)** | `ReadWriteOnce` | The volume **can be mounted as read-write by only one node** at a time. |
| 🔐 **ReadOnlyMany (ROX)** | `ReadOnlyMany` | The volume **can be mounted as read-only by many nodes** simultaneously. |
| 🧮 **ReadWriteMany (RWX)** | `ReadWriteMany` | The volume **can be mounted as read-write by many nodes** simultaneously. |
- PersistentVolumeReclaimPolicy: Retain

To retain data of pod in pv

- storage class name - tells about type of storage for pv
- hostPath - path where 1Gi storage is allocated

now apply

```bash
kubectl apply -f persistent-volume.yml
>> persistentvolume/local-pv created

kubectl get pv
>>
NAME       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS      CLAIM   STORAGECLASS    VOLUMEATTRIBUTESCLASS   REASON   AGE
local-pv   1Gi        RWO            Retain           Available           local-storage   <unset>                          38s
```

here, status : available

👉 means pv is now available and made but we now need to claim it (to use it)

creating another yml for claiming

```yaml
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: local-pvc
  namespace: nginx

spec:
  accessModes:
   - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
  storageClassName: local-storage
  
```

- under resources, we are requesting the amount of storage to claim

now apply

```bash
kubectl apply -f persistent-volume-claim.yml

kubectl get pv
>>
NAME       CAPACITY   ACCESS MODES   RECLAIM POLICY   STATUS   CLAIM               STORAGECLASS    VOLUMEATTRIBUTESCLASS   REASON   AGE
local-pv   1Gi        RWO            Retain           Bound    default/local-pvc   local-storage   <unset>                          9m33s
```

No status is bound

making deployment and adding pv in it

```yaml
kind: Deployment
apiVersion: apps/v1
metadata:
  name: nginx-deployment
  namespace: nginx
spec:
  replicas: 2
  selector:
    matchLabels:
      app: nginx
  
  template:
    metadata:
      labels:
        app: nginx

    spec:
      containers:
      - name: nginx
        image: nginx:latest
        ports:
        - containerPort: 80
        volumeMounts:
        - mountPath: /var/www/html
          name: my-volume
      volumes:
        - name: my-volume
          persistentVolumeClaim:
            claimName: local-pvc

```

volume mounts and mount path

- asking pod to create a volume named “my-volume”
- then we give the path where our pod data stored to my-volume - this path is mountPath
- then we connect my-volume with pvc

### Service

It allows you to **access Pods reliably**, even if Pods **keep changing** (because Pods are temporary)

creating service using yml

```yaml
kind: Service
apiVersion: v1
metadata:
  name: nginx-service
  namespace: nginx

spec:
  selector:
    app: nginx
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
  type: ClusterIP
```

- selector

helps in connecting to pod that have label `app : nginx`

- port, nodePort, targetPort

| Field | Meaning |
| --- | --- |
| `port` | **The port on the Service**. Other Pods inside the cluster talk to this port. |
| `targetPort` | **The port on the Pod/container** where your app (e.g., Nginx) is actually listening. |

now check all resources

```bash
kubectl apply -f service.yml
kubectl get all -n nginx
>>
NAME                                    READY   STATUS      RESTARTS   AGE
pod/every-min-backup-29350908-lmcqz     0/1     Completed   0          2m18s
pod/every-min-backup-29350909-t5khp     0/1     Completed   0          78s
pod/every-min-backup-29350910-h4dxs     0/1     Completed   0          18s
pod/nginx-daemonsets-mqjk4              1/1     Running     0          2d4h
pod/nginx-deployment-6dc7f95c58-8pth6   1/1     Running     0          84m
pod/nginx-deployment-6dc7f95c58-z9qmw   1/1     Running     0          84m

NAME                    TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
service/nginx-service   ClusterIP   10.96.167.197   <none>        80/TCP    61s

NAME                              DESIRED   CURRENT   READY   UP-TO-DATE   AVAILABLE   NODE SELECTOR   AGE
daemonset.apps/nginx-daemonsets   1         1         1       1            1           <none>          2d4h

NAME                               READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/nginx-deployment   2/2     2            2           84m

NAME                                          DESIRED   CURRENT   READY   AGE
replicaset.apps/nginx-deployment-6dc7f95c58   2         2         2       84m

NAME                             SCHEDULE    TIMEZONE   SUSPEND   ACTIVE   LAST SCHEDULE   AGE
cronjob.batch/every-min-backup   * * * * *   <none>     False     0        18s             29h

NAME                                  STATUS     COMPLETIONS   DURATION   AGE
job.batch/every-min-backup-29350908   Complete   1/1           5s         2m18s
job.batch/every-min-backup-29350909   Complete   1/1           5s         78s
job.batch/every-min-backup-29350910   Complete   1/1           5s         18s
```

now run the service

```bash
kubectl port-forward svc/nginx-service 8080:80 -n nginx
```

![image.png](attachment:7ed981c7-7eda-4ffe-bfee-9922d5cf37bc:a5afdfdb-6c5c-402c-88c4-228a7e98b96e.png)

![image.png](attachment:ad38d2f5-2701-4ca2-a5c1-4598c5a7d47e:image.png)

- `kubectl` → tool to talk to your Kubernetes cluster
- `port-forward` → make a tunnel from your computer to the cluster.
- `svc/nginx-service` → the Service you want to connect to (`svc/` = Service).
- `8080:80` → map your computer port 8080 → Service port 80
    - 8080 = what you type in your browser
    - 80 = the port inside Kubernetes where Nginx is running
- `n nginx` → the namespace where your Service is (like a folder for your apps)

### Implementing things

Create a basic website with html, css, js

![image.png](attachment:ee4c8c57-2ee5-49c7-a68b-244a6e2f6687:image.png)

creating a docker image for web app

```docker
# Use lightweight Python image
FROM python:3.10-slim

# Set working directory
WORKDIR /app

# Copy all frontend files
COPY . .

# Expose port 8000 for frontend
EXPOSE 8000

# Start a lightweight web server
CMD ["python", "-m", "http.server", "8000"]
```

- When you build a Docker image for your website, you are basically packaging your entire website + everything it needs to run (like Nginx or Python) into one self-contained box called an image.
- That image can then be run anywhere — on your laptop, a server, or the cloud — and it will behave exactly the same way

create image

```bash
docker build -t web-app-k8s .

docker images
>>
REPOSITORY     TAG       IMAGE ID       CREATED         SIZE
web-app-k8s    latest    021241de4ba0   9 seconds ago   209MB
kindest/node   v1.34.0   7416a61b42b1   7 weeks ago     1.48GB
```

now sign up on docker hub and generate personal access token

This is used to login on our system

Tagging and pushing to docker hub

If you want to upload (push) your image to Docker Hub, you must tag it with your Docker Hub username and repository name

```bash
docker tag web-app-k8s:latest neelkumar01/web-app-k8s:latest
docker images
>>
REPOSITORY                TAG       IMAGE ID       CREATED          SIZE
neelkumar01/web-app-k8s   latest    021241de4ba0   21 minutes ago   209MB
web-app-k8s               latest    021241de4ba0   21 minutes ago   209MB
kindest/node              v1.34.0   7416a61b42b1   7 weeks ago      1.48GB
```

push image to docker hub

```bash
docker push neelkumar01/web-app-k8s:latest
```

Now our image is available on docker hub so we can use it in our pod

create `/k8s` folder

- create deployment.yml

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: web-app-deployment
  namespace: web-app
  labels:
    app: web-app
spec:
  replicas: 1
  selector:
    matchLabels:
      app: web-app
  template:
    metadata:
      labels:
        app: web-app
    spec:
      containers:
      - name: web-app
        image: neelkumar01/web-app-k8s
        ports:
        - containerPort: 8000

```

- create namespace.yml

```yaml
kind: Namespace
apiVersion: v1
metadata:
  name: web-app
```

- create service.yml

```yaml
kind: Service
apiVersion: v1
metadata:
  name: web-app-service
  namespace: web-app
spec:
  selector:
    app: web-app
  ports:
    - protocol: TCP
      port: 8000
      targetPort: 8000
  type: ClusterIP
```

- apply all yml files

```bash
kubectl apply -f namespace.yml
kubectl apply -f deployment.yml
kubectl apply -f service.yml

kubectl get pods -n web-app
>>
NAME                                 READY   STATUS    RESTARTS   AGE
web-app-deployment-dd46b4854-bfk9f   1/1     Running   0          80s
```

- launching the service live to be accessed on browser

```bash
kubectl port-forward service/web-app-service -n web-app 8000:8000 -n web-app
```

![image.png](attachment:6f45d025-0288-4231-9e9b-53dd2e621b14:image.png)

### Ingress

This is used for rerouting services

![image.png](attachment:47a190b1-fa66-49d3-a629-1df1c846636d:image.png)

- to use this, we need to remove web-app deployment and services from namespace web-app and put both into nginx namespace
- with this, both our services of nginx and web-app will run in same namespace

```bash
kubectl get deployment -n web-app
>> 
NAME                 READY   UP-TO-DATE   AVAILABLE   AGE
web-app-deployment   1/1     1            1           6h40m

kubectl delete deployment/web-app-deployment -n web-app
kubectl delete service/web-app-service -n web-app
```

now change namespace to nginx then

```bash
kubectl apply -f deployment.yml -f service.yml
```

now we have both service in same namespace

apply ingress nginx

```bash
kubectl apply -f https://kind.sigs.k8s.io/examples/ingress/deploy-ingress-nginx.yaml
```

create ingress.yml in nginx

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: my-ingress
  namespace: nginx
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  ingressClassName: nginx
  rules:
  - http:
      paths:
      - pathType: Prefix
        path: /nginx
        backend:
          service:
            name: nginx-service
            port:
              number: 80
      - pathType: Prefix
        path: /app
        backend:
          service:
            name: web-app-service
            port:
              number: 8000
```

now apply

```bash
kubectl apply -f ingress.yml
kubectl port-forward --namespace ingress-nginx service/ingress-nginx-controller 8080:80
```

- use of annotations

When someone opens

👉 `http://example.com/app`

Then the Ingress sends that exact path `/app` to your backend pod or service.

So your backend receives: /app

But your backend web app might only work on the root path

So it doesn’t understand /app → gives “404 Not Found”

What annotation do is,

Whenever you see `/app`, send it to `/` instead.

- `ingressClassName: nginx` → This Ingress will be handled by the Nginx Ingress Controller

for /app

![image.png](attachment:449bd560-02c0-4b20-9978-e5aff2588bcb:image.png)

for /nginx

![image.png](attachment:dcd2e2d6-3fb6-4a11-a834-19f231c98d64:image.png)

### Stateful Sets

Starting fresh, delete nginx namespace

```bash
kubectl delete -f namespace.yml
kubectl get pods -n nginx
>> No resources found in nginx namespace

cd ..
mkdir mysql
cd mysql
```

- Stateless → “Doesn’t remember anything.”
    
    Every time you send a request, it treats it like the first time. Example: a basic website that shows the same page to everyone
    
- Stateful → “Remembers stuff.”
    
    It keeps track of what happened before. Example: your shopping cart in an online store 
    

create namespace mysql

```yaml
kind: Namespace
apiVersion: v1
metadata:
  name: mysql
```

apply it

```bash
kubectl apply -f namespace.yml
```

<aside>
💡

Important info for statefulset.yml

</aside>

- labels in template.metadata
    - This is tagging your pod with a name.
    - Think of it like putting a sticker on your pod: app=web-app.
    - Kubernetes uses this to identify the pod later.
- selector.matchLabels
    - This tells Kubernetes which pods to manage.
    - It matches the pods that have the same label (app=web-app).
    - Basically: “Hey Kubernetes, only control the pods that have this sticker.”

- env - environment variables
    - MySQL sees the `MYSQL_ROOT_PASSWORD` env variable → sets root password.
    - MySQL sees `MYSQL_DATABASE` → creates a database called my-database

These are key-value pairs that are injected into the container at runtime.

Your containerised application can read these values when it starts.

```yaml
kind: StatefulSet
apiVersion: apps/v1
metadata:
  name: mysql-statefulset
  namespace: mysql

spec:
  serviceName: mysql-service
  replicas: 3
  selector:
    matchLabels:
      app: mysql

  template:
    metadata:
      labels:
        app: mysql
    spec:
      containers:
      - name: mysql
        image: mysql:8.0
        ports:
        - containerPort: 3306
        env:
        - name: MYSQL_ROOT_PASSWORD
          value: root
        - name: MYSQL_DATABASE
          value: my-database
        volumeMounts:
        - name: mysql-data
          mountPath: /var/lib/mysql

  volumeClaimTemplates:
  - metadata:
      name: mysql-data
    spec:
      accessModes: [ "ReadWriteOnce" ]
      resources:
        requests:
          storage: 1Gi
```

- service.yml - headless service as clusterIP is none because we don’’t want our database to expose to other user

```yaml
kind: Service
apiVersion: v1
metadata:
  name: mysql-service
  namespace: mysql
spec:
  clusterIP: None
  selector:
    app: mysql
  ports:
  - name: mysql
    protocol: TCP
    port: 3306
    targetPort: 3306
    
```

apply them

```bash
kubectl apply -f service.yml -f statefulset.yml
kubectl get pods -n mysql
>>
NAME                  READY   STATUS    RESTARTS   AGE
mysql-statefulset-0   1/1     Running   0          15m
mysql-statefulset-1   1/1     Running   0          62s
mysql-statefulset-2   1/1     Running   0          58s
```

going into mysql service we created

```bash
kubectl exec -it mysql-statefulset-0 -n mysql -- bash
bash-5.1# mysql -u root -p
Enter password: root
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 8
Server version: 8.0.44 MySQL Community Server - GPL

Copyright (c) 2000, 2025, Oracle and/or its affiliates.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql>
```

now in mysql,

```sql
show databases;
>>
+--------------------+
| Database           |
+--------------------+
| information_schema |
| my-database        |
| mysql              |
| performance_schema |
| sys                |
+--------------------+
5 rows in set (0.01 sec)
```

There is database with name `my-database` this is what we created using yml file

if we delete, any pod

```sql
kubectl delete pod mysql-statefulset-0 -n mysql
kubectl get pods -n mysql
>>
NAME                  READY   STATUS    RESTARTS   AGE
mysql-statefulset-0   1/1     Running   0          10s
mysql-statefulset-1   1/1     Running   0          11m
mysql-statefulset-2   1/1     Running   0          11m
```

as we delete, a pod with same name is created

this is how states are managed here with stateful sets

### Config maps

- Store plain text configuration data (like environment variables, URLs, file paths, etc.) separately from your app code.
- Why: So you can change app settings without rebuilding the container image.

create configMap.yml

```yaml
kind: ConfigMap
apiVersion: v1
metadata:
  name: mysql-config-map
  namespace: mysql
data:
  MYSQL_DATABASE: my-database
```

in statefulset.yml change part of code with this

```yaml
        - name: MYSQL_DATABASE
          valueFrom:
            configMapKeyRef:
              name: mysql-config-map
              key: MYSQL_DATABASE
```

apply chnages

```bash
kubectl apply -f configMap.yml -f statefulset.yml
```

### Secrets

- Purpose: Store confidential data like passwords, tokens, certificates, etc.
- Why: These should not be visible in plain YAML or logs

create secret.yml

```yaml
kind: Secret
apiVersion: v1
metadata:
  name: mysql-secret
  namespace: mysql
data:
  MYSQL_ROOT_PASSWORD: cm9vdAo=  #password in base64 encoding: "cm9vdAo=" means "root"
```

similarly make changes in statefulset.yml

```yaml
        - name: MYSQL_ROOT_PASSWORD
          valueFrom:
            secretKeyRef:
              name: mysql-secret
              key: MYSQL_ROOT_PASSWORD
```

### Resource quotas

![image.png](attachment:2deff8ba-39f6-4c42-8e84-a469f954318a:image.png)

- we have 1 pod of nginx and 1 pod of mysql
- if traffic came on nginx pod and it consumes all 4Gi space, nothing lefts for mysql pod
- so we manage this by providing resource quotas

we can do so by adding this code in our deployment yml file

```yaml
        resources:
          requests:
            cpu: 100m
            memory: 128Mi
          limits:
            cpu: 200m
            memory: 256Mi
```

Here,

- `cpu: 100m`  - this means 0.1 cpu core

0.1 cpu core - to use 10% processing ability of 1 core

- memory of 128Mi means to use 128 mebibytes (~134 Mb) of RAM

final deployment.yml

```yaml
kind: Deployment
apiVersion: apps/v1
metadata:
  name: nginx-deployment
  namespace: nginx
spec:
  replicas: 2
  selector:
    matchLabels:
      app: nginx
  
  template:
    metadata:
      labels:
        app: nginx

    spec:
      containers:
      - name: nginx
        image: nginx:latest
        ports:
        - containerPort: 80
        resources:
          requests:
            cpu: 100m
            memory: 128Mi
          limits:
            cpu: 200m
            memory: 256Mi
        volumeMounts:
        - mountPath: /var/www/html
          name: my-volume
      volumes:
        - name: my-volume
          persistentVolumeClaim:
            claimName: local-pvc

```

### Probes

Probes are health checks that Kubernetes runs inside a container to check if it’s healthy, alive, and ready to serve traffic

- liveness probe - check if container is alive and running
- readiness probe - check if container is read to handle traffic
- startup probe - check if container is successfully started up

to use, add this in deployment

 

```yaml
        livenessProbe:
          httpGet:
            path: /
            port: 8000
        readinessProbe:
          httpGet:
            path: /
            port: 8000
```

final deployment.yml

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: web-app-deployment
  labels:
    app: web-app
  namespace: nginx
spec:
  replicas: 1
  selector:
    matchLabels:
      app: web-app
  template:
    metadata:
      labels:
        app: web-app
    spec:
      containers:
      - name: web-app
        image: neelkumar01/web-app-k8s
        ports:
        - containerPort: 8000
        livenessProbe:
          httpGet:
            path: /
            port: 8000
        readinessProbe:
          httpGet:
            path: /
            port: 8000
```

now check

```bash
kubectl apply -f service.yml
kubectl apply -f deployment.yml
kubectl get pods -n nginx
>>
NAME                                  READY   STATUS    RESTARTS   AGE
nginx-deployment-5b48d69c8b-bnsf9     1/1     Running   0          22m
nginx-deployment-5b48d69c8b-fd5t5     1/1     Running   0          22m
web-app-deployment-5b8bf8cccb-tk57j   1/1     Running   0          14s

kubectl describe pod/web-app-deployment-5b8bf8cccb-tk57j -n nginx
```

### Taints and tolerations

- taint - to tell kubernetes scheduler not to schedule pod to a particular worker node
- toleration - forcing pod to run on a tainted server

```bash
# updating worker nodes from 1 to 3
kind get clusters
nginx-cluster

kind delete cluster --name nginx-cluster
>> Deleting cluster "nginx-cluster" ...

kind create cluster --name nginx-cluster --config config.yml

kubectl get nodes
>>
NAME                          STATUS   ROLES           AGE     VERSION
nginx-cluster-control-plane   Ready    control-plane   7m18s   v1.34.0
nginx-cluster-worker          Ready    <none>          7m4s    v1.34.0
nginx-cluster-worker2         Ready    <none>          7m4s    v1.34.0
nginx-cluster-worker3         Ready    <none>          7m4s    v1.34.0

# tainting all worker node

kubectl taint node nginx-cluster-worker prod=true:NoSchedule
kubectl taint node nginx-cluster-worker2 prod=true:NoSchedule
kubectl taint node nginx-cluster-worker3 prod=true:NoSchedule

kubectl apply -f namespace.yml
kubectl apply -f pod.yml
kubectl get pods -n nginx
>>
NAME        READY   STATUS    RESTARTS   AGE
nginx-pod   0/1     Pending   0          10s
```

- status here is pending because all pods are tainted and pod can’t be scheduled to any of worker node
- control plane is always tainted

untainting worker node

```bash
kubectl taint node nginx-cluster-worker prod=true:NoSchedule-
>> node/nginx-cluster-worker untainted

kubectl get pods -n nginx
>>
NAME        READY   STATUS    RESTARTS   AGE
nginx-pod   1/1     Running   0          5m7s
```

now our pod is running !

```bash
# again taint all nodes
kubectl taint node nginx-cluster-worker prod=true:NoSchedule

# delete pod
kubectl delete -f pod.yml
>> pod "nginx-pod" deleted from nginx namespace
```

applying toleration

add the code in pod.yml to implement toleration

```yaml
  tolerations:
  - key: "prod"
    operator: "Equal"
    value: "true"
    effect: "NoSchedule"
```

final pod.yml

```yaml
kind: Pod
apiVersion: v1
metadata:
  name: nginx-pod
  namespace: nginx
spec:
  containers:
  - name: nginx
    image: nginx:latest
    ports:
    - containerPort: 80
  tolerations:
  - key: "prod"
    operator: "Equal"
    value: "true"
    effect: "NoSchedule"
```

apply everything

```bash
kubectl apply -f pod.yml
>> pod/nginx-pod created

kubectl get pods -n nginx
>>
NAME        READY   STATUS    RESTARTS   AGE
nginx-pod   1/1     Running   0          9s
```

even in tainted nodes, pod is running

### HPA

HPA - horizontal pod autoscaling

- You have a web app running with 2 pods.
- Traffic increases → CPU usage goes above 80%.
- HPA automatically increases pods from 2 → 5.
- When traffic drops, it scales back down to 2 pods

VPA - vertical pod autoscaling

- here we don’t increase the number of pods
- we increase the resources and limits of pod

• If you are using a Kind cluster install Metrics Server

```bash
kubectl apply -f https://github.com/kubernetes-sigs/metrics-server/releases/latest/download/components.yaml
```

• Edit the Metrics Server Deployment

```bash
kubectl -n kube-system edit deployment metrics-server
```

• Add the security bypass to deployment under `container.args`

```yaml
- --kubelet-insecure-tls
- --kubelet-preferred-address-types=InternalIP,Hostname,ExternalIP
```

• Restart the deployment

```bash
kubectl -n kube-system rollout restart deployment metrics-server
```

• Verify if the metrics server is running

```yaml
kubectl get pods -n kube-system
kubectl top nodes
```

check metrics

```bash
kubectl top pod -n nginx
>>
NAME        CPU(cores)   MEMORY(bytes)   
nginx-pod   0m           8Mi         
```

```bash
cd ..
mkdir apache
cd apache
```

create namespace apache

- yml file

```yaml
kind: Namespace
apiVersion: v1
metadata:
  name: apache
```

- apply

```bash
kubectl apply -f namespace.yml
```

create deployment yaml file

```yaml
kind: Deployment
apiVersion: apps/v1
metadata:
  name: apache-deployment
  namespace: apache
spec:
  replicas: 1
  selector:
    matchLabels:
      app: apache
  template:
    metadata:
      name: apache
      labels:
        app: apache
    spec:
      containers:
      - name: apache
        image: httpd:latest
        ports:
        - containerPort: 80
        resources:
          requests:
            cpu: 100m
            memory: 128Mi
          limits:
            cpu: 200m
            memory: 256Mi
```

apply and check

```bash
kubectl apply -f deployment.yml
>> deployment.apps/apache-deployment created

kubectl get pods -n apache
>>
NAME                                 READY   STATUS    RESTARTS   AGE
apache-deployment-5db86f8c66-5prxp   1/1     Running   0          55s
```

create service yaml file

```yaml
kind: Service
apiVersion: v1
metadata:
  name: apache-service
  namespace: apache
spec:
  selector:
    app: apache
  ports:
  - protocol: TCP
    port: 80  # exposed port in cluster
    targetPort: 80  # container port
  type: ClusterIP
```

now,

```bash
kubectl get all -n apache
>>
NAME                                     READY   STATUS    RESTARTS   AGE
pod/apache-deployment-5db86f8c66-5prxp   1/1     Running   0          72m
NAME                     TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
service/apache-service   ClusterIP   10.96.221.217   <none>        80/TCP    65m
NAME                                READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/apache-deployment   1/1     1            1           72m
NAME                                           DESIRED   CURRENT   READY   AGE
replicaset.apps/apache-deployment-5db86f8c66   1         1         1       72m

kubectl scale deployment apache-deployment -n apache --replicas=3
```

create hpa yaml file for horizontal pod autoscaling

- scaleTargetRef - tells which thing to scale

```yaml
apiVersion: autoscaling/v2
kind: HorizontalPodAutoscaler
metadata:
  name: apache-hpa
  namespace: apache
spec:
  scaleTargetRef:
    kind: Deployment
    name: apache-deployment
    apiVersion: apps/v1

  minReplicas: 1
  maxReplicas: 5

  metrics:
    - type: Resource
      resource:
        name: cpu
        target:
          type: Utilization
          averageUtilization: 5
```

now apply

```bash
kubectl apply -f hpa.yml
kubectl get hpa -n apache
>>
NAME         REFERENCE                      TARGETS      MINPODS   MAXPODS   REPLICAS   AGE
apache-hpa   Deployment/apache-deployment   cpu: 1%/5%   1         5         1          17s
```

here we have current pod utilising cpu 1% out of given 5%

- if traffic increase cpu utilisation will increase
- then pods also autoscale automatically

### VPA

VPA - vertical pod autoscaler

install pre requisites

```bash
git clone https://github.com/kubernetes/autoscaler.git

ls
>>
autoscaler	deployment.yml	hpa.yml		namespace.yml	service.yml

cd autoscaler
ls
>>
addon-resizer			LICENSE
balancer			multidimensional-pod-autoscaler
builder				OWNERS
cluster-autoscaler		OWNERS_ALIASES
code-of-conduct.md		README.md
CONTRIBUTING.md			SECURITY_CONTACTS
hack				vertical-pod-autoscaler

cd vertical-pod-autoscaler

./hack/vpa-up.sh
```

create vpa yaml file

```yaml
kind: VerticalPodAutoscaler
apiVersion: autoscaling.k8s.io/v1
metadata:
  name: apache-vpa
  namespace: apache

spec:
  targetRef:
    name: apache-deployment
    apiVersion: apps/v1
    kind: Deployment
  updatePolicy:
    updateMode: "Auto"
```

monitor and check

```bash
kubectl apply -f vpa.yml

kubectl get vpa -n apache
>>
NAME         MODE   CPU   MEM     PROVIDED   AGE
apache-vpa   Auto   25m   250Mi   True       76m
```

- here cpu and memory will increase if traffic increases on our pod

### Node affinity

Kubernetes checks the labels on each node, and if they match your affinity rules → the Pod can be scheduled there

### RBAC

RBAC - role based access control

Without RBAC, anyone could access or modify anything in your cluster — dangerous!

With RBAC, you can:

- Limit access (principle of least privilege)
- Create developer/admin/operator roles
- Protect critical resources

![image.png](attachment:5e312af5-a2bc-4b52-a2c0-cd169121d94c:image.png)

1. service account
- it is kubernetes managed account
- used by pods
- operate at namespace level
- role and role bindings are used

1. user
- real person like admin / developer
- operate at cluster level
- cluster roles and cluster role bindings are used

let’s start with basic setup

using namespace “apache”

```bash
# delete all resources in apache
kubectl delete -f .

kubectl auth whoami
>> 
ATTRIBUTE        VALUE
Username         kubernetes-admin
Groups           [kubeadm:cluster-admins system:authenticated]

kubectl auth can-i get pods
>> yes

kubectl apply -f namespace.yml

kubectl apply -f deployment.yml

# This is risky !
kubectl auth can-i delete deployment
>> yes
```

create role yaml file

```yaml
kind: Role
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: apache-manager
  namespace: apache
rules:
  - apiGroups: ["*"]
    resources: ["*"]
    verbs: ["*"]
```

create service-account yaml file

```yaml
kind: ServiceAccount
apiVersion: v1
metadata:
  name: apache-user
  namespace: apache
```

now

```bash
kubectl get serviceaccount -n apache
>>
NAME          SECRETS   AGE
apache-user   0         19s
default       0         3h15m

kubectl auth can-i get pods --as=apache-user -n apache
>> no
```

Here we can see permission to get pods as apache-user is denied because we have not binded role with service-account

create role-binding yaml file

```yaml
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: apache-manager-rolebinding
  namespace: apache
subjects:
- kind: User
  name: apache-user
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role
  name: apache-manager
  apiGroup: rbac.authorization.k8s.io

```

- Give the user named `apache-user` the permissions we’re about to reference
- Bind the Role named `apache-manager` to the user `apache-user` in the `apache` namespace

```bash
kubectl get rolebinding -n apache
>>
NAME                         ROLE                  AGE
apache-manager-rolebinding   Role/apache-manager   3m42s

# now check
kubectl auth can-i get pods --as=apache-user -n apache
>> yes

kubectl auth can-i get deployment --as=apache-user -n apache
>> yes
```

### Monitoring and logging

```bash
cd ..
mkdir dashboard
cd dashboard

# Deploy the Dashboard Apply the Kubernetes Dashboard manifest:
kubectl apply -f https://raw.githubusercontent.com/kubernetes/dashboard/v2.7.0/aio/deploy/recommended.yaml
```

create dashboard-admin-user YAML file

```yaml
kind: ServiceAccount
apiVersion: v1
metadata:
  name: admin-user
  namespace: kubernetes-dashboard

---

kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1
metadata:
  name: admin-user-binding
  namespace: kubernetes-dashboard

subjects:
- kind: ServiceAccount
  name: admin-user
  namespace: kubernetes-dashboard
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
```

```bash
kubectl apply -f dashboard-admin-user.yml

# create token to access dashboard
kubectl -n kubernetes-dashboard create token admin-user
>> eyJhbGciOiJSUzI1NiIsImtpZCI6InNRNGV0MThBT3RxUlAtOFBLTDhZMFl4clEtLTZuQ1hBZ3BldzIyM1VVeDgifQ.eyJhdWQiOlsiaHR0cHM6Ly9rdWJlcm5ldGVzLmRlZmF1bHQuc3ZjLmNsdXN0ZXIubG9jYWwiXSwiZXhwIjoxNzYxMzU0ODY0LCJpYXQiOjE3NjEzNTEyNjQsImlzcyI6Imh0dHBzOi8va3ViZXJuZXRlcy5kZWZhdWx0LnN2Yy5jbHVzdGVyLmxvY2FsIiwianRpIjoiZGQ1NmU3NWUtMTg1Ni00ZWM0LTkxYzMtMjQxZGNjMGEzNmQzIiwia3ViZXJuZXRlcy5pbyI6eyJuYW1lc3BhY2UiOiJrdWJlcm5ldGVzLWRhc2hib2FyZCIsInNlcnZpY2VhY2NvdW50Ijp7Im5hbWUiOiJhZG1pbi11c2VyIiwidWlkIjoiNWZmMmUyOWQtOTJmZi00YzJhLThlYTMtN2RkZTE2YmNiZTg1In19LCJuYmYiOjE3NjEzNTEyNjQsInN1YiI6InN5c3RlbTpzZXJ2aWNlYWNjb3VudDprdWJlcm5ldGVzLWRhc2hib2FyZDphZG1pbi11c2VyIn0.F5GXETVBdDGRDvXIDQ0YTN81cP7iI59IR1cvO7pLyNtKvye9rWbLDJMGDMB45Q_8n1Zs2PTfe65rHqjK046DG9f5g8ycMMoz8O84njxsWq0SBqGTs0ZXxI1UU_GgQwK-WLjh5rfcfYYuRkdiDYjYRizD_5zMxijltBiOdTk9eWKy2xbzVYEO1oMXWrqHLV6GLZJVtNbyNpLk504ZfqtIqQ7FR5FF-kDp1lPzZgLvSlGcmtxViFozFRSRx12fHBzkQvuIpwG473HgsGwAcNRRX7hai-XLUZtNJQiHMBpnRz4lqqUrXMFuD_TemwjjT0FQHsOllx5gzFmJxI1uWVos5w

#Access the Dashboard Start the Dashboard using kubectl proxy:
kubectl proxy
```

Open the Dashboard in your browser at this url,

http://localhost:8001/api/v1/namespaces/kubernetes-dashboard/services/https:kubernetes-dashboard:/proxy/

### CRD

CRD - custom resource definition

- when we want to create our own custom resources which are not present in kubernetes

```bash
mkdir crd
cd crd
touch devops-crd.yml
```

crd yaml file

```yaml
kind: CustomResourceDefinition
apiVersion: apiextensions.k8s.io/v1
metadata:
  name: crdboxes.crdgroup.kube
spec:
  group: crdgroup.kube
  names:
    plural: crdboxes
    singular: crdbox
    kind: CrdBoxes
    shortNames:
      - cb
      - box
  scope: Namespaced
  versions:
  - name: v1
    served: true
    storage: true
    schema:
      openAPIV3Schema:
        type: object
        properties:
          spec:
            type: object
            properties:
              name: 
                type: string
              weight:
                type: string
              color:
                type: string
              material:
                type: string
```

```bash
kubectl get crd
>>
NAME                          CREATED AT
crdboxes.crdgroup.kube        2025-10-25T00:44:52Z

# use your crd
touch cr.yml
```

yaml file to use our custom resource

```yaml
kind: CrdBoxes
apiVersion: crdgroup.kube/v1
metadata:
  name: black-box
spec:
  name: secret-black-box
  weight: 5 kg
  color: black
  material: cardboard
  
  ---

kind: CrdBoxes
apiVersion: crdgroup.kube/v1
metadata:
  name: blue-box
spec:
  name: secret-blue-box
  weight: 12 kg
  color: blue
  material: cardboard  
```

```bash
kubectl apply -f cr.yml

kubectl get crdboxes
>>
NAME        AGE
black-box   2m30s
blue-box    10s
```

### Init container and sidecar container

Init container

- inside pod there is always a main container
- now there are some requirements or dependencies that are needed to run/make/create main container
- all those things run in init container

Sidecar container

- we have a pod and inside it main container is there
- now there are some things which needs to be stored or any other type of resource that needs to be handled side by side
- all this is done inside sidecar container
