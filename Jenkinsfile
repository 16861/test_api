pipeline {
  agent any
  stages {
    stage('Build') {
      steps {
        git(url: 'https://github.com/16861/test_api', branch: 'master', poll: true)
      }
    }
  }
}