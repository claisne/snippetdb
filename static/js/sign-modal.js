
(function () {
	var signVisible = false;
	var signLink = document.getElementById('sign-modal-link');
	var signModal = document.getElementById('sign-modal');
	var signModalWrapper = document.getElementById('sign-modal-wrapper');
	var signModalContainer = document.getElementById('sign-modal-container');

	if (signLink != null) {
        signLink.onclick = function (evt) {
            evt.preventDefault();
            signModalWrapper.className = 'sign-modal-show';
        };
	}

	if (signModal!= null) {
        signModal.onclick = function (evt) {
            evt.stopPropagation();
        };
	}

	if (signModalContainer != null) {
        signModalContainer.onclick = function (evt) {
            signModalWrapper.className = 'sign-modal-hidden';
        };
	}
})();
