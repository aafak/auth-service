# auth-service
Handles user authentication and authorization.

# Install postgres
https://github.com/aafak/dev-setup/blob/main/postgres/install%20postgres_using_k8s.md

# Build
```
aafak@aafak-virtual-machine:~/go_apps/auth-service$ ls
cmd  Dockerfile  go.mod  go.sum  internal  Jenkinsfile  LICENSE  Makefile  README.md
aafak@aafak-virtual-machine:~/go_apps/auth-service$  sudo docker build --build-arg PROXY=http://proxy.corp.net:8080 -t aafak/auth-service:latest .
```

# Create a container

```
aafak@aafak-virtual-machine:~/go_apps/auth-service$ sudo docker run -d -p 8080:8080 aafak/auth-service:latest
55e26637e120baf45d6f07ab2103aef14c2ce8d41a0494b7e78257dc1d933743
aafak@aafak-virtual-machine:~/go_apps/auth-service$ docker ps
CONTAINER ID   IMAGE                                 COMMAND                  CREATED         STATUS         PORTS                                                                                                                                  NAMES
55e26637e120   aafak/auth-service:latest             "./main"                 6 seconds ago   Up 5 seconds   0.0.0.0:8080->8080/tcp, :::8080->8080/tcp                                                                                              happy_spence
0947afc9cbca   gcr.io/k8s-minikube/kicbase:v0.0.45   "/usr/local/bin/entr…"   2 weeks ago     Up 2 weeks     127.0.0.1:32782->22/tcp, 127.0.0.1:32781->2376/tcp, 127.0.0.1:32780->5000/tcp, 127.0.0.1:32779->8443/tcp, 127.0.0.1:32778->32443/tcp   minikube
aafak@aafak-virtual-machine:~/go_apps/auth-service$ docker logs -f 55e26637e120

2024/10/08 05:49:13 /app/internal/repository/postgres.go:41
[error] failed to initialize database, got error failed to connect to `host=192.168.49.2 user=aafak database=authz`: dial error (timeout: dial tcp 192.168.49.2:32738: connect: connection timed out)
Failed to connect to DB
[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /register                 --> github.com/aafak/auth-service/internal/handler.UserHandler.RegisterUser-fm (3 handlers)
[GIN-debug] GET    /users                    --> github.com/aafak/auth-service/internal/handler.UserHandler.GetUser-fm (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080
[GIN] 2024/10/08 - 05:49:30 | 404 |       1.268µs |      172.18.0.1 | GET      "/userscurl"
[GIN] 2024/10/08 - 05:49:30 | 404 |         784ns |      172.18.0.1 | GET      "/"
[GIN] 2024/10/08 - 05:49:32 | 404 |       1.191µs |      172.18.0.1 | GET      "/"
[GIN] 2024/10/08 - 05:49:36 | 404 |         988ns |      172.18.0.1 | GET      "/"
Getting user....
[GIN] 2024/10/08 - 05:49:41 | 200 |     1.27052ms |      172.18.0.1 | GET      "/users"
Getting user....
[GIN] 2024/10/08 - 05:49:46 | 200 |     287.233µs |      172.18.0.1 | GET      "/users"


aafak@aafak-virtual-machine:~$ curl http://localhost:8080/users
{"id":1,"username":"aafak"}
aafak@aafak-virtual-machine:~$


aafak@aafak-virtual-machine:~$ curl -X POST http://localhost:8080/register -H "Content-Type: application/json" -d '{"username": "user1", "password": "psw1"}'
aafak@aafak-virtual-machine:~$

```


# Test with helm
Install helm using: https://github.com/aafak/dev-setup/tree/main/k8s/helm

