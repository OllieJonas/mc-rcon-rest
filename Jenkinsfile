pipeline {
    agent any
    stages {
        stage('Test') {
            steps {
                echo "Running tests for ${env.BUILD_ID} on ${env.JENKINS_URL} ..."
                sh 'go test'
            }
        }

        stage('Build') {
            steps {
                echo "Building ${env.BUILD_ID} on ${env.JENKINS_URL} ..."
                sh "docker build -t ${env.JOB_NAME} ."
            }
        }

        stage('Deploy') {
            environment {
                DEPLOY_SERVER_URL = 'projects.olliejonas.com'
                DEPLOY_SERVER_USER = 'root'
            }
            steps {
                echo "Deploying ${env.BUILD_ID} onto ${DEPLOY_SERVER_URL} ..."
            }
        }

        stage('Cleanup') {
            steps {
                echo "Performing cleanup..."
                sh "docker image prune --force"
                // sh "docker image rm {env.JENKINS_URL} --force"
            }
        }
    }
}