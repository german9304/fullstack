pipeline {
    agent none
    stages {
        stage('Back-end') {
            agent any
            stages {
                stage('Build') {
                    steps {
                        sh '/usr/local/bin/docker-compose --version'
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