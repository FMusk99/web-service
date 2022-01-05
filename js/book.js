document.getElementById("searchBooks") ?.addEventListener("click", findByKeyword);
document.getElementById("generateBookData") ?.addEventListener("click", generateBooks);
document.getElementById("createBook") ?.addEventListener("click", createBook);
var jsonCreateBook = document.getElementById("jsonCreateBook") // Create book by json data in textarea element.
jsonCreateBook.onclick = function() {
	createBook(w3review ?.value)
}

function testFunc() {
	popupMessage("TEST", COMMON_NOTI.SUCCESS)
}
var keyword = document.getElementById("textIn")
var xmlht
var url = "http://127.0.0.1:8080/v1/"

var category = [
	{}
]
function findByKeyword() {
	getBooks(keyword ?.value ?? '')
}

function generateBooks() {
	createBooks(`${url}books`)
}


function loadFetchDoc(url) {
	fetch(url)
		.then(response => response.json())
		.then(res => {
			let ul = document.getElementById("result-data"); // get id of ul tag
			ul.classList.remove("skeletion-card")
			//e.firstElementChild can be used.
			if(res.code === 200) {
				for(let candidate of res.data) {
				let li = document.createElement("li");
				li.classList.add("result-box__item")
				li.classList.add("va-card")
				li.appendChild(document.createTextNode(`${candidate ?.name}( ${candidate ?.interests ?? ""} )`));
				ul.appendChild(li);
				}
			} else {
				let li = document.createElement("li");
				li.appendChild(document.createTextNode(`Not have data!`));
				ul.appendChild(li);
			}

	});
}

function getBooks(value) {
	let ul = document.getElementById("result-data"); // get id of ul tag
	//e.firstElementChild can be used.
	var child = ul.lastElementChild;
	while (child) {
		ul.removeChild(child);
		child = ul.lastElementChild;
	}
	ul.classList.add("skeletion-card")
	// Use AJAX for call web services
	xmlht ?.abort()
	xmlht = new XMLHttpRequest();
	xmlht.onreadystatechange = function () {
		if (xmlht.readyState == XMLHttpRequest.DONE) { // XMLHttpRequest.DONE == 4
			let ul = document.getElementById("result-data"); // get id of ul tag
			ul.classList.remove("skeletion-card")
			if (xmlht.status == 200) {
				const res = JSON.parse(xmlht.responseText)
				// let ul = document.getElementById("result-data"); // get id of ul tag
				if (res.code === 200) {
					for (let book of res.data) {
						let obj  = convertBook(book)
						let li = document.createElement("li");
						li.classList.add("result-box__item")
						li.classList.add("va-card")
						li.appendChild(document.createTextNode(`${obj ?.title} - ${obj ?.author}`));
						var closeTag = document.createElement("SPAN")
						closeTag.innerHTML = "&times"
						closeTag.classList.add("closebtn")
						closeTag.onclick = function() {
							deleteBook(book.id)
						}
						li.appendChild(closeTag)
						ul.appendChild(li);
					}
				} else if (res ?.code === 404){
					let li = document.createElement("li");
					li.appendChild(document.createTextNode(`Not have data!`));
					ul.appendChild(li);
				}
				
			} else if(res ?.code === 404) {
				popupMessage("Bad request!",COMMON_NOTI.ERROR)
			} else {
				popupMessage("Error server!",COMMON_NOTI.ERROR)
			}
		}
	};

	xmlht.open("GET", `${url}books?keyword=${value}`, true);
	xmlht.send();
}

