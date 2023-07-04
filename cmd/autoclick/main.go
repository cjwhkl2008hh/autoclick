package main

import (
	"fmt"
	"time"

	g "github.com/AllenDang/giu"
	r "github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
)

const hint = "1.鼠标移动到上键点击F3\n2.鼠标移动到右键点击F4\nF7开始第9关\nF8开始第10关\n\n"

var text = hint
var current_x = 0
var current_y = 0
var top_x = 0
var top_y = 0
var right_x = 0
var right_y = 0
var running = 0

var level9 = "uurrurrdllllllulddrdrruurrrrddlurulllllrddlluluurddurrrrrddlurullllluldrrruu"
var level10 = "ddrluurrrrurddduuulldldduuldrururrdddlllulldrrrrrurrddluruldlllluurdldrrlluu"

func loop() {
	text = hint
	if top_x == 0 {
		text += "上键：未映射\n"
	} else {
		text += fmt.Sprintf("上键：X=%d,Y=%d\n", top_x, top_y)
	}
	if right_x == 0 {
		text += "右键：未映射\n"
	} else {
		text += fmt.Sprintf("右键：X=%d,Y=%d\n", right_x, right_y)
	}
	if running == 0 {
		text += "正在运行：无\n"
	} else {
		text += fmt.Sprintf("正在运行：第%d关\n", running)
	}
	g.SingleWindow().Layout(
		g.Label(text),
	)
}

func execute(instructs string) {
	left_x := 2*top_x - right_x
	left_y := right_y
	down_x := top_x
	down_y := 2*right_y - top_y
	for _, instruct := range instructs {
		switch instruct {
		case 'u':
			r.Move(top_x, top_y)
		case 'l':
			r.Move(left_x, left_y)
		case 'd':
			r.Move(down_x, down_y)
		case 'r':
			r.Move(right_x, right_y)
		case 's':
			time.Sleep(time.Second)
			continue
		}
		r.Toggle("left")
		time.Sleep(time.Millisecond * time.Duration(50))
		r.Toggle("left", "up")
		time.Sleep(time.Millisecond * time.Duration(300))
		if running == 0 {
			return
		}
	}
	running = 0
	g.Update()
}

func process() {
	events := hook.Start()
	for e := range events {
		if e.Kind == hook.MouseMove {
			current_x = (int)(e.X)
			current_y = (int)(e.Y)
		} else if e.Kind == hook.KeyUp {
			switch e.Keycode {
			case hook.Keycode["f3"]:
				top_x = current_x
				top_y = current_y
			case hook.Keycode["f4"]:
				right_x = current_x
				right_y = current_y
			case hook.Keycode["f7"]:
				if top_x > 0 && right_x > 0 && running == 0 {
					running = 9
					go execute(level9)
				}
			case hook.Keycode["f8"]:
				if top_x > 0 && right_x > 0 && running == 0 {
					running = 10
					go execute(level10)
				}
			case hook.Keycode["esc"]:
				running = 0
			}
			g.Update()
		}
	}
}

func main() {
	wnd := g.NewMasterWindow("自动点击器", 200, 200, g.MasterWindowFlagsNotResizable)
	go process()
	wnd.Run(loop)
}
