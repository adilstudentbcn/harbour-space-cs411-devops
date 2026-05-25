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
                    sh '''
                    mkdir -p ~/.ssh
                    ssh-keyscan -H target >> ~/.ssh/known_hosts
                    scp -i $SSH_KEY main $SSH_USER@target:~
                    '''
                }
            }
        }
    }    


        
    
}

