pipeline {
    agent {
        dockerfile true
    }
    stages {
//         stage('Pre Test') {
//             steps {
//                 echo 'Installing Dependencies'
//                 sh 'go get -u github.com/gin-gonic/gin'
//             }
//         }
        stage('build') {
            steps {
                sh 'go get -u github.com/gin-gonic/gin'
                sh 'go build'
            }
        }
    }
}