# Install auth service using helm
```
aafak@aafak-virtual-machine:~/go_apps/auth-service$ ls
cmd  Dockerfile  go.mod  go.sum  helm  internal  Jenkinsfile  LICENSE  Makefile  README.md

aafak@aafak-virtual-machine:~/go_apps/auth-service$ ls helm/
Chart.yaml  templates  values-prod.yaml  values.yaml
aafak@aafak-virtual-machine:~/go_apps/auth-service$

aafak@aafak-virtual-machine:~/go_apps/auth-service$ ls helm/templates/
configmap.yaml  deployment.yaml  service.yaml
aafak@aafak-virtual-machine:~/go_apps/auth-service$


aafak@aafak-virtual-machine:~/go_apps/auth-service$ helm install  auth-service ./helm -f ./helm/values.yaml
NAME: auth-service
LAST DEPLOYED: Tue Oct  8 12:52:15 2024
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None

aafak@aafak-virtual-machine:~$ helm list
NAME            NAMESPACE       REVISION        UPDATED                                 STATUS          CHART                  APP VERSION
auth-service    default         1               2024-10-08 12:52:15.200480116 +0530 IST deployed        auth-service-0.0.0     0.0.0
aafak@aafak-virtual-machine:~$
```

# Verify deployments
```
aafak@aafak-virtual-machine:~/go_apps/auth-service$ kubectl get svc
NAME           TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
auth-service   ClusterIP   10.98.104.53    <none>        8080/TCP         54s
kubernetes     ClusterIP   10.96.0.1       <none>        443/TCP          15d
postgres       NodePort    10.96.153.222   <none>        5433:32738/TCP   14d

aafak@aafak-virtual-machine:~$ kubectl get cm
NAME                 DATA   AGE
kube-root-ca.crt     1      15d
postgres-configmap   10     27m
aafak@aafak-virtual-machine:~$

aafak@aafak-virtual-machine:~/go_apps/auth-service$ kubectl get deploy
NAME           READY   UP-TO-DATE   AVAILABLE   AGE
auth-service   1/1     1            1           60s
postgres       1/1     1            1           14d
aafak@aafak-virtual-machine:~/go_apps/auth-service$


aafak@aafak-virtual-machine:~/go_apps/auth-service$ kubectl get pods
NAME                            READY   STATUS    RESTARTS   AGE
auth-service-56dd954cfb-lghhw   1/1     Running   0          48s
postgres-7768954946-7bl85       1/1     Running   0          14d

aafak@aafak-virtual-machine:~/go_apps/auth-service$ kubectl logs -f auth-service-56dd954cfb-lghhw

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /register                 --> github.com/aafak/auth-service/internal/handler.UserHandler.RegisterUser-fm (3 handlers)
[GIN-debug] GET    /users                    --> github.com/aafak/auth-service/internal/handler.UserHandler.GetUser-fm (3 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8080

Getting user....
[GIN] 2024/10/08 - 07:37:36 | 200 |     556.541µs |      10.244.0.1 | GET      "/users"
registring user....
[GIN] 2024/10/08 - 07:40:05 | 200 |   19.203618ms |      10.244.0.1 | POST     "/register"

```

# Test the auth service
```
aafak@aafak-virtual-machine:~/go_apps/auth-service$ kubectl get svc
NAME           TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
auth-service   ClusterIP   10.98.104.53    <none>        8080/TCP         54s
kubernetes     ClusterIP   10.96.0.1       <none>        443/TCP          15d
postgres       NodePort    10.96.153.222   <none>        5433:32738/TCP   14d

Change clusterIP to NodePort
aafak@aafak-virtual-machine:~/go_apps/auth-service$ kubectl edit svc auth-service

aafak@aafak-virtual-machine:~$ kubectl get svc
NAME           TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
auth-service   NodePort    10.98.104.53    <none>        8080:30847/TCP   14m
kubernetes     ClusterIP   10.96.0.1       <none>        443/TCP          15d
postgres       NodePort    10.96.153.222   <none>        5433:32738/TCP   15d
aafak@aafak-virtual-machine:~$ minikube ip
192.168.49.2
aafak@aafak-virtual-machine:~$ curl http://192.168.49.2:30847/users
{"id":1,"username":"aafak"}
aafak@aafak-virtual-machine:~$


aafak@aafak-virtual-machine:~$ curl -X POST http://192.168.49.2:30847/register -H "Content-Type: application/json" -d '{"username": "user1", "password": "psw1"}'
{"message":"User registered successfully"}
aafak@aafak-virtual-machine:~$


aafak@aafak-virtual-machine:~$ kubectl get svc
NAME           TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
auth-service   NodePort    10.98.104.53    <none>        8080:30847/TCP   14m
kubernetes     ClusterIP   10.96.0.1       <none>        443/TCP          15d
postgres       NodePort    10.96.153.222   <none>        5433:32738/TCP   15d
aafak@aafak-virtual-machine:~$ minikube ip
192.168.49.2
aafak@aafak-virtual-machine:~$ psql -h 192.168.49.2 -p 32738 -d authz
Password for user aafak: test
psql (12.20 (Ubuntu 12.20-0ubuntu0.20.04.1), server 13.16 (Debian 13.16-1.pgdg120+1))
WARNING: psql major version 12, server major version 13.
         Some psql features might not work.
Type "help" for help.

authz=#
authz=# select * from users;
 id | username | password
----+----------+----------
  1 | user1    | psw1
(1 row)

authz=#


aafak@aafak-virtual-machine:~$  curl -X POST http://192.168.49.2:30847/register -H "Content-Type: application/json" -d '{"username": "user2", "password": "psw2"}'
{"message":"User registered successfully"}
aafak@aafak-virtual-machine:~$

authz=# select * from users;
 id | username | password
----+----------+----------
  1 | user1    | psw1
  2 | user2    | psw2
(2 rows)

authz=#

```

