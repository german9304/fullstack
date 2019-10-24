pipeline {
    agent none
    stages {
        stage('Back-End') {
            agent {
                docker {
                    image 'golang:1.12'
                }
            }
            steps {
                echo 'running here'
            }
            stages {
                stage('Build') {
                    steps {
                        echo 'running here'
                        sh 'go version'
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
            steps {
                sh 'node -v'
                echo 'Checking directory'
                sh './frontend/scripts/build.sh'
            }
        }
    }
}