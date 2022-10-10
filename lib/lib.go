package lib

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func SelectItem(v *string, label string, items []string) error {
	prompt := promptui.Select{
		Label: label,
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		fmt.Printf("selectItem failed %v\n", err)
		return err
	}

	*v = result

	return nil
}
