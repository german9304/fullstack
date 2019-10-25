pipeline {
    agent none
    stages {
        stage('Back-end') {
            agent any
            environment {
                PATH = "/usr/local/bin"
            }
            stages {
                stage('Build') {
                    steps {
                        echo 'Path is'
                        echo '$PATH'
                    }
                }
                stage('Test') {
                    agent {
                        docker {
                            image 'golang:1.12'
                        }
                    }
                    steps {
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