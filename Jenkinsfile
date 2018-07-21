pipeline {
    agent any

    environment {
        GOPATH = "${env.WORKSPACE}"
    }

    stages {
        dir('src/SupermarketAPI') {
            stage('checkout') {
                steps {

                    checkout scm
                    sh 'go env'

                }
            }
            stage('install depencies') {
                steps {
                    sh 'go get'
                    sh 'go install'
                }

            }
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