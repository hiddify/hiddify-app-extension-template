package hiddify_extension

import (
	"fmt"
	"strconv"

	ui "github.com/hiddify/hiddify-core/extension/ui"
)

const (
	CountKey      = "count"
	ConsoleKey    = "console"
	ButtonTestKey = "button_test"
)

// GetUI returns the UI form for the extension
func (e *ExampleExtension) GetUI() ui.Form {
	// Create a form depending on whether there is a background task or not
	if e.cancel != nil {
		return e.getRunningUI()
	}
	return e.getStoppedUI()
}

// setFormData validates and sets the form data from input
func (e *ExampleExtension) setFormData(data map[string]string) error {
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
	return nil // Return nil if data is set successfully
}

func (e *ExampleExtension) getRunningUI() ui.Form {

	return ui.Form{
		Title:       "project_urlname",
		Description: "project_description",
		Fields: [][]ui.FormField{
			{{
				Type:  ui.FieldConsole,
				Key:   ConsoleKey,
				Label: "Console",
				Value: e.console, // Display console output
				Lines: 20,
			}},
			{{Type: ui.FieldButton, Key: ui.ButtonCancel, Label: "Cancel"}},
		},
	}
}
func (e *ExampleExtension) getStoppedUI() ui.Form {
	// Inital page
	return ui.Form{
		Title:       "project_urlname",
		Description: "project_description",
		Fields: [][]ui.FormField{
			{{
				Type:        ui.FieldInput,
				Key:         CountKey,
				Label:       "Count",
				Placeholder: "This will be the count",
				Required:    true,
				Value:       fmt.Sprintf("%d", e.Base.Data.Count), // Default value from stored data
				Validator:   ui.ValidatorDigitsOnly,               // Only allow digits
			}},
			{{
				Type:  ui.FieldConsole,
				Key:   ConsoleKey,
				Label: "Console",
				Value: e.console, // Display current console output
				Lines: 20,
			}},
			{
				{Type: ui.FieldButton, Key: ButtonTestKey, Label: "Test"},
				{Type: ui.FieldButton, Key: ui.ButtonSubmit, Label: "Submit"},
			},
		},
	}
}

// addAndUpdateConsole adds messages to the console and updates the UI
func (e *ExampleExtension) addAndUpdateConsole(message ...any) {
	e.console = fmt.Sprintln(message...) + e.console // Prepend new messages to the console output
	e.UpdateUI(e.GetUI())                            // Update the UI with the new console content
}
