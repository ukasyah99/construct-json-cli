package lib

import (
	"errors"
	"fmt"
	"strconv"

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

func Input(v *string, label string) error {
	prompt := promptui.Prompt{
		Label: label,
	}

	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("selectItem failed %v\n", err)
		return err
	}

	*v = result

	return nil
}

func InputNumber(v *int, label string) error {
	validate := func(input string) error {
		_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid number")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    label,
		Validate: validate,
	}

	result, err := prompt.Run()
	if err != nil {
		return err
	}

	num, err := strconv.Atoi(result)
	if err != nil {
		return err
	}

	*v = num

	return nil
}
