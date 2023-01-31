pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                docker build ./dockerfiles/dockerfile
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}
