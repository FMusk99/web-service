var COMMON_NOTI = {
    ERROR: "ERROR",
    SUCCESS: "SUCCESS"
}
var w3review = document.querySelector("#w3review")
var jsonTestBtnTxt = document.querySelector("#jsonTestBtnTxt")
var count = 1
// alert a message that use for all resource
function popupMessage(message, type = COMMON_NOTI.SUCCESS, time =4000) {
    count += 1
    var alert = document.getElementById("alert");
    if(!alert) {
        alert = document.createElement("DIV");
        alert.id = "alert"
        alert.className = 'alert-container'
    }
    
    var alertTag = document.createElement("DIV");
    alertTag.className = `alert alert--${type.toLowerCase()} to-left-animation`
    var contentTag = document.createElement("STRONG")
    var capitalText = type
    contentTag.innerHTML = `${capitalText.charAt(0).toUpperCase() + capitalText.slice(1).toLowerCase()}!` //${count}
    alertTag.appendChild(contentTag);
    var text = document.createTextNode(message);
    alertTag.appendChild(text)
    document.getElementById("closePopup") ?.remove();
   
    // close icon
    var closeTag = document.createElement("SPAN")
    closeTag.innerHTML = "&times"
    closeTag.classList.add("closebtn")
    // closeTag.id = "closePopup"
   // close icon

    alertTag.appendChild(closeTag)
    closeTag.addEventListener("click", removeNode.bind(closeTag.parentNode))   
    alert.appendChild(alertTag)
    document.body.appendChild(alert)
    closePopup.bind(document.getElementById("alert").lastChild)(time)

}

window.onscroll = function() {onscrollPopup()};

function onscrollPopup() {
    var alert = document.getElementById("alert");
    if(alert) {
        var sticky = alert.offsetTop;
        if (window.pageYOffset > sticky) {
            alert.classList.remove("to-left-animation");
            alert.classList.add("delay-forward-animation");
        } else {
            alert.classList.remove("delay-forward-animation");
        }
    }
}

//close popupMessage
function closePopup(time) {
    setTimeout(() => {
        removeNode.bind(this)()
    }, time);
   
}

function removeNode() {
    if(this.style.display !== 'none' && this) {
        this.classList.add("face-out-animation")
        setTimeout(() => {
            this.remove();
            this.style.display='none'
            var x = document.getElementsByClassName("alert")
            x.length < 1 ?  document.getElementById("alert") ?.remove(): null
        }, 500);
    }
    
}


w3review ?.addEventListener("keydown", function(event) {
    // document.getElementById("output-box").innerHTML += "Sorry! <code>preventDefault()</code> won't let you check this!<br>";
    if(event ?.code === "KeyV") { // paste from clipboard
        event.preventDefault();
        navigator.clipboard.readText().then(data => {
            formatJsonView.bind(this)(data)
        })
        
    }
    
}, false);

jsonTestBtnTxt.onclick = function() {
    formatJsonView.bind(w3review)(w3review ?.value)
}
function formatJsonView(raw) {
    try {
        js = JSON.stringify(JSON.parse(raw,undefined, 4), undefined, 4);
        this.value = js
    } catch(e) {
        alert(e); // error in the above json (in this case, yes)!
    }
}

var btnJsonList = document.getElementById("jsonList");
// Get the modal
// Get the button that opens the modal

// When the user clicks on the button, open the modal
btnJsonList.onclick = function() {
    openModal(()=>{
		document.getElementById("jsonClonBook") ?.addEventListener("click", cloneBookRequest);
        document.getElementById("jsonClonCategory") ?.addEventListener("click", cloneCategoryRequest);
	})
}
