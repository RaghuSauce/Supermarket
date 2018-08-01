pipeline {
    agent any

    environment {
        GOPATH = "${env.WORKSPACE}"
    }

    stages {
        stage('checkout') {
            steps {
                dir('src/SupermarketAPI') {
                    checkout scm
                    sh 'go env'
                }
            }
        }
        stage('install depencies') {
            steps {
                dir('src/SupermarketAPI') {
                    sh 'go get'
                    sh 'go install'
                }
            }
        }
        stage('Build static bin') {
            steps {
                dir('src/SupermarketAPI') {
                    sh './build.sh'
                }
            }
        }
        stage('Build docker image') {
            steps {
                dir('src/SupermarketAPI') {
                    sh 'docker build -t supermarket_api:latest .'
                }
            }
        }
        stage('temp run') {
            steps {
                sh 'docker run -d -p 8081:8081 supermarket_api:latest'
            }
        }
    }
}