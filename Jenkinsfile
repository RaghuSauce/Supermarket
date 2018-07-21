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
        stage('temp run'){
            steps{
                sh 'docker run -d -p 8081:8081 supermarket_api:latest'
            }
        }
    }
    // pipeline{
    //     agent  any

    //     enviroment {
    //         GOPATH=''
    //     }
    //     stages{
    //         state ("prep workspace"){
    //             steps{
    //                 dir('src'){
    //                 }
    //             }
    //         }
    //         stage("pull"){
    //             steps{
    //                 sh "git pull https://github.com/RaghuSauce/SupermarketAPI.git"
    //             }
    //         stage("build staic bin"){
    //             steps{
    //                 sh "./build.sh"
    //             }
    //         }
    //         stage("build docker image"){
    //             steps{
    //                 sh "docker build -t supermarket_api:latest ."
    //             }
    //         }
    //         stage("run docker image"){
    //             steps{
    //                 sh "docker run -p 8081:8080 -d supermarket_api:latest"
    //             }
    //         }

    //         }
    //     }
    // }