pipeline {
    agents none
    stages {
        stage('Back-End') {
            agent {
                docker {
                    image 'golang:1.12'
                }
            }
            steps {
                sh 'go version'
            }
        }
        stage('Front-End'){
            agent {
                image 'node:10.16'
            }
            steps {
                sh 'node -v'
            }
        }
    }
}