package statemachine

import (
	"fmt"
	"study-statemachine/light"
)

type RedLighting struct {
	mc *MachineContext
}

var (
	_ EventListener = (*RedLighting)(nil)
	_ State         = (*RedLighting)(nil)
)

func (s RedLighting) Entry() {
	fmt.Println("RedLighting#Entry")

	s.Do() // Doの処理は非同期実行だけど割愛
}

func (s RedLighting) Do() {
	light.RedLighting()
}

func (s RedLighting) Exit() {
	// 本当はDoをキャンセルする処理が必用

	fmt.Println("RedLighting#Exit")
}

func (l RedLighting) OnStart() {
	l.mc.changeState(&BlueLighting{mc: l.mc})
}

func (l RedLighting) OnStop() {
	fmt.Println("[Red] Power off")

	l.mc.terminate()
}
