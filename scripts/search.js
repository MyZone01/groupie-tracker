const searchInput = document.getElementById("search");
const searchForm = document.querySelector("form.search")
const suggestionsContainer = document.getElementById('suggestions');
function debounce(func, delay) {
    let timeoutId;

    return function () {
        const context = this;
        const args = arguments;

        clearTimeout(timeoutId);

        timeoutId = setTimeout(function () {
            func.apply(context, args);
        }, delay);
    };
}

function createSuggestionHeading(headingText) {
    const heading = document.createElement('h3');
    heading.textContent = headingText;
    return heading;
}

function createSuggestionList(suggestions) {
    const list = document.createElement('div');
    list.classList.add("list");

    suggestions.forEach(suggestion => {
        const item = document.createElement('a');
        wordLink = suggestion.split("@")
        item.textContent = wordLink[0];
        item.setAttribute("href", `/artist/${wordLink[1]}`)
        list.appendChild(item);
    });

    return list;
}

function showSuggestions(suggestions) {
    suggestionsContainer.innerHTML = '';
    if (!suggestionsContainer.classList.contains("show")) {
        suggestionsContainer.classList.add("show");
    }

    if (suggestions.Names && suggestions.Names.length > 0) {
        const namesHeading = createSuggestionHeading('Names');
        const namesList = createSuggestionList(suggestions.Names);
        suggestionsContainer.appendChild(namesHeading);
        suggestionsContainer.appendChild(namesList);
    }

    if (suggestions.Members && suggestions.Members.length > 0) {
        const membersHeading = createSuggestionHeading('Members');
        const membersList = createSuggestionList(suggestions.Members);
        suggestionsContainer.appendChild(membersHeading);
        suggestionsContainer.appendChild(membersList);
    }

    if (suggestions.Locations && suggestions.Locations.length > 0) {
        const locationsHeading = createSuggestionHeading('Locations');
        const locationsList = createSuggestionList(suggestions.Locations);
        suggestionsContainer.appendChild(locationsHeading);
        suggestionsContainer.appendChild(locationsList);
    }

    if (suggestions.FirstAlbums && suggestions.FirstAlbums.length > 0) {
        const albumsHeading = createSuggestionHeading('First Albums');
        const albumsList = createSuggestionList(suggestions.FirstAlbums);
        suggestionsContainer.appendChild(albumsHeading);
        suggestionsContainer.appendChild(albumsList);
    }

    if (suggestions.CreationDates && suggestions.CreationDates.length > 0) {
        const datesHeading = createSuggestionHeading('Creation Dates');
        const datesList = createSuggestionList(suggestions.CreationDates);
        suggestionsContainer.appendChild(datesHeading);
        suggestionsContainer.appendChild(datesList);
    }
}

const handleInput = async (event) => {
    event.preventDefault();
    if (searchInput.value !== "") {
        const formData = new FormData(searchForm);
        const response = await fetch('/suggestion', {
            method: 'POST',
            body: formData
        });
        const data = await response.text();
        suggestions = JSON.parse(data)
        showSuggestions(suggestions)
    } else {
        hideSuggestion()
    }
}

const hideSuggestion = () => {
    suggestionsContainer.classList.remove("show");
    suggestionsContainer.innerHTML = '';
}

const debouncedInputHandler = debounce(handleInput, 500);

searchInput.addEventListener('input', debouncedInputHandler);