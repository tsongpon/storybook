pipeline {
    agent { docker { image 'instrumentisto/dep' } }
    stages {
        stage('pull dependencies') {
            steps {
                sh 'dep ensure'
            }
        }
        stage('unit test') {
            steps {
                sh 'go test ./... -v'
            }
        }
        stage('build image') {
            steps {
                sh 'docker build -t tsongpon/yoneebook .'
            }
        }
    }
}