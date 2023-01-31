pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                echo 'Building..'
                docker version
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
