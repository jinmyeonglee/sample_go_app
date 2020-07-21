pipeline {
    agent {
        docker {
            image 'golang'
            args '-p 8080:8080'
        }
    }
    stages {
        stage('Pre Test') {
            steps {
                echo 'Installing Dependencies'
                sh 'go get -u github.com/gin-gonic/gin'
            }
        }
        stage('build') {
            steps {
                sh 'go build'
            }
        }
    }
}