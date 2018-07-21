// new build style base on this blog
// https://medium.com/@zarkopafilis/building-a-ci-system-for-go-with-jenkins-4ab04d4bacd0
node {
    try {
        withEnv(["GOPATH=${env.WORKSPACE}"]) {
            stage('Check Environment') {
                //For debugging purposes (Check go version and path to working directory)
                sh 'go version'
                sh 'pwd'
            }

            stage('checkout') {
                steps {
                    dir('src/SupermarketAPI') {
                        checkout scm
                        sh 'go env'
                    }
                }
            }
            
            stage("install dependencies") {
                sh 'go get'
                sh 'go install'
            }
        }
    } catch (e) {
        //do something
    }
}






// pipeline {
//     agent any

//     environment {
//         GOPATH = "${env.WORKSPACE}"
//     }

//     stages {
//         stage('checkout'){
//             steps{
//                 dir('src/SupermarketAPI'){
//                     checkout scm
//                     sh 'go env'
//                 }
//             }
//         }
//         stage ('install depencies'){
//             steps{
//                 sh 'go get'
//                 sh 'go install'
//             }
//         }
//     }
// }

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