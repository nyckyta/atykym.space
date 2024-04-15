const handleTerminalEnterHit = (event) => {
    // Prevent default is needed here to avoid inserting <br> tag clicking on enter, since 'new line' action
    // is done manually below.
    //
    // Btw, for some reasons this event produces two <br> tags inside a span on single event, though I
    // expected only one. I did not find any explanation regarding this. This works in the same way in both
    // Crhome and Firefox. Weird.
    event.preventDefault()

    let currentInput = document.getElementById("tty-main")
    let ttyClone = currentInput.cloneNode(true)
    ttyClone.setAttribute('id', '')

    let editableArea = ttyClone.children[1]
    editableArea.setAttribute('contenteditable', false)
    editableArea.setAttribute('id', '')
    editableArea.setAttribute('autofocus', false)

    currentInput.children[1].textContent = ""
    currentInput.parentElement.insertBefore(ttyClone, currentInput)
}

window.addEventListener("load", () => {
    window.addEventListener("click", () => {
        document.getElementById("tty-input").focus()
    }, false)

    document.getElementById("tty-input").addEventListener("keydown", (event) => {
        console.log(`Trigger event ${event.code}`)
        if (event.code == "Enter") {
            handleTerminalEnterHit(event)
        }
    })
})