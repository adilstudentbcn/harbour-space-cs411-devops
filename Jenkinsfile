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
                    
                    # 1. Stop the running service to unlock the file
                    ssh -i $SSH_KEY $SSH_USER@target 'sudo systemctl stop main.service'
                    
                    # 2. Safely copy the new binary
                    scp -i $SSH_KEY main $SSH_USER@target:~
                    
                    # 3. Start the service back up
                    ssh -i $SSH_KEY $SSH_USER@target 'sudo systemctl start main.service'
                    '''
                }
            }
        }
    }    


        
    
}

