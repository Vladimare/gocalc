# Scalable Golang calculator server

## Create deployment

```shell
minikube start --driver=docker
minikube addons enable ingress
docker build . --tag gocalc:1.0.0
minikube image load gocalc:1.0.0
kubectl create ns go-calc
kubectl apply -f gocalc.yaml
kubectl apply -f ingress.yaml
```

## Check the configuration

### Check the IP address

```shell
kubectl get ingress -n go-calc --watch
```

### Set the domain name

```shell
sudo echo '<ingress ip address> go.calc' >> /etc/hosts
```

### Connect to the application

Use the following pattern for the calculations:
`http://go.calc/plus?a=4&b=3`, where `plus` is an operation, `a` and `b` are operands in this addition, i.e. here we calculate `4+3`.
For other operations, use `minus`, `multiply` and `divide` respectively.

## Delete deployment

```shell
kubectl delete ingress gocalc-ingress -n go-calc
kubectl delete service go-calc -n go-calc
kubectl delete deployment gocalc -n go-calc
```
