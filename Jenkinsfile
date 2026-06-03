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
                    // Copy to /tmp first, then move with sudo to ensure permissions are handled
                    sh "scp -o StrictHostKeyChecking=no -i ${SSH_KEY} ${APP_NAME} ${TARGET_USER}@${TARGET_IP}:/tmp/"
                    sh "ssh -o StrictHostKeyChecking=no -i ${SSH_KEY} ${TARGET_USER}@${TARGET_IP} 'mv /tmp/${APP_NAME} /home/${TARGET_USER}/${APP_NAME} && chmod +x /home/${TARGET_USER}/${APP_NAME} && nohup /home/${TARGET_USER}/${APP_NAME} > app.log 2>&1 &'"
                }
            }
        }
    }
}
