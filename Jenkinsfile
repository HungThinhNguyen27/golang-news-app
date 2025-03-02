pipeline {
    agent any

    tools {
        go '1.24.0'
    }
    
    environment {
        CHANGED_SERVICES = ''
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

        stage('Detect Changed Services') {
            steps {
                script {
                    echo "Checking for changed services..."
                    def changedFiles = sh(script: 'git diff --name-only HEAD~1', returnStdout: true).trim().split('\n')
                    
                    def services = ['crawl-service', 'article-service']
                    def changedServices = []

                    for (service in services) {
                        if (changedFiles.any { it.startsWith(service + "/") }) {
                            changedServices.add(service)
                        }
                    }

                    CHANGED_SERVICES = changedServices.join(' ')
                    if (CHANGED_SERVICES == '') {
                        echo "No services changed. Skipping build."
                        currentBuild.result = 'SUCCESS'
                        return
                    }
                    
                    echo "Services to build: ${CHANGED_SERVICES}"
                }
            }
        }

        stage('Build & Test Changed Services') {
            when {
                expression { return CHANGED_SERVICES != '' }
            }
            steps {
                script {
                    for (service in CHANGED_SERVICES.split(' ')) {
                        def workdir = "${WORKSPACE}/${service}"
                        
                        echo "Building & Testing ${service}..."

                        sh """
                        cd ${workdir}
                        go mod tidy
                        go test ./... -v
                        """
                    }
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
    }
}
