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
    }
}
