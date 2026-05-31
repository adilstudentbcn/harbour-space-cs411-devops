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
            agent { label 'docker' }
            steps {
                sh "docker pull ${IMAGE}"
                sh "docker rm -f myapp || true"
                sh "docker run -d -p 4444:4444 --name myapp ${IMAGE}"
            }
        }
    }
}
