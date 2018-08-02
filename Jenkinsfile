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
                    sh 'docker build -t raghusauce011/supermarketchallange:latest .'
                }
            }
        }

        stage('Integration Test') {
            steps {
                sh 'docker run --name supermarket_api --rm -d -p 8081:8081 raghusauce011/supermarketchallange:latest'
                sh 'go test supermarket_service/handlers_integration_test.go -integration'
                sh 'docker stop supermarket_api'
            }

        }
        // TODO version images properly
        stage('Publish to Dockerhub') {
            steps {
                gitHash=sh 'git show -s --format=%h'
                withDockerRegistry([credentialsId: "DockerHubLogin", url: ""]) {
                    sh 'docker push raghusauce011/supermarketchallange:latest'
                    sh 'docker push raghusauce011/supermarketchallange:${gitHash}'
                }
            }
        }
    }
}