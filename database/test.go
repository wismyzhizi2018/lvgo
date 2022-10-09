package main

import (
	"errors"
	"fmt"
)

//策略模式
//比如：在你发送好友消息的时候
//
//解决在多重相似算法情况下使用了 if…else 和 switch…case 带来的复杂性和臃肿性问题。
//
//策略模式适用于以下场景：
//
//针对同一类型问题，有多种处理方式，每一种都能独立解决问题。
//
//需要自由切换选择不同类型算法场景。
//
//需要闭屏算法具体实现规则场景。
//
//一般要实现一个较为完整的策略模式，需要如下组成单元：
//
//上下文控制函数：用来拒绝切换不同算法，屏蔽高层模块（调用者）对策略、算法的直接访问，封装可能存在的变化–可以用简单工程与单例模式封装该函数
//
//抽象要实现的策略接口：定义一个 interface，决定好内部包含的具体函数方法定义。
//
//具体的策略角色：实现每个类实现抽象接口的方法，进行内部具体算法的维护实现即可。

type RongyunSendMessage interface {
	Send(fromUser string, toUser string, Content string) error
	DownloadOrder()
	FormatOrder()
	CheckOrderSign()
	InsertOrder()
	ParseOrder()
}
type RongyunSendImgMessage struct{}

func (r RongyunSendImgMessage) DownloadOrder() {
	//TODO implement me
	panic("implement me")
}

func (r RongyunSendImgMessage) FormatOrder() {
	//TODO implement me
	panic("implement me")
}

func (r RongyunSendImgMessage) CheckOrderSign() {
	//TODO implement me
	panic("implement me")
}

func (r RongyunSendImgMessage) InsertOrder() {
	//TODO implement me
	panic("implement me")
}

func (r RongyunSendImgMessage) ParseOrder() {
	//TODO implement me
	panic("implement me")
}

func (r RongyunSendImgMessage) Send(fromUser string, toUser string, Content string) error {
	// 判断用户状态
	// 判断用户权限
	// 整理发送数据
	// 发送消息
	// 消息记录入库
	fmt.Printf("fromUser [%v] => toUser [%s] : content [%v]\n", fromUser, toUser, Content)
	return nil
}

type RongyunSendImgsMessage struct{}

func (r RongyunSendImgsMessage) DownloadOrder() {
	//TODO implement me
	panic("implement me")
}

func (r RongyunSendImgsMessage) FormatOrder() {
	//TODO implement me
	panic("implement me")
}

func (r RongyunSendImgsMessage) CheckOrderSign() {
	//TODO implement me
	panic("implement me")
}

func (r RongyunSendImgsMessage) InsertOrder() {
	//TODO implement me
	panic("implement me")
}

func (r RongyunSendImgsMessage) ParseOrder() {
	//TODO implement me
	panic("implement me")
}

func (r RongyunSendImgsMessage) Send(fromUser string, toUser string, Content string) error {
	// 判断用户状态
	// 判断用户权限
	// 整理发送数据
	// 发送消息
	// 消息记录入库
	fmt.Printf("fromUser [%v] => toUser [%s] : content [%v]\n", fromUser, toUser, Content)
	return nil
}

type RongyunSendVideoMessage struct{}

func (r RongyunSendVideoMessage) DownloadOrder() {
	//TODO implement me
	panic("implement me")
}

func (r RongyunSendVideoMessage) FormatOrder() {
	//TODO implement me
	panic("implement me")
}

func (r RongyunSendVideoMessage) CheckOrderSign() {
	//TODO implement me
	panic("implement me")
}

func (r RongyunSendVideoMessage) InsertOrder() {
	//TODO implement me
	panic("implement me")
}

func (r RongyunSendVideoMessage) ParseOrder() {
	//TODO implement me
	panic("implement me")
}

func (r RongyunSendVideoMessage) Send(fromUser string, toUser string, Content string) error {
	// 判断用户状态
	// 判断用户权限
	// 整理发送数据
	// 发送消息
	// 消息记录入库
	fmt.Printf("fromUser [%v] => toUser [%s] : content [%v]\n", fromUser, toUser, Content)
	return nil
}

type MessageParams struct {
	Type     string // 消息类型：img=单图，imgs=多图，video=视频 ...
	Content  string // 消息内容
	FromUser string // 发送方
	ToUser   string // 接收方
}

func main() {
	// 发送单图
	sendImg := MessageParams{
		Type:     "img",
		Content:  "发送单图消息",
		FromUser: "A",
		ToUser:   "B",
	}
	// 发送多图
	sendImgs := MessageParams{
		Type:     "imgs",
		Content:  "发送多图",
		FromUser: "A",
		ToUser:   "B",
	}
	// 发送视频
	sendVideo := MessageParams{
		Type:     "video",
		Content:  "发送视频",
		FromUser: "A",
		ToUser:   "B",
	}
	SendMessage(sendImg)   // 单图
	SendMessage(sendImgs)  // 多图
	SendMessage(sendVideo) // 视频
}

var Template = map[string]RongyunSendMessage{
	"img":   new(RongyunSendImgMessage),
	"imgs":  new(RongyunSendImgsMessage),
	"video": new(RongyunSendVideoMessage),
}

func SendMessage(params MessageParams) error {
	if _, ok := Template[params.Type]; !ok {
		return errors.New("tagID invalid")
	}
	return Template[params.Type].Send(params.FromUser, params.ToUser, params.Content)
}
