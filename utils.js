const validateResponse = (response) => {
    if (response == null || !("error" in response)) return true;
    console.log("Go Wasm response", response);
    alert(response.error);
    return false;
}


const refreshResult = (result, response) => {
    if (response == null || !("count" in response)) return;

    result.innerHTML = "";
    for (let i = 0; i < response.count; i++) {
        if (0 < i) result.innerHTML += '\n'
        result.innerHTML += response[`value${i}`];
    }

    if (result.style.height !== "auto") result.style.height = "auto";
    if (result.innerHTML.indexOf('\n') !== -1) result.style.height = `${result.scrollHeight}px`;

    result.select();
}

document.addEventListener("DOMContentLoaded", () => {
    document.getElementById("title").innerText = document.title;
    document.getElementById("description").innerText = document.getElementsByName('description').item(0).content;
});
