package main

import (
	"fmt"
	"time"
)

type Document struct {
	Head string		`json:"head"`
	Body string		`json:"body"`
	Footer string	`json:"footer"`
}

type docBuilder interface {
	buildDocumentHead(head string)
	buildDocumentBody(body string)
	buildDocumentFooter()
	composeDocument() Document
}

type defaultBuilder struct{
	header string
	body string
	footer string
}

func newDefaultBuilder() *defaultBuilder{
	return &defaultBuilder{}
}

func (b *defaultBuilder) buildDocumentHead(head string){
	b.header=fmt.Sprintf("Здоровки! Это тебе пишет %s",head)
}

func (b *defaultBuilder) buildDocumentBody(body string){
	b.body=fmt.Sprintf("Короче, тут такое дело: %s",body)
}

func (b *defaultBuilder) buildDocumentFooter(){
	b.footer=fmt.Sprintf("Все, жду ответа. Письмо написано:%s",time.Now().String())
}

func (b *defaultBuilder) composeDocument() Document{
	return Document{
		b.header,
		b.body,
		b.footer,
	}
}

type officialBuilder struct{
	header string
	body string
	footer string
}

func newOfficialBuilder() *officialBuilder{
	return &officialBuilder{}
}

func (b *officialBuilder) buildDocumentHead(head string){
	b.header=fmt.Sprintf("Добрый день. Вам пишет %s",head)
}

func (b *officialBuilder) buildDocumentBody(body string){
	b.body=fmt.Sprintf("У нас к вам есть важное дело: %s",body)
}

func (b *officialBuilder) buildDocumentFooter(){
	b.footer=fmt.Sprintf("Ждем вашего ответа. Письмо написано:%s",time.Now().String())
}

func (b *officialBuilder) composeDocument() Document{
	return Document{
		b.header,
		b.body,
		b.footer,
	}
}

func getDocBuilder(docType string)docBuilder{
	if docType=="default"{
		return &defaultBuilder{}
	}
	if docType=="official"{
		return &officialBuilder{}
	}
	return nil
}

const(
	OfficialDocument="official"
	InformalLetter="default"
)

func BuildLetter(head string, body string, docType string)Document{
	a:=getDocBuilder(docType)
	a.buildDocumentHead(head)
	a.buildDocumentBody(body)
	a.buildDocumentFooter()
	return a.composeDocument()
}
