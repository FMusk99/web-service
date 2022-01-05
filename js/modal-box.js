// <!-- Trigger/Open The Modal -->
// <button id="myBtn">Open Modal</button>

// <!-- The Modal -->
// <div id="myModal" class="modal">

//   <!-- Modal content -->
//   <div class="modal-content">
//     <span class="close">&times;</span>
//     <p>Some text in the Modal..</p>
//   </div>
// </div>

function openModal(callback) {
   
    // only apply for one layer(one modal)
    let modal = document.getElementById("myModal");
    if(!modal) {
        var modalBox = document.createElement("DIV");   // Create a modal element
        modalBox.id = "myModal"
        modalBox.className = "modal"
        // modalBox.innerHTML = "CLICK ME";                   // Insert text
        var temp = document.getElementsByTagName("template")[0];
        var clon = temp.content.cloneNode(true);
        modalBox.appendChild(clon);
        document.body.appendChild(modalBox); 
        modal = modalBox
    }
    // Get the modal
    // Get the button that opens the modal
    // let btn = document.getElementById("jsonList");
    //Get the <span> element that closes the modal
    let span = document.getElementsByClassName("close")[0];

    // // When the user clicks on the button, open the modal
    // btn.onclick = function() {
    //     modal.style.display = "block";
    // }
    modal.style.display = "block";

    // // When the user clicks on <span> (x), close the modal
    span.onclick = function() {
        modal.style.display = "none";
    }

    //  When the user clicks anywhere outside of the modal, close it
    window.onclick = function(event) {
        if (event.target == modal) {
            modal.style.display = "none";
        }
    }
    callback()
}