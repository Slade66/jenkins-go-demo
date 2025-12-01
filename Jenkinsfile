pipeline {
    agent any
    environment {
        IMAGE_NAME = "my-go-app"
        CONTAINER_NAME = "my-go-service"
        PORT = "8091"
    }
    stages {
        stage('CI: Build & Test') {
            steps {
                script {
                    echo '从 Git 拉取代码完成，开始构建...'
                    // 启动 Go 容器
                    docker.image('golang:1.20-alpine').inside {
                        sh 'go version'
                        // 如果 go.mod 不存在则初始化
                        sh '[ -f go.mod ] || go mod init myapp'
                        sh 'go test -v ./...'
                        sh 'go build -o myapp main.go'
                    }
                }
            }
        }
        stage('CD: Deploy') {
            steps {
                script {
                    echo '构建 Docker 镜像...'
                    sh "docker build -t ${IMAGE_NAME} ."

                    echo '部署中...'
                    sh "docker rm -f ${CONTAINER_NAME} || true"
                    sh "docker run -d -p ${PORT}:8090 --name ${CONTAINER_NAME} ${IMAGE_NAME}"
                }
            }
        }
    }
}