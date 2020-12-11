### wechat-bot-go

wechat bot golang client

### **接收**消息类型

- 群文本消息

```json
{
  "content":"高速老是在修路",
  "id":"20201211160304",
  "receiver":"430122245@chatroom", 
  "sender":"wxid_t0rcxgk3o01w12",
  "srvid":1,
  "time":"2020-12-11 16:03:04",
  "type":1
}
```

注意到 receiver 为群聊 roomid, sender 为发送的人

- 个人文本消息

```json
{
  "content":"？",
  "id":"20201211160346",
  "receiver":"self",
  "sender":"wxid_uspyqfp09fb621",
  "srvid":1,
  "time":"2020-12-11 16:03:46",
  "type":1
}
```

注意到 receiver 为 self，也就是说接收个人消息的话可以通过这个 key 来进行判断

### acknowledgement

- [wechat-bot](https://github.com/cixingguangming55555/wechat-bot)