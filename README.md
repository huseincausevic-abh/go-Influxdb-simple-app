# Go-Influxdb-simple-app

This is just a simple API written in go programming language that uses InfluxDB as a data storage.

Prerequisites:
- Kubernetes cluster
- kubectl command-line tool for communication with the cluster

# Deployment

To deploy this simple app go inside k8s folder and execute deploy shell script. Script has two parameters -n for namespace and -m for mode.
Valid mode is 'apply', 'delete' and 'recreate'.

```
cd k8s
bash deploy.sh -n <NAMESPACE> -m <MODE>
```
