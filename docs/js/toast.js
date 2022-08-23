const toast = document.getElementById("toast");
let isVisible = false;

const showToast = (status, message) => {
    if (isVisible) return false

    toast.innerHTML = message
    toast.classList.add('is-' + status)
    toast.classList.add('is-show')
    isVisible = true
};

toast.addEventListener("animationend", () => {
    toast.innerHTML = "";
    toast.className = "init";
    isVisible = false;
});

document.getElementById("copy").addEventListener("click", (ev) => {
  showToast("success", "Copied!")
});
