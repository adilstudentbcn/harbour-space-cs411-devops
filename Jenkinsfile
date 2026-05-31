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
                withCredentials([sshUserPrivateKey(credentialsId: 'docker-ssh', keyFileVariable: 'SSH_KEY', usernameVariable: 'SSH_USER')]) {
                    sh '''
                    ssh -i $SSH_KEY -o StrictHostKeyChecking=no $SSH_USER@docker "docker pull ${IMAGE}"
                    ssh -i $SSH_KEY -o StrictHostKeyChecking=no $SSH_USER@docker "docker rm -f myapp || true"
                    ssh -i $SSH_KEY -o StrictHostKeyChecking=no $SSH_USER@docker "docker run -d -p 4444:4444 --name myapp ${IMAGE}"
                    '''
                }
            }
        }
    }
}
