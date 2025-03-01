pipeline {
    agent any

    parameters {
        choice(name: 'SERVICE', choices: ['crawl-service', 'article-service'], description: 'Select the service to build & test')
    }

    environment {
        GO_VERSION = "1.23" 
        WORKDIR = "${params.SERVICE}" 
    }

    stages {
        stage('Checkout Code') {
            steps {
                script {
                    echo "Cloning repository..."
                    checkout scm
                }
            }
        }

        stage('Setup Go Environment') {
            steps {
                script {
                    echo "Setting up Go environment..."
                    sh "go version"
                }
            }
        }

        stage('Download Dependencies') {
            steps {
                script {
                    echo "Downloading dependencies for ${WORKDIR}..."
                    sh "cd ${WORKDIR} && go mod tidy"
                }
            }
        }

        stage('Run Unit Tests') {
            steps {
                script {
                    echo "Running unit tests for ${WORKDIR}..."
                    sh "cd ${WORKDIR} && go test ./... -v"
                }
            }
        }
    }

    post {
        always {
            script {
                echo "Cleaning up workspace..."
                deleteDir()
            }
        }
        success {
            script {
                echo "✅ Unit tests passed for ${WORKDIR}!"
            }
        }
        failure {
            script {
                echo "❌ Unit tests failed for ${WORKDIR}!"
            }
        }
    }
}
