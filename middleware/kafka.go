package middleware

/**
1、安装部署
docker pull wurstmeister/zookeeper
docker pull wurstmeister/kafka
docker run -d --name zookeeper -p2181:2181 wurstmeister/zookeeper
docker run -d --name kafka -p 9092:9092 -e KAFKA_BROKER_ID=0 -e KAFKA_ZOOKEEPER_CONNECT=docker.for.mac.host.internal:2181/kafka -e KAFKA_ADVERTISED_LISTENERS=PLAINTEXT://docker.for.mac.host.internal:9092 -e KAFKA_LISTENERS=PLAINTEXT://0.0.0.0:9092 wurstmeister/kafka

终端1发布：
docker exec -it kafka bash
cd /opt/kafka_xxx/bin
./kafka-console-producer.sh --broker-list localhost:9092 --topic first-topic

终端2消费：
docker exec -it kafka bash
cd /opt/kafka_xxx/bin
./kafka-console-consumer.sh --bootstrap-server localhost:9092 --topic first-topic --from-beginning
*/

/**
docker补充知识：
docker exec 在运行的容器中执行命令，-i 即使没有附加也保持STDIN打开，-t分配一个伪终端
docker kill/stop docker stop：先发sigterm信号给docker，运行其在一定时间内进行一些操作后再关闭。 docker kill：直接发送sigkill信号杀死容器。
docker run/start docker run：从镜像启动一个容器。 docker start：运行已停止的容器。
-d设置后台运行
--name zookeeper指定容器别名
-p 2181:2181绑定容器端口到宿主端口
*/

/**
2、简介
领英团队开发
基于java和scala开发， 设计中大量使用批处理和异步思想
相比其他消息中间件优势：
- 极致性能：最高每秒千万级别
- 生态系统兼容性太牛逼：Kafka 与周边生态系统的兼容性是最好的没有之一，尤其在大数据和流计算领域。

发布--订阅模型 兼容 队列模型

概念：
Topic（主题） : Producer 将消息发送到特定的主题，Consumer 通过订阅特定的 Topic(主题) 来消费消息
Partition（分区） : Partition 属于 Topic 的一部分。一个 Topic 可以有多个 Partition ，并且同一 Topic 下的 Partition 可以分布在不同的 Broker 上，这也就表明一个 Topic 可以横跨多个 Broker

Kafka 通过给特定 Topic 指定多个 Partition, 而各个 Partition 可以分布在不同的 Broker 上, 这样便能提供比较好的并发能力（负载均衡）。
Partition 可以指定对应的 Replica 数, 这也极大地提高了消息存储的安全性, 提高了容灾能力，不过也相应的增加了所需要的存储空间。

ZooKeeper 主要为 Kafka 提供元数据的管理的功能。
- Broker 注册
- Topic 注册
- 负载均衡
*/

/**
3、消息发布相关

发布之前 消息可能会经过 拦截器、序列化器、分区器
批处理思想：
- 发布缓存区 RecordAccumulator ，默认32M。 sender是一个dan
*/
