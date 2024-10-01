package hiddify_extension

import (
	"fmt"
	"strconv"

	ui "github.com/hiddify/hiddify-core/extension/ui"
)

// Field name constants for easy reference, use similar name to the json key
const (
	CountKey      = "count"
	InputKey      = "input"
	PasswordKey   = "password"
	EmailKey      = "email"
	SelectKey     = "select"
	TextAreaKey   = "textarea"
	SwitchKey     = "switchVal"
	CheckboxKey   = "checkbox"
	RadioboxKey   = "radiobox"
	ContentKey    = "content"
	ConsoleKey    = "console"
	ButtonTestKey = "button_test"
)

// GetUI returns the UI form for the extension
func (e *HiddifyAppDemoExtension) GetUI() ui.Form {
	// Create a form depending on whether there is a background task or not
	if e.cancel != nil {
		return e.getRunningUI()
	}
	return e.getStoppedUI()
}

// setFormData validates and sets the form data from input
func (e *HiddifyAppDemoExtension) setFormData(data map[string]string) error {
	// Check if CountKey exists in the provided data
	if val, ok := data[CountKey]; ok {
		if intValue, err := strconv.Atoi(val); err == nil {
			// Validate that the count is greater than 5
			if intValue < 5 {
				return fmt.Errorf("please use a number greater than 5")
			} else {
				e.Base.Data.Count = intValue // Set valid count value
			}
		} else {
			return err // Return parsing error
		}
	}
	if val, ok := data[InputKey]; ok {
		e.Base.Data.Input = val
	}
	if val, ok := data[PasswordKey]; ok {
		e.Base.Data.Password = val
	}
	if val, ok := data[EmailKey]; ok {
		e.Base.Data.Email = val
	}
	if val, ok := data[SelectKey]; ok {
		if selectedValue, err := strconv.ParseBool(val); err == nil {
			e.Base.Data.Selected = selectedValue
		} else {
			return err
		}
	}
	if val, ok := data[TextAreaKey]; ok {
		e.Base.Data.Textarea = val
	}
	if val, ok := data[SwitchKey]; ok {
		if selectedValue, err := strconv.ParseBool(val); err == nil {
			e.Base.Data.SwitchVal = selectedValue
		} else {
			return err
		}
	}
	// if val, ok := data[CheckboxKey]; ok {
	// 	e.checkbox = val
	// }
	if val, ok := data[ContentKey]; ok {
		e.Base.Data.Content = val
	}
	if val, ok := data[RadioboxKey]; ok {
		e.Base.Data.Radiobox = val
	}

	return nil // Return nil if data is set successfully
}

func (e *HiddifyAppDemoExtension) getRunningUI() ui.Form {
	return ui.Form{
		Title:       "hiddify-app-demo-extension",
		Description: "Awesome Extension hiddify_app_demo_extension created by hiddify",

		Fields: [][]ui.FormField{
			{{
				Type:  ui.FieldConsole,
				Key:   ConsoleKey,
				Label: "Console",
				Value: e.console, // Display console output
				Lines: 20,
			}},
			{{
				Type:  ui.FieldButton,
				Key:   ui.ButtonCancel,
				Label: "Cancel",
			}},
		},
	}
}

func (e *HiddifyAppDemoExtension) getStoppedUI() ui.Form {
	return ui.Form{
		Title:       "hiddify-app-demo-extension",
		Description: "Awesome Extension hiddify_app_demo_extension created by hiddify",

		Fields: [][]ui.FormField{
			{{
				Type:        ui.FieldInput,
				Key:         CountKey,
				Label:       "Count",
				Placeholder: "This will be the count",
				Required:    true,
				Value:       fmt.Sprintf("%d", e.Base.Data.Count),
				Validator:   ui.ValidatorDigitsOnly,
			}},
			{{
				Type:        ui.FieldInput,
				Key:         InputKey,
				Label:       "Hi Group",
				Placeholder: "Hi Group flutter",
				Required:    true,
				Value:       e.Base.Data.Input,
			}},
			{{
				Type:  ui.FieldSwitch,
				Key:   SelectKey,
				Label: "Select Label",
				Value: strconv.FormatBool(e.Base.Data.Selected),
			}},
			{{
				Type:        ui.FieldTextArea,
				Key:         TextAreaKey,
				Label:       "TextArea Label",
				Placeholder: "Enter your text",
				Required:    true,
				Value:       e.Base.Data.Textarea,
			}},
			{{
				Type:  ui.FieldSwitch,
				Key:   SwitchKey,
				Label: "Switch Label",
				Value: strconv.FormatBool(e.Base.Data.SwitchVal),
			}},
			// {
			// 	Type:     ui.Checkbox,
			// 	Key:      CheckboxKey,
			// 	Label:    "Checkbox Label",
			// 	Required: true,
			// 	Value:    e.checkbox,
			// 	Items: []ui.SelectItem{
			// 		{
			// 			Label: "A",
			// 			Value: "A",
			// 		},
			// 		{
			// 			Label: "B",
			// 			Value: "B",
			// 		},
			// 	},
			// },
			{{
				Type:     ui.FieldRadioButton,
				Key:      RadioboxKey,
				Label:    "Radio Label",
				Required: true,
				Value:    e.Base.Data.Radiobox,
				Items: []ui.SelectItem{
					{
						Label: "A",
						Value: "A",
					},
					{
						Label: "B",
						Value: "B",
					},
				},
			}},
			{{
				Type:     ui.FieldTextArea,
				Readonly: true,
				Key:      ContentKey,
				Label:    "Content",
				Value:    e.Base.Data.Content,
				Lines:    10,
			}},
			{
				{
					Type:  ui.FieldButton,
					Key:   ButtonTestKey,
					Label: "Test",
				},
				{
					Type:  ui.FieldButton,
					Key:   ui.ButtonSubmit,
					Label: "Submit",
				},
			},
		},
	}
}
