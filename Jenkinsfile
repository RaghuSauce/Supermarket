pipeline {
    agent any

    environment {
        GOPATH = "${env.WORKSPACE}"
    }

    stages {
        stage('Announce') {
            steps {
                slackSend(color: '#00FF00', message: "Starting: Job '${env.JOB_NAME} [${env.BUILD_NUMBER}]' (${env.BUILD_URL})")
            }
        }

        stage('checkout') {
            steps {
                dir('src/SupermarketChallenge') {
                    checkout scm
                    sh 'go env'
                }
            }
        }

        stage('install dependencies') {
            steps {
                dir('src/SupermarketChallenge') {
                    sh 'go get'
                    sh 'go install'
                }
            }
        }

        stage('unit test') {
            steps {
                dir('src/SupermarketChallenge') {
                    sh 'go test ./...'
                }
            }
        }

        stage('Build static bin') {
            steps {
                dir('src/SupermarketChallenge') {
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

                dir('src/SupermarketChallenge') {
                    script {
                        versionNum = readFile('VERSION')
                        version = sh([script: "${versionNum}"+= "|" += "${gitHash}"])
                    }
                    sh 'docker build -t raghusauce011/supermarketchallenge:latest .'
                    sh "docker tag raghusauce011/supermarketchallenge:latest raghusauce011/supermarketchallenge:${version}"
                }

            }
        }

        stage('Integration Test') {
            steps {
                sh 'docker run --name supermarket_api --rm -d -p 8081:8081 raghusauce011/supermarketchallenge:latest'
                sh 'go test smservice/smHandlers_integration_test.go'
                sh 'docker stop supermarket_api'
            }

        }
        stage('Publish to Dockerhub') {
            steps {
                withDockerRegistry([credentialsId: "DockerHubLogin", url: ""]) {
                    sh 'docker push raghusauce011/supermarketchallenge:latest'
                    sh "docker push raghusauce011/supermarketchallenge:${gitHash}"
                }

            }
        }
        stage('Deploy to GKE') {
            steps {
                sh "echo ${gitHash}"
                sh "kubectl set image deployment/supermarket-api-deployment supermarket-api-deployment=raghusauce011/supermarketchallenge:${gitHash}"
            }
        }
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
    }
}