function createBooks(url) {
	let books = [
		{title: "Clean code", author:"Robert C.Martin", pages :200, quantity: 1000},
		{title: "Clean architecture", author:"Robert C.Martin", pages :200, quantity: 3000},
		{title: "Types and Programming Languages", author:"Benjamin C. Pierce", pages :200, quantity: 1000},
		{title: "Code Complete", author:"Steve McConnell", pages :200, quantity: 1000},
		{title: "Algorithms to Live", author:"Brian Christian, Tom Griffiths", pages :200, quantity: 1000},
		{title: "The Computer Book ", author:"Simson L. Garfinkel, Rachel H. Grunspan", pages :200, quantity: 1000},
		{title: "The Second Machine Age ", author:"Erik Brynjolfsson and Andrew McAfee", pages :200, quantity: 1000},
		{title: "Explain the Cloud Like I’m 10 ", author:"Todd Hoff", pages :200, quantity: 1000},
		{title: "Structure and Interpretation of Computer Programs 2nd Edition", author:"Harold Abelson, Julie Sussman and Gerald Jay Sussman", pages :200, quantity: 1000},
		{title: "The Pragmatic Programmer: From Journey To Mastery", author:"David Thomas and Andrew Hunt", pages :200, quantity: 1000},
		{title: "Lauren Ipsum: A Story About Computer Science and Other Improbable Things", author:"Lauren Ipsum", pages :200, quantity: 1000},
		{title: "Game Theory for Security and Risk Management", author:"Stefan Rass, Stefan Schauer", pages :200, quantity: 1000},
		{title: "Stefan Rass, Stefan Schauer", author:"Charles N. Fischer, Richard J. Leblanc", pages :200, quantity: 1000},
		{title: "Quantum Computing Since Democritus", author:"Scott Aaronson", pages :200, quantity: 1000},
		{title: "Algorithms", author:"Robert Sedgewick and Kevin Wayne", pages :200, quantity: 1000},
		{title: "Robert Sedgewick and Kevin Wayne", author:"Steve Krug", pages :200, quantity: 1000},
		{title: "Modern Operating Systems", author:"Andrew Tannenbaum", pages :200, quantity: 1000},
		{title: "A Programming Approach to Computability", author:"A.J. Kfoury, Robert N. Moll, Michael A. Arbib", pages :200, quantity: 1000},
		{title: "Code Complete 2: a Practical Handbook of Software Construction", author:"Steve McConnell", pages :200, quantity: 1000},
		
		
		
	]
	var formData = new FormData();
	formData.append("book", JSON.stringify(books));
	xmlht ?.abort()
	xmlht = new XMLHttpRequest();
	xmlht.onreadystatechange = function () {
		if (xmlht.readyState == XMLHttpRequest.DONE) { // XMLHttpRequest.DONE == 4
			let ul = document.getElementById("result-data"); // get id of ul tag
			ul.classList.remove("skeletion-card")
			if (xmlht.status == 200) {
				popupMessage("created Books",COMMON_NOTI.SUCCESS)
			} else if (xmlht.status == 404) {
				// let li = document.createElement("li");
				// li.appendChild(document.createTextNode(`Not have data!`));
				// ul.appendChild(li);
			} else {
				console.log(xmlht)
				// alert('something else other than 200 was returned');
			}
		}
	};

	xmlht.open("POST", url, true);
	xmlht.send(formData);
}

function createBook(json) {
	if(!json){
		var raw = new Book()
		const formData = new FormData(document.getElementById('bookFrom'))
		for (var [key,value] of formData.entries()) {
			key === "pages" || key === "quantity" ?  raw[key] = parseInt(value) : raw[key] = value
		}
		console.log(raw)
		var json = JSON.stringify(raw);
	}

	var xhr = new XMLHttpRequest();
	xhr.onreadystatechange = function () {
	if (xhr.readyState == XMLHttpRequest.DONE) { // XMLHttpRequest.DONE == 4
		if (xhr.status == 200) {
			popupMessage('Success', COMMON_NOTI.SUCCESS, 4000)
		} else if (xhr.status == 404) {
			popupMessage("Bad request!",  COMMON_NOTI.ERROR)
		} else {
			popupMessage('error server!',  COMMON_NOTI.ERROR)
			console.log(xhr)
		}
	}
	};
	  
	xhr.open("POST", `${url}books`);
	xhr.setRequestHeader("Content-Type", "application/json");
	xhr.send(json);
	
}


function convertBook(obj) {
	return new Book(obj.id, obj.title, obj.author, obj.pages, obj.quantity, obj.created_at, obj.updated_at)
}

class Book {

	constructor(id, title, author, pages, quantity, created_at, updated_at) {
		this.id = id
		this.author = author
		this.title = title
		this.pages = parseInt(pages)
		this.quantity = parseInt(quantity)
		this.created_at = created_at
		this.updated_at = updated_at
	}
}

function cloneBookRequest() {
	/* Get the text field */
	var copyText = {title: "Clean code", author:"Robert C.Martin", pages :200, quantity: 1000};
  
	/* Select the text field */
	// copyText.select();
	// copyText.setSelectionRange(0, 99999); /* For mobile devices */
  
	 /* Copy the text inside the text field */
	navigator.clipboard.writeText(JSON.stringify(copyText));
	popupMessage("Clone data successfully!")
}

function deleteBook(id){
	console.log("delete: ", id)
	var xhr = new XMLHttpRequest();
	xhr.onreadystatechange = function () {
	if (xhr.readyState == XMLHttpRequest.DONE) { // XMLHttpRequest.DONE == 4
		if (xhr.status == 200) {
			popupMessage('Success', COMMON_NOTI.SUCCESS, 4000)
		} else if (xhr.status == 404) {
			popupMessage("Bad request!",  COMMON_NOTI.ERROR)
		} else {
			popupMessage('error server!',  COMMON_NOTI.ERROR)
			console.log(xhr)
		}
	}
	};
	  
	xhr.open("DELETE", `${url}books/${id}`);
	// xhr.setRequestHeader("Content-Type", "application/json");
	xhr.send();
}	