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
                script {
                    gitHash = sh([script: "git show -s --format=%h", returnStdout: true]).trim()
                    echo "GitHash:${gitHash}"
                }
                dir('src/SupermarketAPI') {
                    sh 'docker build -t raghusauce011/supermarketchallange:latest .'
                    sh "docker tag raghusauce011/supermarketchallange:latest raghusauce011/supermarketchallange:${gitHash}"
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
                withDockerRegistry([credentialsId: "DockerHubLogin", url: ""]) {
                    sh 'docker push raghusauce011/supermarketchallange:latest'
                    sh "docker push raghusauce011/supermarketchallange:${gitHash}"
                }

            }
        }
    }

    pre{
           slackSend(color: '#00FF00', message: "Starting: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' (${env.BUILD_URL})")
    }
    post {
        // only triggered when blue or green sign
        success {
            slackSend(color: '#00FF00', message: "SUCCESSFUL: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' (${env.BUILD_URL})")
        }
        // triggered when red sign
        failure {
            slackSend(color: '#FF0000', message: "FAILED: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' (${env.BUILD_URL})")
        }
        // trigger every-works
    }
}