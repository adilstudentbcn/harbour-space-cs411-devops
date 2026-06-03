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
                sshagent(['aws-ec2-key']) {
                    sh "scp -o StrictHostKeyChecking=no ${APP_NAME} ${TARGET_USER}@${TARGET_IP}:/home/${TARGET_USER}/"
                    sh "ssh -o StrictHostKeyChecking=no ${TARGET_USER}@${TARGET_IP} './${APP_NAME} &'"
                }
            }
        }
    }
}
