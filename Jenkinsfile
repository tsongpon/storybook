pipeline {
    agent { docker { image 'golang' } }
    stages {
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