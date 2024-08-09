# GRPC Example - Unary

To build the server/client containers use:
```
#Build Docker-based containers using specific Dockerfile
sudo docker build -f dockerfile.client -t grpc-client-op .
sudo docker build -f dockerfile.server -t grpc-server-op .
```
> The client container has a ENV variable to set the IP address of the server *SET_IPADDR*

```
sudo docker run -p 20001:20001 grpc-server-op
sudo docker run -e SET_IPADDR=IP_ADDR_TO_BE_SET -p 20002:20002 grpc-client-op 
```
