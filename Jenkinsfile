pipeline {
    agent none
    stages {
        stage('Back end') {
            agent {
                docker {
                    image 'golang:1.12'
                }
            }
            stages {
                stage('Build') {
                    agent {
                        docker {
                            image 'node:10.16'
                        }
                    }
                    steps {
                        echo 'build stage'
                        sh 'node -v'
                    }
                }
                stage('Test') {
                    steps {
                        echo 'test step'
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