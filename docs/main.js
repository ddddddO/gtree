// Google tag (gtag.js) ref: https://source.sakuraweb.com/coding/gtag/
const gaScript = document.createElement('script');
gaScript.src = "https://www.googletagmanager.com/gtag/js?id=G-94DG34ENH8";
gaScript.async = true;
document.head.appendChild(gaScript);
window.dataLayer = window.dataLayer || [];
function gtag(){dataLayer.push(arguments);}
gtag('js', new Date());
gtag('config', 'G-94DG34ENH8');

const go = new Go();
let mod, instance;
WebAssembly.instantiateStreaming(fetch("main.wasm"), go.importObject).then((result) => {
    mod = result.module;
    instance = result.instance;
    document.getElementById("gtree").disabled = false;

    console.clear();
    go.run(instance);
    instance = WebAssembly.instantiate(mod, go.importObject);
});

const clearMarkdown = () => {
  document.getElementById("in").value = "";
  gtag('event', 'clear button', {'event_category': 'clear button','event_label': 'button click'});
};

const generateTree = () => {
  gtree();
  gtag('event', 'tree button', {'event_category': 'tree button','event_label': 'button click'});
};

const copyToClipboard = () => {
  const tree = document.getElementById("treeView");
  if (tree === null) {
    return;
  }
  const clipboard = window.navigator.clipboard;
  clipboard.writeText(tree.innerHTML);

  gtag('event', 'copy button', {'event_category': 'copy button','event_label': 'button click'});
};

