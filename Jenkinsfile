pipeline {
    agent none
    stages {
        stage('Back-end') {
            agent any
            stages {
                stage('Build') {
                   steps {
                        echo "check docker compose"
                        step([$class: 'DockerComposeBuilder', dockerComposeFile: '/var/jenkins_home/workspace/fullstack-application/docker-compose.yml', option: [$class: 'StartAllServices'], useCustomDockerComposeFile: true])
                   }
                }
            }
        }

        stage('Front-End'){
            agent any
            stages {
                stage('Build') {
                    steps {
                        step([$class: 'DockerComposeBuilder', dockerComposeFile: '/var/jenkins_home/workspace/fullstack-application/docker-compose.yml', option: [$class: 'StopAllServices'], useCustomDockerComposeFile: true])
                    }
                }
            }
        }
    }
}