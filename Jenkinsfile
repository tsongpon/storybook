pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('pull dependencies') {
            steps {
                sh 'curl https://raw.githubusercontent.com/golang/dep/master/install.sh | sh'
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