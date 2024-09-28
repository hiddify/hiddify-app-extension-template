package hiddify_extension

import (
	"context"
	"fmt"
	"strconv"
	"time"

	ex "github.com/hiddify/hiddify-core/extension"
	ui "github.com/hiddify/hiddify-core/extension/ui"
)

// Field name constants
const (
	CountKey    = "countKey"
	InputKey    = "inputKey"
	PasswordKey = "passwordKey"
	EmailKey    = "emailKey"
	SelectKey   = "selectKey"
	TextAreaKey = "textareaKey"
	SwitchKey   = "switchKey"
	CheckboxKey = "checkboxKey"
	RadioboxKey = "radioboxKey"
	ContentKey  = "contentKey"
)

type ExampleExtensionConfig struct {
	Count     int    `json:"count"`
	Input     string `json:"input"`
	Password  string `json:"password"`
	Email     string `json:"email"`
	Selected  bool   `json:"selected"`
	Textarea  string `json:"textarea"`
	SwitchVal bool   `json:"switchVal"`
	// checkbox  string
	Radiobox string `json:"radiobox"`
	Content  string `json:"content"`
}
type ExampleExtension struct {
	ex.Base[ExampleExtensionConfig]
	cancel context.CancelFunc
}

func (e *ExampleExtension) GetId() string {
	return "example"
}

func (e *ExampleExtension) GetTitle() string {
	return "Example Extension"
}

func (e *ExampleExtension) GetDescription() string {
	return "This is a sample extension."
}

func (e *ExampleExtension) GetUI() ui.Form {
	// e.setFormData(data)
	return e.buildForm()
}

func (e *ExampleExtension) setFormData(data map[string]string) error {
	if val, ok := data[CountKey]; ok {
		if intValue, err := strconv.Atoi(val); err == nil {
			e.Base.Data.Count = intValue
		} else {
			return err
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
	return nil
}

func (e *ExampleExtension) buildForm() ui.Form {
	return ui.Form{
		Title:       "Example Form",
		Description: "This is a sample form.",
		Buttons:     []string{ui.Button_Submit, ui.Button_Cancel},
		Fields: []ui.FormField{
			{
				Type:        ui.FieldInput,
				Key:         CountKey,
				Label:       "Count",
				Placeholder: "This will be the count",
				Required:    true,
				Value:       fmt.Sprintf("%d", e.Base.Data.Count),
				Validator:   ui.ValidatorDigitsOnly,
			},
			{
				Type:        ui.FieldInput,
				Key:         InputKey,
				Label:       "Hi Group",
				Placeholder: "Hi Group flutter",
				Required:    true,
				Value:       e.Base.Data.Input,
			},
			{
				Type:     ui.FieldPassword,
				Key:      PasswordKey,
				Label:    "Password",
				Required: true,
				Value:    e.Base.Data.Password,
			},
			{
				Type:        ui.FieldEmail,
				Key:         EmailKey,
				Label:       "Email Label",
				Placeholder: "Enter your email",
				Required:    true,
				Value:       e.Base.Data.Email,
			},
			{
				Type:  ui.FieldSwitch,
				Key:   SelectKey,
				Label: "Select Label",
				Value: strconv.FormatBool(e.Base.Data.Selected),
			},
			{
				Type:        ui.FieldTextArea,
				Key:         TextAreaKey,
				Label:       "TextArea Label",
				Placeholder: "Enter your text",
				Required:    true,
				Value:       e.Base.Data.Textarea,
			},
			{
				Type:  ui.FieldSwitch,
				Key:   SwitchKey,
				Label: "Switch Label",
				Value: strconv.FormatBool(e.Base.Data.SwitchVal),
			},
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
			{
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
			},
			{
				Type:     ui.FieldTextArea,
				Readonly: true,
				Key:      ContentKey,
				Label:    "Content",
				Value:    e.Base.Data.Content,
				Lines:    10,
			},
		},
	}
}

func (e *ExampleExtension) backgroundTask(ctx context.Context) {
	i := 1
	for {
		select {
		case <-ctx.Done():
			e.Base.Data.Content = strconv.Itoa(i) + " Background task stop...\n" + e.Base.Data.Content
			e.UpdateUI(e.buildForm())

			fmt.Println("Background task stopped")
			return
		case <-time.After(1000 * time.Millisecond):
			txt := strconv.Itoa(i) + " Background task working..."
			e.Base.Data.Content = txt + "\n" + e.Base.Data.Content
			e.UpdateUI(e.buildForm())
			fmt.Println(txt)
		}
		i++
	}
}

func (e *ExampleExtension) SubmitData(data map[string]string) error {
	err := e.setFormData(data)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	e.cancel = cancel

	go e.backgroundTask(ctx)

	return nil
}

func (e *ExampleExtension) Cancel() error {
	if e.cancel != nil {
		e.cancel()
		e.cancel = nil
	}
	return nil
}

func (e *ExampleExtension) Stop() error {
	if e.cancel != nil {
		e.cancel()
		e.cancel = nil
	}
	return nil
}

func NewExampleExtension() ex.Extension {
	return &ExampleExtension{
		Base: ex.Base[ExampleExtensionConfig]{
			Data: ExampleExtensionConfig{
				Input:     "default",
				Password:  "123456",
				Email:     "app@hiddify.com",
				Selected:  false,
				Textarea:  "area",
				SwitchVal: true,
				Radiobox:  "A",
				Content:   "Welcome to Example Extension",
				Count:     10,
			},
		},
	}
}

func init() {
	ex.RegisterExtension(
		ex.ExtensionFactory{
			Id:          "github.com/hiddify/hiddify-app-example-extension/hiddify_extension",
			Title:       "Example Extension",
			Description: "An example extension",
			Builder:     NewExampleExtension,
		},
	)
}
