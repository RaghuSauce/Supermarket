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

        stage('unit test') {
            steps {
                dir('src/SupermarketAPI') {
                    sh 'go test ./...'
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

        stage('Integration Test') {
            steps {
                sh 'docker run --rm -d -p 8081:8081 supermarket_api'
                sh 'go test supermarket_service/handlers_integration_test.go -integration'
                sh 'docker stop supermarket_api'
            }

        }
        // TODO version images properly
        stage('Publish to Dockerhub') {
            withCredwithCredentials([
                [$class: 'UsernamePasswordMultiBinding', credentialsId: 'DockerHubLogin', usernameVariable: 'USERNAME', passwordVariable: 'PASSWORD']
            ])
            steps {
                sh 'docker login --username=$USERNAME --password=$PASSWORD'
                sh 'docker tag supermarket_api $USERNAME/supermarket_api:latest'
                sh 'docker push supermarket_api'   
            }

        }
    }
}