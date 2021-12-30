pipeline {
    agent any
    stages {
        stage('Test') {
            sh 'go test'
        }

        stage('Build') {
            echo 'Building ${env.BUILD_ID} on ${env.JENKINS_URL}...'
            sh 'docker build -t ${env.JOB_NAME} .'
        }

        stage('Deploy') {
            environment {
                DEPLOY_SERVER_URL = 'projects.olliejonas.com'
            }
            echo 'Deploying ${env.BUILD_ID} onto ${DEPLOY_SERVER_URL}...'
        }
    }
}