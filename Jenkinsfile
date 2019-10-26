pipeline {
    agent none
    stages {
        stage('Back-end') {
            agent any
            stages {
                stage('Build') {
                   steps {
                        sh './backend/scripts/build.sh'
                        echo "check docker compose"
                        sh '/usr/local/bin/docker-compose --version'
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