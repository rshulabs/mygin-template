pipeline {
    agent any

    stages {
        stage('拉取仓库') {
            steps {
                git branch: 'main', credentialsId: 'git', url: 'https://github.com/rshulabs/mygin-template.git'
            }
        }
        stage('bulid code') {
            steps {
                sh '''echo "开始构建"
                    echo "构建完成"'''
            }
        }
        stage('upload') {
            steps {
                sshPublisher(publishers: [sshPublisherDesc(configName: '192.168.60.99', transfers: [sshTransfer(cleanRemote: false, excludes: '', execCommand: 'echo "pipeline"', execTimeout: 120000, flatten: false, makeEmptyDirs: false, noDefaultExcludes: false, patternSeparator: '[, ]+', remoteDirectory: 'pipeline-test/', remoteDirectorySDF: false, removePrefix: '', sourceFiles: '**')], usePromotionTimestamp: false, useWorkspaceInPromotion: false, verbose: false)])
            }
        }
    }
}
