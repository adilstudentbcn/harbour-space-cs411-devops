pipeline {
    agent any
    environment {
        IMAGE = "ttl.sh/adilstudentbcn:2h"
    }
    stages {
        stage('Build Image') {
            steps {
                sh "docker build -t ${IMAGE} ."
            }
        }
        stage('Push Image') {
            steps {
                sh "docker push ${IMAGE}"
            }
        }
        stage('Deploy to Kubernetes') {
            steps {
                withCredentials([string(credentialsId: 'k8s-token', variable: 'K8S_TOKEN')]) {
                    sh "kubectl --server=https://kubernetes:6443 --insecure-skip-tls-verify --token=$K8S_TOKEN delete pod myapp --ignore-not-found=true"
                    sh "kubectl --server=https://kubernetes:6443 --insecure-skip-tls-verify --token=$K8S_TOKEN apply -f pod.yaml"
                }
            }
        }
    }
}
