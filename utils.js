const validateResponse = (response) => {
    if (response === null || !("error" in response)) return true;
    alert(response.error);
    return false;
}


const refreshResult = (resultTextAreaName, response) => {
    if (response === null || !("count" in response)) return;

    const textArea = document.getElementById(resultTextAreaName)

    textArea.innerHTML = "";
    for (let i = 0; i < response.count; i++) {
        if (0 < i) textArea.innerHTML += '\n'
        textArea.innerHTML += response[`value${i}`];
    }

    if (textArea.style.height !== "auto") textArea.style.height = "auto";
    if (textArea.innerHTML.indexOf('\n') !== -1) textArea.style.height = `${textArea.scrollHeight}px`;

    textArea.select();
}

const refreshResultArray = (resultInputFixedName, response) => {
    if (response === null || !("count" in response)) return;

    for (let i = 0; i < response.count; i++) {
        const input = document.getElementById(`${resultInputFixedName}${i}`)
        input.value = response[`value${i}`];
        if (i === 0) input.select();
    }
}

document.addEventListener("DOMContentLoaded", () => {
    document.getElementById("title").innerText = document.title;
    document.getElementById("description").innerText = document.getElementsByName('description').item(0).content;
});
