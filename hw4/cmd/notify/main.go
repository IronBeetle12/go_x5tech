package main

import (
	"errors"
	"fmt"
	"go_x5tech/hw4/practice"
)

func main() {
	success()
	empty()
	missingSettings()
	brokenSettings()
	// TODO: домашняя работа. Покажите успешную рассылку, пустой текст,
	// ошибку настройки и продолжение после BrokenNotifier.
	// Импорт учебного пакета: "lesson5/student/practice".
	//fmt.Println("Добавьте сценарии из homework/README.md")
}

func success() {
	fmt.Println("Success")
	email := &practice.EmailNotifier{
		BaseNotifier: practice.BaseNotifier{Channel: "email"},
		Address:      "dep12@dozen.ru",
	}
	push := &practice.PushNotifier{
		BaseNotifier: practice.BaseNotifier{Channel: "push"},
		DeviceID:     "deded-123",
	}
	err := practice.Dispatch([]practice.Notifier{push, email}, "Заказ готов")
	if err != nil {
		fmt.Println("error: ", err)
	}
	fmt.Println("email: ", email.Messages)
	fmt.Println("push: ", push.Messages)
}

func empty() {
	fmt.Println("empty text")
	email := &practice.EmailNotifier{
		BaseNotifier: practice.BaseNotifier{Channel: "email"},
		Address:      "dep12@dozen.ru",
	}
	err := practice.Dispatch([]practice.Notifier{email}, "")
	if errors.Is(err, practice.ErrEmptyMessage) {
		fmt.Println("error: empty message")
	} else {
		fmt.Println("unknown error")
	}
	fmt.Println("email: ", email.Messages)
}

func missingSettings() {
	fmt.Println("missing settings")
	push := &practice.PushNotifier{
		BaseNotifier: practice.BaseNotifier{Channel: "push"},
	}
	err := practice.Dispatch([]practice.Notifier{push}, "Заказ готов")
	var sendErr *practice.SendError
	if errors.As(err, &sendErr) {
		fmt.Println("channel: ", sendErr.Channel)
		fmt.Println("reason: ", sendErr.Reason)
	} else {
		fmt.Println("unknown error")
	}
}

func brokenSettings() {
	fmt.Println("broken settings")
	broken := &practice.BrokenNotifier{
		BaseNotifier: practice.BaseNotifier{Channel: "broken"},
	}
	push := &practice.PushNotifier{
		BaseNotifier: practice.BaseNotifier{Channel: "push"},
		DeviceID:     "device-1",
	}
	errs := practice.SafeDispatch([]practice.Notifier{broken, push}, "Заказ готов")
	for _, err := range errs {
		if errors.Is(err, practice.ErrNotifierPanic) {
			fmt.Println("panic caught:", err)
		} else {
			fmt.Println("ошибка:", err)
		}
	}
	fmt.Println("push:", push.Messages)
}
