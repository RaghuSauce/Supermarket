pipeline{
    agent  anu
    stages{
        stage("pull"){
            steps{
                sh "git pull https://github.com/RaghuSauce/SupermarketAPI.git"
            }
        stage("build staic bin"){
            steps{
                sh "./build.sh"
            }
        }
        stage("build docker image"){
            steps{
                sh "docker build -t supermarket_api:latest ."
            }
        }
        stage("run docker image"){
            steps{
                sh "docker run -p 8081:8080 -d supermarket_api:latest"
            }
        }
        
        }
    }
}