# Verify enviroment variable
```
aafak@aafak-virtual-machine:~$ kubectl get pods
NAME                            READY   STATUS    RESTARTS   AGE
auth-service-56dd954cfb-lghhw   1/1     Running   0          22m
postgres-7768954946-7bl85       1/1     Running   0          15d

aafak@aafak-virtual-machine:~$ kubectl exec -it auth-service-56dd954cfb-lghhw -- env
PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
HOSTNAME=auth-service-56dd954cfb-lghhw
TERM=xterm
LOG_LEVEL=info
APP_ENV=production
CONFIG_JSON={
  "version": "1.0",
  "settings": {
    "theme": "light",
    "max_retries": 3
  }
}

DB_USER=aafak
POSTGRES_URL=postgresql://localhost:5432/mydb
DB_PORT=5433
DB_NAME=authz
DEBUG_MODE=false
DATABASE_URL=postgresql://localhost:5432/mydb
DB_HOST=postgres
DB_PASSWORD=test
POSTGRES_SERVICE_NAME=postgres
KUBERNETES_PORT_443_TCP_PROTO=tcp
POSTGRES_PORT_5433_TCP_PORT=5433
AUTH_SERVICE_SERVICE_PORT=8080
AUTH_SERVICE_SERVICE_PORT_REST=8080
AUTH_SERVICE_PORT_8080_TCP_ADDR=10.98.104.53
AUTH_SERVICE_PORT=tcp://10.98.104.53:8080
AUTH_SERVICE_PORT_8080_TCP_PROTO=tcp
AUTH_SERVICE_PORT_8080_TCP_PORT=8080
KUBERNETES_SERVICE_HOST=10.96.0.1
POSTGRES_SERVICE_HOST=10.96.153.222
POSTGRES_PORT=tcp://10.96.153.222:5433
POSTGRES_PORT_5433_TCP=tcp://10.96.153.222:5433
AUTH_SERVICE_SERVICE_HOST=10.98.104.53
KUBERNETES_PORT_443_TCP_ADDR=10.96.0.1
KUBERNETES_SERVICE_PORT_HTTPS=443
KUBERNETES_PORT=tcp://10.96.0.1:443
KUBERNETES_PORT_443_TCP=tcp://10.96.0.1:443
KUBERNETES_PORT_443_TCP_PORT=443
POSTGRES_SERVICE_PORT=5433
POSTGRES_PORT_5433_TCP_PROTO=tcp
POSTGRES_PORT_5433_TCP_ADDR=10.96.153.222
AUTH_SERVICE_PORT_8080_TCP=tcp://10.98.104.53:8080
KUBERNETES_SERVICE_PORT=443
HOME=/root
aafak@aafak-virtual-machine:~$

```
