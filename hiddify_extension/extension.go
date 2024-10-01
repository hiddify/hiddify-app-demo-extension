package hiddify_extension

import (
	"context"
	"fmt"
	"time"

	"github.com/hiddify/hiddify-core/config"
	"github.com/sagernet/sing-box/option"

	"github.com/fatih/color"
	ex "github.com/hiddify/hiddify-core/extension"
	ui "github.com/hiddify/hiddify-core/extension/ui"
)

// Color definitions for console output
var (
	red    = color.New(color.FgRed).Add(color.Bold)
	green  = color.New(color.FgGreen).Add(color.Underline)
	yellow = color.New(color.FgYellow)
)

// HiddifyAppDemoExtensionData holds the data specific to HiddifyAppDemoExtension
type HiddifyAppDemoExtensionData struct {
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

// HiddifyAppDemoExtension represents the core functionality of the extension
type HiddifyAppDemoExtension struct {
	ex.Base[HiddifyAppDemoExtensionData]                    // Embedding base extension functionality
	cancel                               context.CancelFunc // Function to cancel background tasks
	console                              string             // Stores console output
}

// backgroundTask runs a task in the background, updating the console at intervals
func (e *HiddifyAppDemoExtension) backgroundTask(ctx context.Context) {
	for count := 1; count <= e.Base.Data.Count; count++ {
		select {
		case <-ctx.Done(): // If context is done (cancel is pressed), exit the task
			e.cancel = nil
			e.addAndUpdateConsole(red.Sprint("Background Task Stoped")) // Notify cancellation
			return
		case <-time.After(1 * time.Second): // Wait for a second before the next iteration
			e.addAndUpdateConsole(red.Sprint(count), yellow.Sprint(" Background task ", count, " working..."))
		}
	}
	e.cancel = nil
	e.addAndUpdateConsole(green.Sprint("Background Task Finished Successfully")) // Task completion message
}

// addAndUpdateConsole adds messages to the console and updates the UI
func (e *HiddifyAppDemoExtension) addAndUpdateConsole(message ...any) {
	e.console = fmt.Sprintln(message...) + e.console // Prepend new messages to the console output
	e.UpdateUI(e.GetUI())                            // Update the UI with the new console content
}

// SubmitData processes and validates form submission data
func (e *HiddifyAppDemoExtension) SubmitData(button string, data map[string]string) error {
	switch button {
	case ui.ButtonDialogOk, ui.ButtonDialogClose:
		return nil
	case ui.ButtonCancel:
		return e.stop()
	case ui.ButtonSubmit:
		if err := e.setFormData(data); err != nil {
			e.ShowMessage("Invalid data", err.Error()) // Show error message if data is invalid
			return err                                 // Return the error
		}

		// stop any ongoing background task
		if e.cancel != nil {
			e.cancel()
		}

		// Create a new context for the task and store the cancel function
		ctx, cancel := context.WithCancel(context.Background())
		e.cancel = cancel

		// Run the background task concurrently
		go e.backgroundTask(ctx)

		return nil

	default:
		// Show message for undefined button actions
		return e.ShowMessage("Button "+button+" is pressed", "No action is defined for this button")
	}
}

// Stop stops the ongoing background task if it exists
func (e *HiddifyAppDemoExtension) stop() error {
	if e.cancel != nil {
		e.cancel()     // Stop the task
		e.cancel = nil // Clear the cancel function
	}
	return nil // Return nil after cancellation
}

// Stop is called when the extension is closed
func (e *HiddifyAppDemoExtension) Close() error {
	return e.stop() // Simply delegate to Stop
}

// To Modify user's config before connecting, you can use this function
func (e *HiddifyAppDemoExtension) BeforeAppConnect(hiddifySettings *config.HiddifyOptions, singconfig *option.Options) error {
	return nil
}

// NewHiddifyAppDemoExtension initializes a new instance of HiddifyAppDemoExtension with default values
func NewHiddifyAppDemoExtension() ex.Extension {
	return &HiddifyAppDemoExtension{
		Base: ex.Base[HiddifyAppDemoExtensionData]{
			Data: HiddifyAppDemoExtensionData{ // Set default data
				Input:     "default",
				Password:  "123456",
				Email:     "appdemo@extension.com",
				Selected:  false,
				Textarea:  "area",
				SwitchVal: true,
				Radiobox:  "A",
				Content:   "Welcome to Example Extension",
				Count:     10,
			},
		},
		console: yellow.Sprint("Welcome to ") + green.Sprint("hiddify-app-demo-extension\n"), // Default message
	}
}

// init registers the extension with the provided metadata
func init() {
	ex.RegisterExtension(
		ex.ExtensionFactory{
			Id:          "github.com/hiddify/hiddify_app_demo_extension/hiddify_extension",                      // Package identifier
			Title:       "Hiddify Demo Extension",                                                               // Display title of the extension
			Description: "This extension includes all the UI elements available in the Hiddify Extensions SDK.", // Brief description of the extension
			Builder:     NewHiddifyAppDemoExtension,                                                             // Function to create a new instance
		},
	)
}
