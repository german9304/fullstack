pipeline {
    agent none
    stages {
        stage('Back end') {
            stages {
                stage('Build') {
                    steps {
                         echo 'start container'
                    }

                    steps {
                        sh 'docker-compose -up -d'
                    }
                }
                stage('Test') {
                    steps {
                        sh 'docker ps'
                    }

                    steps {
                        echo 'running container'
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