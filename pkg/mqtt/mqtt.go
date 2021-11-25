package mqtt

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	pmqtt "github.com/eclipse/paho.mqtt.golang"
)

type MqttInfo struct {
	username string
	password string
	port     int
	ip       string
	scadaId  string
	topic    string
}

type DataFormate struct {
	Command string      `json:"Cmd"`
	Value   interface{} `json:"Val"`
}

type DeskMqttFormat struct {
	D  DataFormate `json:"d"`
	Ts string      `json:"ts"`
}

var (
	Info   MqttInfo
	Client pmqtt.Client
)

func SetMQueue(username string, password string, ip string, scadaId string, mqttTopic string, deviceId string) {

	Info = MqttInfo{
		username: username, //"Goy2waYPAGQP:7LQDMXZ1p2eY"
		password: password, //"flDWSFGo037RvoQwhOTu"
		port:     1883,
		ip:       ip,      //"rabbitmq-001-pub.hz.wise-paas.com.cn"
		scadaId:  scadaId, //"scada_Hdym5d86YUdb",
	}
	Info.topic = func() string {
		mqttTopic := mqttTopic //iot-2/evt/+/fmt/scada_Hdym5d86YUdb/#
		// fmt.Println("created topic:", mqttTopic)
		mqttTopic = strings.Replace(mqttTopic, "+", "wacmd", 1)
		mqttTopic = strings.Replace(mqttTopic, "#", deviceId, 1)
		return mqttTopic
	}()

	// fmt.Printf("%+v\n", mqttInfo)
	// fmt.Println("topic:", mqttInfo.topic)

	Client = connectMqtt()

}

var messagePubHandler pmqtt.MessageHandler = func(client pmqtt.Client, msg pmqtt.Message) {
	// fmt.Printf("Received message: %s from topic: %s\n", msg.Payload(), msg.Topic())
}

var connectHandler pmqtt.OnConnectHandler = func(client pmqtt.Client) {
	// log.Println("mqtt connected...")
}

var connectLostHandler pmqtt.ConnectionLostHandler = func(client pmqtt.Client, err error) {
	// fmt.Printf("Connect lost: %v", err)
}

func connectMqtt() pmqtt.Client {

	var broker = Info.ip
	var port = Info.port
	opts := pmqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID("go_mqtt_client")
	opts.SetUsername(Info.username)
	opts.SetPassword(Info.password)

	opts.SetDefaultPublishHandler(messagePubHandler)
	opts.OnConnect = connectHandler
	opts.OnConnectionLost = connectLostHandler
	client := pmqtt.NewClient(opts)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// client.Disconnect(250)
	return client

}

func PublishMQueue(flatJson interface{}) {

	deskMqttFormat := DeskMqttFormat{}

	deskMqttFormat.D.Command = "WV"
	deskMqttFormat.D.Value = flatJson
	deskMqttFormat.Ts = time.Now().Format("2006-01-02T15:04:05.000Z")

	b, err := json.MarshalIndent(deskMqttFormat, "", " ")
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(Green("publish mqtt payload:"), Green(string(b)))

	token := Client.Publish(Info.topic, 0, false, string(b))
	token.Wait()
	time.Sleep(time.Second)

}
