window.onload = function () {
  const go = new Go();
  let mod, instance;
  WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
    mod = result.module;
    instance = result.instance;
    // document.getElementById("gtree").disabled = false;

    console.clear();
    go.run(instance);
    instance = WebAssembly.instantiate(mod, go.importObject);

    // wasmで定義してる関数のloadが終わってないとダメなので以降の処理はこの位置にある
    if (typeof window.gtree !== 'function') {
      return;
    }

    gtree();
  });
};

const clearMarkdown = () => {
  document.getElementById("in").value = "";
};

const copyToClipboard = () => {
  const tree = document.getElementById("treeView");
  if (tree === null) {
    return;
  }
  const clipboard = window.navigator.clipboard;
  clipboard.writeText(tree.innerHTML);
};

const resetParts = () => {
  document.getElementById("parts1").value = "└";
  document.getElementById("parts2").value = "├";
  document.getElementById("parts3").value = "──";
  document.getElementById("parts4").value = "│";
};

const reset = () => {
  resetParts();
  gtree();
};

const clearTxt = () => {
  clearMarkdown();
  document.getElementById("treeView").remove();
};
