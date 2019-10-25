pipeline {
    agent none
    stages {
        stage('Back-end') {
            agent any
            stages {
                stage('Build') {
                    agent {
                        docker {
                            image 'node:10.16'
                        }
                    }
                    steps {
                        sh 'yarn run prisma deploy'
                        sh 'yarn run prisma generate'
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