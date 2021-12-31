document.getElementById("generateCateData") ?.addEventListener("click", generateCategories);
var xmlhtCate
function generateCategories() {
	createCategories(`${url}categories`)
}

function createCategories(url) {
	let categories = [
		{name: "Action and adventure", description: "Action and adventure"},
		{name: "Alternate history", description: "Alternate history"},
		{name: "Anthology", description: "Anthology"},
		{name: "Chick lit", description: "Chick lit"},
		{name: "Children's", description: "Children's"},
		{name: "Classic", description: "Classic"},
		{name: "Comic book", description: "Comic book"},
		{name: "Coming-of-age", description: "Coming-of-age"},
		{name: "Crime", description: "Crime"},
		{name: "Drama", description: "Drama"},
		{name: "Fairytale", description: "Fairytale"},
		{name: "Fantasy", description: "Fantasy"},
		{name: "Graphic novel", description: "Graphic novel"},
		{name: "Historical fiction", description: "Historical fiction"},
		{name: "Horror", description: "Horror"},
		{name: "Mystery", description: "Mystery"},
		{name: "Paranormal romance", description: "Paranormal romance"},
		{name: "Picture book", description: "Picture book"},
		{name: "Poetry", description: "Poetry"},
		{name: "Political thriller ", description: "Political thriller "},
		{name: "Romance", description: "Romance"},
		{name: "Satire", description: "Satire"},
		{name: "Science fiction", description: "Science fiction"},
		{name: "Short story", description: "Short story"},
		{name: "Suspense", description: "Suspense"},
		{name: "Thriller", description: "Thriller"},
		{name: "Western", description: "Western"},
		{name: "Young adult", description: "Young adult"},
		{name: "Art/architecture", description: "Art/architecture"},
		{name: "Autobiography", description: "Autobiography"},
		{name: "Biography", description: "Biography"},
		{name: "Business/economics", description: "Business/economics"},
		{name: "Crafts/hobbiesc", description: "Crafts/hobbies"},
		{name: "Cookbook", description: "Cookbook"},
		{name: "Diary", description: "Diary"},
		{name: "Dictionary", description: "Dictionary"},
		{name: "Encyclopedia", description: "Encyclopedia"},
		{name: "Guide", description: "Guide"},
		{name: "Health/fitness", description: "Health/fitness"},
		{name: "History", description: "History"},
		{name: "Home and garden", description: "Home and garden"},
		{name: "Humor", description: "Humor"},
		{name: "Journal", description: "Journal"},
		{name: "Math", description: "Math"},
		{name: "Memoir", description: "Memoir"},
		{name: "Philosophy", description: "Philosophy"},
		{name: "Prayer", description: "Prayer"},
		{name: "Religion, spirituality, and new age", description: "Religion, spirituality, and new age"},
		{name: "Textbook", description: "Textbook"},
		{name: "True crime", description: "True crime"},
		{name: "Review", description: "Review"},
		{name: "Science", description: "Science"},
		{name: "Self help", description: "Self help"},
		{name: "Sports and leisure", description: "Sports and leisure"},
		{name: "Travel", description: "Travel"},
		{name: "True crime", description: "True crime"}
	]

	var xhr = new XMLHttpRequest();
	xhr.onreadystatechange = function () {
	if (xhr.readyState == XMLHttpRequest.DONE) { // XMLHttpRequest.DONE == 4
		if (xhr.status == 200) {
			popupMessage('Created Category.', COMMON_NOTI.SUCCESS, 4000)
		} else if (xhr.status == 404) {
			popupMessage('Success', COMMON_NOTI.ERROR)
		} else {
			popupMessage('Success', COMMON_NOTI.ERROR)
			console.log(xhr)
		}
	}
	};
	  
	xhr.open("POST", "http://127.0.0.1:8080/v1/categories");
	xhr.setRequestHeader("Content-Type", "cls/json");
	xhr.send(JSON.stringify(categories));
}

