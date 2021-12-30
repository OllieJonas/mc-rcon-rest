pipeline {
    agent any { dockerfile = true }
    stages {
        stage('Test') {
            steps {
                sh 'go test'
            }
        }
    }
}