package main

import (
    "fmt"
    "os"
    "time"

    mqtt "github.com/eclipse/paho.mqtt.golang"
)

var (
	MQTTBroker []string = []string{
		"tcp://test.mosquitto.org:1883",
		"tcp://broker.hivemq.com:1883",
		"tcp://broker.emqx.io:1883",
	}
)

func main() {
    broker := MQTTBroker[0]
    // 客户端ID
    clientID := "go_mqtt_client_example"
    // 订阅的主题
    topic := "test/topic"

    // 消息处理函数
    messagePubHandler := func(client mqtt.Client, msg mqtt.Message) {
        fmt.Printf("收到消息: [%s] %s\n", msg.Topic(), msg.Payload())
    }

    // 连接丢失处理函数
    connectLostHandler := func(client mqtt.Client, err error) {
        fmt.Printf("连接丢失: %v\n", err)
    }

    // 创建客户端选项
    opts := mqtt.NewClientOptions()
    opts.AddBroker(broker)
    opts.SetClientID(clientID)
    opts.SetDefaultPublishHandler(messagePubHandler)
    opts.OnConnectionLost = connectLostHandler

    // 创建并连接客户端
    client := mqtt.NewClient(opts)
    if token := client.Connect(); token.Wait() && token.Error() != nil {
        fmt.Println("连接失败:", token.Error())
        os.Exit(1)
    }
    fmt.Println("连接成功！")

    // 订阅主题
    if token := client.Subscribe(topic, 1, nil); token.Wait() && token.Error() != nil {
        fmt.Println("订阅失败:", token.Error())
        os.Exit(1)
    }
    fmt.Printf("已订阅主题: %s\n", topic)

    // 发布消息
    text := "Hello MQTT from Go!"
    token := client.Publish(topic, 0, false, text)
    token.Wait()
    fmt.Printf("已发布消息: %s\n", text)

    // 等待消息接收
    time.Sleep(5 * time.Second)

    // 断开连接
    client.Disconnect(250)
    fmt.Println("客户端已断开连接")
}
