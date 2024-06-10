package main

import (
    "fmt"
    "log"
    "note_taking_app/db"
    "note_taking_app/utils"
    "os"
    "path/filepath"
    "github.com/joho/godotenv"
    "fyne.io/fyne/v2"
    "fyne.io/fyne/v2/app"
    "fyne.io/fyne/v2/container"
    "fyne.io/fyne/v2/dialog"
    "fyne.io/fyne/v2/storage"
    "fyne.io/fyne/v2/theme"
    "fyne.io/fyne/v2/widget"
)

func main() {
    // Load environment variables from .env file
    err := godotenv.Load()
    if (err != nil) {
        log.Fatal("Error loading .env file")
    }

    // Initialize the database connection
    db.InitDB()
    log.Println("Database connection initialized")

    // Initialize Fyne application
    a := app.New()
    w := a.NewWindow("Note Taking App")

    // Create input field for note title
    noteTitle := widget.NewEntry()
    noteTitle.SetPlaceHolder("Enter note title...")

    // Create a text area for note content
    noteContent := widget.NewMultiLineEntry()
    noteContent.SetPlaceHolder("Write your note here...")
    noteContent.Wrapping = fyne.TextWrapWord

    // Create a scroll container for the text area
    scrollContainer := container.NewVScroll(noteContent)
    scrollContainer.SetMinSize(fyne.NewSize(600, 0.75*400))

    // Create a character counter label
    charCount := widget.NewLabel("")
    charCount.Hide()

    // Update character count when text changes
    noteContent.OnChanged = func(content string) {
        charCount.SetText(fmt.Sprintf("Character count: %d", len(content)))
    }

    // Function to toggle character counter
    toggleCharCount := func() {
        if charCount.Visible() {
            charCount.Hide()
        } else {
            charCount.Show()
        }
    }

    // Create buttons for rich text formatting
    boldButton := widget.NewButtonWithIcon("", theme.ContentBoldIcon(), func() {
        noteContent.SetText(noteContent.Text + "**bold**")
        charCount.SetText(fmt.Sprintf("Character count: %d", len(noteContent.Text)))
    })
    italicButton := widget.NewButtonWithIcon("", theme.ContentItalicIcon(), func() {
        noteContent.SetText(noteContent.Text + "*italic*")
        charCount.SetText(fmt.Sprintf("Character count: %d", len(noteContent.Text)))
    })
    underlineButton := widget.NewButtonWithIcon("", theme.ContentUndoIcon(), func() { // Placeholder icon for underline
        noteContent.SetText(noteContent.Text + "__underline__")
        charCount.SetText(fmt.Sprintf("Character count: %d", len(noteContent.Text)))
    })

    // Create save function
    saveNote := func() {
        title := noteTitle.Text
        content := noteContent.Text
        db.CreateNote(title, content)
        fmt.Println("Note saved to database")
    }

    // Create save as function
    saveNoteAs := func() {
        title := noteTitle.Text
        content := noteContent.Text

        // Get the user's Documents directory
        documentsDir, err := os.UserHomeDir()
        if err != nil {
            log.Println("Error getting user home directory:", err)
            return
        }
        documentsDir = filepath.Join(documentsDir, "Documents")

        // Set up the file save dialog
        saveDialog := dialog.NewFileSave(
            func(file fyne.URIWriteCloser, err error) {
                if err != nil {
                    log.Println("Error while saving file:", err)
                    return
                }
                if file == nil {
                    log.Println("Save file dialog was cancelled")
                    return
                }

                err = utils.SaveNoteToFile(file.URI().Path(), content)
                if err != nil {
                    log.Println("Error saving note to file:", err)
                } else {
                    fmt.Println("Note saved to file:", file.URI().Path())
                }

                file.Close()
            }, w)
        saveDialog.SetFileName(fmt.Sprintf("%s.txt", title))

        // Set the initial directory to Documents
        uri, err := storage.ListerForURI(storage.NewFileURI(documentsDir))
        if err != nil {
            log.Println("Error setting initial directory:", err)
            return
        }
        saveDialog.SetLocation(uri)

        saveDialog.Show()
    }

    // Create the menu
    saveMenu := fyne.NewMenu("File",
        fyne.NewMenuItem("Save", func() { saveNote() }),
        fyne.NewMenuItem("Save As", func() { saveNoteAs() }),
        fyne.NewMenuItemSeparator(),
        fyne.NewMenuItem("Toggle Character Counter", func() { toggleCharCount() }),
    )
    mainMenu := fyne.NewMainMenu(saveMenu)
    w.SetMainMenu(mainMenu)

    // Create a horizontal box for formatting buttons
    formatBox := container.NewHBox(boldButton, italicButton, underlineButton)

    // Create a vertical box to hold the UI components
    vbox := container.NewVBox(
        widget.NewLabel("My Note Taking App"),
        noteTitle,
        formatBox,
        scrollContainer,
        charCount,
    )

    // Set the window content and size
    w.SetContent(vbox)
    w.Resize(fyne.NewSize(600, 400))
    w.ShowAndRun()
}
