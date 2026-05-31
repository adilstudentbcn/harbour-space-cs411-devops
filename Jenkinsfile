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
        stage('Deploy to Docker VM') {
            steps {
                sh "ssh -o StrictHostKeyChecking=no docker 'docker pull ${IMAGE}'"
                sh "ssh -o StrictHostKeyChecking=no docker 'docker rm -f myapp || true'"
                sh "ssh -o StrictHostKeyChecking=no docker 'docker run -d -p 4444:4444 --name myapp ${IMAGE}'"
            }
        }
    }
}
