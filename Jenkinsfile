pipeline {
    agent none
    stages {
        stage('Back-end') {
            agent {
                docker {
                    image 'golang:1.12'
                }
            }
            stages {
                stage('Build') {
                   steps {
                        sh './backend/scripts/build.sh'
                        echo "check docker compose"
                        sh './backend/scripts/test.sh'
                   } 
                }
            }
        }

        stage('Front-End'){
            agent {
                docker {
                    image 'node:10.16'
                    args '-p 3000:3000'
                }
            }
            stages {
                stage('Build') {
                    steps {
                        sh 'node -v'
                        echo 'Checking directory'
                        sh './frontend/scripts/build.sh'
                    }
                }
            }
        }
    }
}