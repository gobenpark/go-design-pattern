package bridge

import (
	"fmt"
	"github.com/pkg/errors"
	"io"
)

/*
1. 콘솔에 print할 프린터 생성
2. io.Writer인터페이스를 사용해서 print할 프린터 생성

요구사항
1. 메시지를 받아들여 출력하는 `PrinterAPI`
2. 메시지를 콘솔에 출력하는 API 구현
3. 메시지를 io.Writer인터페이스에 출력하는 API 구현
4. printing type을 구현하는 print메소드를 사용하는 Printer 추상화
5. Printer 와 PrinterAPI인터페이스를 구현하는 normal 생성
6. normal프린터는 메시지를 구현체에 전달
7. Printer 추상화 및 PrinterAPI 인터페이스를 구현하는 Packt 프린터
8. Packt 프린터는 Message from Packt:를 프린트 할때마다 출력
*/

// 1 메시지를 받아들여 출력하는 `PrinterAPI`
type PrinterAPI interface {
	PrintMessage(string) error
}

// 2 메시지를 콘솔에 출력하는 API 구현
type PrinterImpl1 struct{}

func (p *PrinterImpl1) PrintMessage(msg string) error {
	fmt.Printf("%s\n", msg)
	return nil
}

// 3 메시지를 io.Writer인터페이스에 출력하는 API 구현
type PrinterImpl2 struct {
	Writer io.Writer
}

func (d *PrinterImpl2) PrintMessage(msg string) error {
	if d.Writer == nil {
		return errors.New("You need to pass an io.Writer to PrinterImpl2")
	}
	fmt.Fprintf(d.Writer, "%s", msg)
	return nil
}

type TestWriter struct {
	Msg string
}

func (t *TestWriter) Write(p []byte) (n int, err error) {
	n = len(p)
	if n > 0 {
		t.Msg = string(p)
		return n, nil
	}
	err = errors.New("Content received on Writer was empty")
	return
}

//4 printing type을 구현하는 print메소드를 사용하는 Printer 추상화
type PrinterAbstraction interface {
	Print() error
}

//5 Printer 와 PrinterAPI인터페이스를 구현하는 normal 생성
type NormalPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (c *NormalPrinter) Print() error {
	c.Printer.PrintMessage(c.Msg)
	return nil
}

// 7 Printer 추상화 및 PrinterAPI 인터페이스를 구현하는 Packt 프린터
type PacktPrinter struct {
	Msg     string
	Printer PrinterAPI
}

func (c *PacktPrinter) Print() error {
	c.Printer.PrintMessage(fmt.Sprintf("Message from Packt: %s", c.Msg))
	return nil
}
