# GPT-asker
## English
> personal GPT language model asker script using command line
### Quick Start
Firstly you need to add your own Openai-api-key in the file `application.yaml`, make sure you have installed go1.20 or higher version, then enter the following command in your terminal
```bash
go build gpt-asker
./gpt-asker
```
then you can chat with GPT through local command line
### Environment Test
```bash
git checkout hello
go run hello.go
```
this script is an environment test script, it will send a message goes like "Hello! I'm using go API to chat with you" to GPT, you can check whether your system environment is successfully running with this project
## 中文
> 个人GPT命令行提问脚本
### 快速开始
首先你需要在`application.yaml`中添加自己的Openai-api-key，并确保安装了go1.20版本，然后在命令行终端中输入以下命令
```bash
go build gpt-asker
./gpt-asker
```
 接下来就可以通过命令行与gpt聊天了
### 环境测试
```bash
git checkout hello
go run hello.go
```
该脚本是环境测试脚本，固定向gpt发送一条"Hello! I'm using go API to chat with you"的消息，你可以通过该脚本的运行情况判断你的运行环境是否正常