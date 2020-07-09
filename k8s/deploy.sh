#!/bin/bash
set -
while getopts ":m:n:" opt; do
  case $opt in
    m) MODE=$OPTARG ;;
    n) NAMESPACE=$OPTARG ;;
  esac
done

GREEN='\033[0;32m'
RED='\033[0;31m'
WHITE='\033[0m'

die () {
   echo -e >&2 "$@"
   echo -e ${WHITE}
   exit 1
 }

function check_if_pod_is_running() {
  retry=15
  for((i=0; i<${retry}; i+=1)); do
  number_of_replicas=$(kubectl get statefulset -n ${NAMESPACE} | grep influxdb-demo | awk '{print $2}')
  if [[ ${number_of_replicas} == "1/1" ]]; then
    echo -e "${GREEN}InfluxDB pod is ready!${WHITE}"
    return 0
  fi
  echo "Waiting for InfluxDB pod to be ready..."
  sleep 5
  done
  echo -e "${RED}Waited for ${retry} retries but InfluxDB pod is not ready!"
  return 1
}

delimeter() {
  echo -e "${GREEN}-------------------------------------------------------------------------------------------------------${WHITE}"
}

if [[ ${MODE} = "" ]]; then
   die "${RED}Unexpected that parameter mode (-m) is not provided or empty"
fi

delimeter
echo -e "${GREEN}Using mode: ${MODE} on resources..."
delimeter

if [[ ${NAMESPACE} == "" ]]; then
  echo "Namespace is not provided, using default namespace...";
  NAMESPACE="default";
fi

if [[ ${MODE} == "delete" || ${MODE} == "recreate" ]]; then
  delimeter
  echo -e "${GREEN}Deleting resources..."
  delimeter
  kubectl delete -f influxdb-secret.yaml -n ${NAMESPACE};
  kubectl delete -f influxdb-service.yaml -n ${NAMESPACE};
  kubectl delete -f influxdb-statefulset.yaml -n ${NAMESPACE};
  kubectl delete -f influxdb-auth-job.yaml -n ${NAMESPACE};
  kubectl delete -f app-service.yaml -n ${NAMESPACE}
  kubectl delete -f app-deployment.yaml -n ${NAMESPACE}
fi

if [[ ${MODE} == "apply" || ${MODE} == "recreate" ]]; then
  delimeter
  echo -e "${GREEN}Applying InfluxDB resources..."
  delimeter
  kubectl apply -f influxdb-secret.yaml -n ${NAMESPACE};
  kubectl apply -f influxdb-service.yaml -n ${NAMESPACE};
  kubectl apply -f influxdb-statefulset.yaml -n ${NAMESPACE};
  check_if_pod_is_running
  if [[ $? == 1 ]]
  then
    die
  fi
  kubectl apply -f influxdb-auth-job.yaml -n ${NAMESPACE};
  kubectl apply -f app-service.yaml -n ${NAMESPACE}
  sleep 2
  kubectl apply -f app-deployment.yaml -n ${NAMESPACE}
fi

delimeter
