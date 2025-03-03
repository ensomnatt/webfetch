const distroImageContainer = document.querySelector(".fastfetch-logo");
const fastFetchTextLineKey = document.querySelectorAll(".fastfetch-text-line-key");

fastFetchTextLineKey.forEach((key) => {
  if (key.textContent.includes("OS")) {
    const separator = key.nextElementSibling;
    const value = separator.nextElementSibling;
    const image = document.createElement("img")
    image.className = "fastfetch-logo-image"
    if (value.textContent.includes("Arch") || value.includes("arch")) {
      image.src = "/static/assets/Arch-linux-logo.png"
    } else if (value.textContent.includes("Gentoo") || value.includes("gentoo")) {
      image.src = "/static/assets/gentoo-logo.png"
    } else if (value.textContent.includes("Fedora") || value.includes("fedora")) {
      image.src = "/static/assets/fedora-logo.png"
    } else if (value.textContent.includes("Mint") || value.includes("mint")) {
      image.src = "/static/assets/linux-mint-logo.png"
    }

    distroImageContainer.appendChild(image)
  }
})
