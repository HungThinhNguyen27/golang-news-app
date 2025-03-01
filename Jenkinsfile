pipeline {
    agent any

    tools {
        go '1.24.0'
    }

    parameters {
        choice(name: 'SERVICE', choices: ['crawl-service', 'article-service'], description: 'Select the service to build & test')
    }

    environment {
        WORKDIR = "${WORKSPACE}/${params.SERVICE}"
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

        stage('Check Go Version') {
            steps {
                sh 'go version'
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
                echo "Unit tests passed for ${WORKDIR}!"
            }
        }
        failure {
            script {
                echo "Unit tests failed for ${WORKDIR}!"
            }
        }
    }
}
