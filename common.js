// -----------------------------------------------------------------------------
// メソッド
// -----------------------------------------------------------------------------

const goRun = () => {
    const go = new Go();
    WebAssembly.instantiateStreaming(fetch("./main.wasm"), go.importObject).then((source) => go.run(source.instance));
}

const validateResponse = (response) => {
    if (response === null || !("error" in response)) return true;
    alert(response.error);
    return false;
}

const refreshResultTextArea = (resultTextAreaName, response) => {
    if (response === null || !("count" in response)) return;

    // レスポンスをマッピング
    const textArea = document.getElementById(resultTextAreaName)
    textArea.innerHTML = "";
    for (let i = 0; i < response.count; i++) {
        if (0 < i) textArea.innerHTML += '\n'
        textArea.innerHTML += response[`value${i}`];
    }

    // テキストに応じてサイズを調整
    if (textArea.style.height !== "auto") textArea.style.height = "auto";
    if (textArea.innerHTML.indexOf('\n') !== -1) textArea.style.height = `${textArea.scrollHeight}px`;

    // テキストを全選択
    textArea.select();
}

const refreshResultInput = (resultInputFixedName, response) => {
    if (response === null || !("count" in response)) return;

    for (let i = 0; i < response.count; i++) {
        // レスポンスをマッピング
        const input = document.getElementById(`${resultInputFixedName}${i}`)
        input.value = response[`value${i}`];

        // 先頭要素を全選択
        if (i === 0) input.select();
    }
}

// -----------------------------------------------------------------------------
// イベントハンドラ
// -----------------------------------------------------------------------------

document.addEventListener("DOMContentLoaded", () => {
    // 共通要素を読み込む
    const load = (elementId, path, after = false) => {
        const request = new XMLHttpRequest();
        request.onreadystatechange = () => {
            if (request.readyState !== 4 || request.status !== 200) return;
            const element = document.getElementById(elementId);
            element.innerHTML = after ? element.innerHTML + request.responseText : request.responseText + element.innerHTML;
        }
        request.open("GET", path, true);
        request.send(null);
    }
    const pathName = location.pathname.substring(0, location.pathname.lastIndexOf('/'))
    const commonDir = 0 <= ['', '/toolbox'].indexOf(pathName) ? "./common" : "../common";
    load("layouts.header", `${commonDir}/header.html`);
    load("layouts.footer", `${commonDir}/footer.html`, true);

    // メタデータをマッピング
    const intervalId = setInterval(() => {
        const title = document.getElementById("title");
        const description = document.getElementById("description");
        if (title && description) {
            clearInterval(intervalId);
            title.innerText = document.title;
            description.innerText = document.getElementsByName('description').item(0).content;
        }
    }, 100);
});
