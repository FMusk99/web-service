var COMMON_NOTI = {
    ERROR: "ERROR",
    SUCCESS: "SUCCESS"
}


// alert a message that use for all resource
function popupMessage(message, type = COMMON_NOTI.SUCCESS, time =4000) {
    document.getElementById("closePopup") ?.remove();
    console.log('data')
    var alertTag = document.createElement("DIV");
    alertTag.id = "alert"
    alertTag.className = `alert alert--${type.toLowerCase()}`
    var contentTag = document.createElement("STRONG")
    var capitalText = type
    contentTag.innerHTML = `${capitalText.charAt(0).toUpperCase() + capitalText.slice(1).toLowerCase()}! `
    alertTag.appendChild(contentTag);
    var text = document.createTextNode(message);
    alertTag.appendChild(text)
    var closeTag = document.createElement("SPAN")
    closeTag.innerHTML = "&times"
    closeTag.classList.add("closebtn")
    closeTag.id = "closePopup"
    closeTag.addEventListener("click", closePopup)   
    alertTag.appendChild(closeTag)
    document.body.appendChild(alertTag)
    setTimeout(function(){
        closePopup.bind(document.getElementById('closePopup'))()
    }, time);

}

function closePopup() {
    this.parentElement ?.remove();
    this.parentElement.style.display='none'
}