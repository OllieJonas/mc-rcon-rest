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
                echo "Building ${env.BUILD_TAG} on ${env.JENKINS_URL} ..."
                sh "docker build -t ${env.BUILD_TAG} ."
            }
        }

        stage('Deploy') {
            environment {
                DEPLOY_SERVER_URL = "projects.olliejonas.com"
                DEPLOY_SERVER_USER = "root"
            }
            steps {
                echo "Deploying ${env.BUILD_TAG} onto ${DEPLOY_SERVER_URL} ..."
                sh "docker save -o ${env.BUILD_TAG}.tar ${env.BUILD_TAG}:latest | gzip > ${env.BUILD_TAG}.tar.gz"

                sshagent(credentials: ['projects']) {
                    sh '''
                        [ -d ~/.ssh ] || mkdir ~/.ssh && chmod 0700 ~/.ssh
                        ssh-keyscan -t rsa,dsa ${DEPLOY_SERVER_URL} >> ~/.ssh/known_hosts
                        cat << 'EOT' | ssh ${DEPLOY_SERVER_USER}@${DEPLOY_SERVER_URL} /bin/bash
                        echo whoami
                        ls -l
                        EOT
                    '''
                }
            }
        }

        stage('Cleanup') {
            steps {
                echo "Performing cleanup..."
                sh "docker image prune --force" // in case it created any dangling images
                sh "docker image rm ${env.BUILD_TAG} --force" // dont want the image left on the agent

                // remove any zipped stuff
                sh "rm ${env.BUILD_TAG}.tar"
                sh "rm ${env.BUILD_TAG}.tar.gz"
            }
        }
    }
}