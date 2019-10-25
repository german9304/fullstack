pipeline {
    agent none
    stages {
        stage('Back end') {
            agent none
            stages {
                stage('Build') {
                    steps {
                        echo 'start docker container'
                        sh 'docker-compose up -d'
                    }
                }
                stage('Test') {
                    steps {
                        echo 'running container'
                        sh 'docker ps'
                    }
                }
            }
        }

        stage('middle') {
            steps {
                sh 'docker-compose stop'
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