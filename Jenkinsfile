pipeline {
    agent any
    environment {
        TARGET_IP   = '13.60.54.247'
        TARGET_USER = 'ubuntu'
        APP_NAME    = 'myapp'
    }
    stages {
        stage('Build') {
            steps {
                sh "go build -o ${APP_NAME} ."
            }
        }
        stage('Deploy') {
            steps {
                withCredentials([sshUserPrivateKey(credentialsId: 'aws-ec2-key', keyFileVariable: 'SSH_KEY')]) {
                    sh "scp -o StrictHostKeyChecking=no -i ${SSH_KEY} ${APP_NAME} ${TARGET_USER}@${TARGET_IP}:/home/${TARGET_USER}/"
                    sh "ssh -o StrictHostKeyChecking=no -i ${SSH_KEY} ${TARGET_USER}@${TARGET_IP} 'chmod +x ${APP_NAME} && ./${APP_NAME} &'"
                }
            }
        }
    }
}
