package statemachine

import (
	"fmt"
	"study-statemachine/light"
)

type BlueLighting struct {
	mc *MachineContext
}

var (
	_ EventListener = (*BlueLighting)(nil)
	_ State         = (*BlueLighting)(nil)
)

func (s BlueLighting) Entry() {
	fmt.Println("BlueLighting#Entry")

	s.Do() // Doの処理は非同期実行だけど割愛
}

func (s BlueLighting) Do() {
	light.BlueLighting()
}

func (s BlueLighting) Exit() {
	// 本当はDoをキャンセルする処理が必用

	fmt.Println("BlueLighting#Exit")
}

func (l BlueLighting) OnStart() {
	l.mc.changeState(&RedLighting{mc: l.mc})
}

func (l BlueLighting) OnStop() {
	fmt.Println("[Blue] Power off")

	l.mc.terminate()
}
