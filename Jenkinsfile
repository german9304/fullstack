pipeline {
    agent none
    stages {
        stage('Back-End') {
            parallel {
                stage('Start') {
                    agent {
                        docker {
                            image 'golang:1.12'
                        }
                    }
                    stages {
                        stage('Build') {
                            steps {
                                echo 'build stage'
                                sh 'go version'
                            }
                        }
                        stage('Test') {
                            steps {
                                echo 'test stage'
                            }
                        }
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