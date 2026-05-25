pipeline {
    agent any

    triggers {
        pollSCM('H/5 * * * *')   // check the repo every ~5 min
    }

    tools { go "1.24.1" }

    stages {
        stage('Build') {
            steps { sh "go build main.go" }
        }
    }
}
