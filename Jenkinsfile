pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('build') {
            steps {
                sh 'go get -u github.com/gin-gonic/gin'
                sh 'go run main.go'
            }
        }
    }
}