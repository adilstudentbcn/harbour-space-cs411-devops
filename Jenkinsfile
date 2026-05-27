pipeline {
    agent any

    triggers {
        pollSCM('H/5 * * * *') 
    }

    tools {
       go "1.24.1"
    }

    stages {
        
        stage('Test') {
            steps {
                sh "go test ./..."
            }
        }
        
        stage('Build') {
            steps {
                sh "go build main.go"
            }
        }

        stage('Deploy') {
            steps {
                withCredentials([sshUserPrivateKey(
                    credentialsId: 'target-ssh', 
                    keyFileVariable: 'SSH_KEY', 
                    usernameVariable: 'SSH_USER'
                )]) {
                    sh 'ansible-playbook -i hosts.ini --private-key=$SSH_KEY --user=$SSH_USER playbook.yml'
                }
            }
        }
    }    
}